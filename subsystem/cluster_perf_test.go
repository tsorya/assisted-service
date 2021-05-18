package subsystem

import (
	"context"
	"fmt"
	"github.com/go-openapi/swag"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/assisted-service/client/installer"
	"github.com/openshift/assisted-service/models"
)


var _ = Describe("AAAAA", func() {

	var (
		ctx     = context.Background()
	)

	BeforeEach(func() {
		clearDB()
	})

	Context("Filter by opensfhift cluster ID", func() {

		BeforeEach(func() {

			// _ = installCluster(*cluster.ID)
		})

		It("AAAAAA", func() {
			worker := func(i int) {
				registerClusterReply, err := userBMClient.Installer.RegisterCluster(ctx, &installer.RegisterClusterParams{
					NewClusterParams: &models.ClusterCreateParams{
						BaseDNSDomain:            fmt.Sprintf("example%d.com", i),
						ClusterNetworkHostPrefix: 23,
						Name:                     swag.String(fmt.Sprintf("test-cluster%d", i)),
						OpenshiftVersion:         swag.String(openshiftVersion),
						PullSecret:               swag.String(pullSecret),
						SSHPublicKey:             sshPublicKey,
					},
				})
				Expect(err).NotTo(HaveOccurred())
				cluster1 := registerClusterReply.GetPayload()
				registerHostsAndSetRolesDHCP(*cluster1.ID, 5)
			}
			maxGoroutines := 10
			guard := make(chan struct{}, maxGoroutines)

			for i := 0; i < 300; i++ {
				guard <- struct{}{} // would block if guard channel is already filled
				go func(n int) {
					worker(n)
					<-guard
				}(i)
			}
		})

	})
})
