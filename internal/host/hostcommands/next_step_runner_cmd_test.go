package hostcommands

import (
	"encoding/json"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/google/uuid"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/models"
)

func getNextStepRequest(args []string) *models.NextStepCmdRequest {
	request := models.NextStepCmdRequest{}
	err := json.Unmarshal([]byte(args[0]), &request)
	Expect(err).NotTo(HaveOccurred())
	return &request
}

var _ = Describe("Format command for starting next step agent", func() {

	var config NextStepRunnerConfig

	infraEnvId := strfmt.UUID(uuid.New().String())
	hostID := strfmt.UUID(uuid.New().String())
	serviceURL := uuid.New().String()
	image := uuid.New().String()

	BeforeEach(func() {
		config = NextStepRunnerConfig{
			InfraEnvID:          infraEnvId,
			HostID:              hostID,
			NextStepRunnerImage: image,
		}
	})

	It("standard formatting", func() {
		config.ServiceBaseURL = serviceURL
		command, args, err := GetNextStepRunnerCommand(&config)
		Expect(err).ToNot(HaveOccurred())
		Expect(command).Should(Equal(""))
		request := getNextStepRequest(*args)
		Expect(swag.BoolValue(request.Insecure)).Should(BeFalse())
		Expect(request.HostID.String()).Should(Equal(hostID.String()))
		Expect(request.InfraEnvID.String()).Should(Equal(infraEnvId.String()))
		Expect(swag.StringValue(request.BaseURL)).Should(Equal(serviceURL))
		Expect(request.CaCertPath).Should(BeEmpty())
		Expect(swag.StringValue(request.AgentVersion)).Should(Equal(image))
	})

	It("trim service URL", func() {
		config.ServiceBaseURL = fmt.Sprintf(" %s ", serviceURL)
		Expect(config.ServiceBaseURL).ShouldNot(Equal(serviceURL))
		_, args, err := GetNextStepRunnerCommand(&config)
		Expect(err).ToNot(HaveOccurred())
		request := getNextStepRequest(*args)
		Expect(swag.StringValue(request.BaseURL)).Should(Equal(serviceURL))
	})

	It("without custom CA certificate", func() {
		config.UseCustomCACert = false
		_, args, err := GetNextStepRunnerCommand(&config)
		Expect(err).ToNot(HaveOccurred())
		request := getNextStepRequest(*args)
		Expect(request.CaCertPath).Should(BeEmpty())
	})

	It("with custom CA certificate", func() {
		config.UseCustomCACert = true
		_, args, err := GetNextStepRunnerCommand(&config)
		Expect(err).ToNot(HaveOccurred())
		request := getNextStepRequest(*args)
		Expect(request.CaCertPath).Should(Equal(common.HostCACertPath))
	})

	It("certificate verification on", func() {
		config.SkipCertVerification = false
		_, args, err := GetNextStepRunnerCommand(&config)
		Expect(err).ToNot(HaveOccurred())
		request := getNextStepRequest(*args)
		Expect(swag.BoolValue(request.Insecure)).Should(BeFalse())
	})

	It("certificate verification off", func() {
		config.SkipCertVerification = true
		_, args, err := GetNextStepRunnerCommand(&config)
		Expect(err).ToNot(HaveOccurred())
		request := getNextStepRequest(*args)
		Expect(swag.BoolValue(request.Insecure)).Should(BeTrue())
	})
})
