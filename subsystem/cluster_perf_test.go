package subsystem

import (
	"context"
	"encoding/json"
	"fmt"
	"net"
	"time"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/assisted-service/client/installer"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/models"
)

var _ = Describe("AAAAA", func() {

	var (
		ctx = context.Background()
	)

	BeforeEach(func() {
		// TODO delete all clusters with api

		clearDB()
	})
	It("AAAAAA", func() {
		worker := func(i int) {
			registerClusterReply, err := userBMClient.Installer.RegisterCluster(ctx, &installer.RegisterClusterParams{
				NewClusterParams: &models.ClusterCreateParams{
					BaseDNSDomain:            fmt.Sprintf("example%d.com", i),
					ClusterNetworkHostPrefix: 23,
					Name:                     swag.String(fmt.Sprintf("test-perf-cluster%d", i)),
					OpenshiftVersion:         swag.String(openshiftVersion),
					PullSecret:               swag.String(pullSecret),
					SSHPublicKey:             sshPublicKey,
				},
			})
			Expect(err).NotTo(HaveOccurred())
			cluster1 := registerClusterReply.GetPayload()
			registerHostsAndSetRolesDHCPPerf(*cluster1.ID, 5)
		}
		maxGoroutines := 10
		guard := make(chan struct{}, maxGoroutines)

		for i := 0; i < 150; i++ {
			guard <- struct{}{} // would block if guard channel is already filled
			go func(n int) {
				worker(n)
				<-guard
			}(i)
		}
		time.Sleep(1 * time.Hour)
	})
})

func registerHostsAndSetRolesDHCPPerf(clusterID strfmt.UUID, numHosts int) []*models.Host {
	ctx := context.Background()
	hosts := make([]*models.Host, 0)
	apiVip := "1.2.3.8"
	ingressVip := "1.2.3.9"

	for i := 0; i < numHosts; i++ {
		hostname := fmt.Sprintf("h%d", i)
		host := registerNode(ctx, clusterID, hostname)
		var role models.HostRoleUpdateParams
		if i < 3 {
			role = models.HostRoleUpdateParamsMaster
		} else {
			role = models.HostRoleUpdateParamsWorker
		}
		_, _ = userBMClient.Installer.UpdateCluster(ctx, &installer.UpdateClusterParams{
			ClusterUpdateParams: &models.ClusterUpdateParams{HostsRoles: []*models.ClusterUpdateParamsHostsRolesItems0{
				{ID: *host.ID, Role: role},
			}},
			ClusterID: clusterID,
		})
		hosts = append(hosts, host)
	}

	_, _ = userBMClient.Installer.UpdateCluster(ctx, &installer.UpdateClusterParams{
		ClusterUpdateParams: &models.ClusterUpdateParams{
			MachineNetworkCidr: swag.String("1.2.3.0/24"),
		},
		ClusterID: clusterID,
	})

	go func() {
		for {
			for i, h := range hosts {
				getNextStepsPerf(clusterID, *h.ID)
				generateDhcpStepReply(h, apiVip, ingressVip)
				generateHWPostReplies(ctx, h, fmt.Sprintf("h%d", i), validHwInfo)
			}
			generateFullMeshConnectivityPerf(ctx, "1.2.3.10", hosts...)
			time.Sleep(1 * time.Minute)
		}
	}()

	return hosts
}

func getNextStepsPerf(clusterID, hostID strfmt.UUID) models.Steps {
	steps, err := agentBMClient.Installer.GetNextSteps(context.Background(), &installer.GetNextStepsParams{
		ClusterID: clusterID,
		HostID:    hostID,
	})
	if err != nil {
		return models.Steps{}
	}
	return *steps.GetPayload()
}

