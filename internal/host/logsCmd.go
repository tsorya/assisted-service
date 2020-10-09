package host

import (
	"context"
	"time"

	"github.com/go-openapi/strfmt"

	"github.com/sirupsen/logrus"

	"github.com/openshift/assisted-service/models"
)

type logsCmd struct {
	baseCmd
	instructionConfig InstructionConfig
}

func NewLogsCmd(log logrus.FieldLogger, instructionConfig InstructionConfig) *logsCmd {
	return &logsCmd{
		baseCmd:           baseCmd{log: log},
		instructionConfig: instructionConfig,
	}
}

func (i *logsCmd) GetStep(ctx context.Context, host *models.Host) (*models.Step, error) {
	// added to run upload logs if install command fails
	if host.LogsCollectedAt != strfmt.DateTime(time.Time{}) {
		return nil, nil
	}
	logsCommand, err := CreateUploadLogsCmd(host, i.instructionConfig.ServiceBaseURL,
		i.instructionConfig.InventoryImage, i.instructionConfig.SkipCertVerification, false)
	if err != nil {
		return nil, err
	}
	step := &models.Step{
		StepType: models.StepTypeExecute,
		Command:  "/usr/bin/podman",
		Args: []string{
			logsCommand,
		},
	}

	return step, nil
}
