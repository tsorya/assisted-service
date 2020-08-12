package host

import (
	"context"
	"fmt"
	"strings"

	"github.com/jinzhu/gorm"
	"github.com/openshift/assisted-service/internal/cluster/validations"
	"github.com/openshift/assisted-service/internal/common"

	"github.com/sirupsen/logrus"

	"github.com/openshift/assisted-service/models"
)

type uploadLogsCmd struct {
	baseCmd
	instructionConfig InstructionConfig
	db                *gorm.DB
}

func NewUploadLogs(log logrus.FieldLogger, db *gorm.DB, instructionConfig InstructionConfig) *uploadLogsCmd {
	return &uploadLogsCmd{
		baseCmd:           baseCmd{log: log},
		instructionConfig: instructionConfig,
		db:                db,
	}
}

func (u *uploadLogsCmd) GetStep(ctx context.Context, host *models.Host) (*models.Step, error) {

	var cluster common.Cluster
	if err := u.db.First(&cluster, "id = ?", host.ClusterID).Error; err != nil {
		u.log.Errorf("failed to get cluster %s", host.ClusterID)
		return nil, err
	}

	creds, err := validations.ParsePullSecret(cluster.PullSecret)
	if err != nil {
		return nil, err
	}
	r, ok := creds["cloud.openshift.com"]
	if !ok {
		return nil, fmt.Errorf("pull secret does not contain auth for cloud.openshift.com")
	}

	step := &models.Step{
		StepType: models.StepTypeExecute,
		Command:  "/usr/bin/logs_sender",
		Args: []string{
			"-tag", "agent", "-tag", "installer", "-url", strings.TrimSpace(u.instructionConfig.ServiceBaseURL),
			"-cluster-id", string(host.ClusterID), "-host-id", string(*host.ID), "-pull-secret-token", r.AuthRaw,
		},
	}
	return step, nil
}
