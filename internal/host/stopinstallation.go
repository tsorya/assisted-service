package host

import (
	"context"
	"github.com/go-openapi/strfmt"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/openshift/assisted-service/models"
)

type stopInstallationCmd struct {
	baseCmd
	instructionConfig InstructionConfig
}

func NewStopInstallationCmd(log logrus.FieldLogger, instructionConfig InstructionConfig) *stopInstallationCmd {
	return &stopInstallationCmd{
		baseCmd: baseCmd{log: log},
		instructionConfig: instructionConfig,
	}
}

func (h *stopInstallationCmd) GetStep(ctx context.Context, host *models.Host) (*models.Step, error) {
	command := "/usr/bin/podman"
	cmdArgs := ""
	if host.LogsCollectedAt != strfmt.DateTime(time.Time{}) {
		logsCommand, err := CreateUploadLogsCmd(host, h.instructionConfig.ServiceBaseURL,
			h.instructionConfig.InventoryImage, h.instructionConfig.SkipCertVerification, false)
		if err == nil {
			cmdArgs += logsCommand
		}
	}


	step := &models.Step{
		StepType: models.StepTypeExecute,
		Command: command,
		Args: []string{
			"kill", "--all",
		},
	}
	return step, nil
}