func generateDhcpStepReply(h *models.Host, apiVip, ingressVip string) {
	avip := strfmt.IPv4(apiVip)
	ivip := strfmt.IPv4(ingressVip)
	r := models.DhcpAllocationResponse{
		APIVipAddress:     &avip,
		IngressVipAddress: &ivip,
	}
	b, err := json.Marshal(&r)
	Expect(err).ToNot(HaveOccurred())
	_, _ = agentBMClient.Installer.PostStepReply(context.TODO(), &installer.PostStepReplyParams{
		ClusterID: h.ClusterID,
		HostID:    *h.ID,
		Reply: &models.StepReply{
			ExitCode: 0,
			StepType: models.StepTypeDhcpLeaseAllocate,
			Output:   string(b),
			StepID:   string(models.StepTypeDhcpLeaseAllocate),
		},
	})
}

func generateHWPostReplies(ctx context.Context, h *models.Host, name string, inventory *models.Inventory) {
	generateHWPostStepReplyPerf(ctx, h, inventory, name)
	generateNTPPostStepReplyPerf(ctx, h, []*models.NtpSource{common.TestNTPSourceSynced})
}

func generateHWPostStepReplyPerf(ctx context.Context, h *models.Host, hwInfo *models.Inventory, hostname string) {
	hwInfo.Hostname = hostname
	hw, err := json.Marshal(&hwInfo)
	Expect(err).NotTo(HaveOccurred())
	_, _ = agentBMClient.Installer.PostStepReply(ctx, &installer.PostStepReplyParams{
		ClusterID: h.ClusterID,
		HostID:    *h.ID,
		Reply: &models.StepReply{
			ExitCode: 0,
			Output:   string(hw),
			StepID:   string(models.StepTypeInventory),
			StepType: models.StepTypeInventory,
		},
	})
}

func generateNTPPostStepReplyPerf(ctx context.Context, h *models.Host, ntpSources []*models.NtpSource) {
	response := models.NtpSynchronizationResponse{
		NtpSources: ntpSources,
	}
	bytes, err := json.Marshal(&response)
	Expect(err).NotTo(HaveOccurred())
	_, _ = agentBMClient.Installer.PostStepReply(ctx, &installer.PostStepReplyParams{
		ClusterID: h.ClusterID,
		HostID:    *h.ID,
		Reply: &models.StepReply{
			ExitCode: 0,
			Output:   string(bytes),
			StepID:   string(models.StepTypeNtpSynchronizer),
			StepType: models.StepTypeNtpSynchronizer,
		},
	})
}

func generateConnectivityPostStepReplyPerf(ctx context.Context, h *models.Host, connectivityReport *models.ConnectivityReport) {
	fa, err := json.Marshal(connectivityReport)
	Expect(err).NotTo(HaveOccurred())
	_, _ = agentBMClient.Installer.PostStepReply(ctx, &installer.PostStepReplyParams{
		ClusterID: h.ClusterID,
		HostID:    *h.ID,
		Reply: &models.StepReply{
			ExitCode: 0,
			Output:   string(fa),
			StepID:   string(models.StepTypeConnectivityCheck),
			StepType: models.StepTypeConnectivityCheck,
		},
	})
}

func generateFullMeshConnectivityPerf(ctx context.Context, startIPAddress string, hosts ...*models.Host) {
	ip := net.ParseIP(startIPAddress)
	hostToAddr := make(map[strfmt.UUID]string)

	for _, h := range hosts {
		hostToAddr[*h.ID] = ip.String()
		ip[len(ip)-1]++
	}

	var connectivityReport models.ConnectivityReport
	for _, h := range hosts {

		l2Connectivity := make([]*models.L2Connectivity, 0)
		for id, addr := range hostToAddr {

			if id == *h.ID {
				continue
			}

			l2Connectivity = append(l2Connectivity, &models.L2Connectivity{
				RemoteIPAddress: addr,
				Successful:      true,
			})
		}

		connectivityReport.RemoteHosts = append(connectivityReport.RemoteHosts, &models.ConnectivityRemoteHost{
			HostID:         *h.ID,
			L2Connectivity: l2Connectivity,
		})
	}

	for _, h := range hosts {
		generateConnectivityPostStepReplyPerf(ctx, h, &connectivityReport)
	}
}
