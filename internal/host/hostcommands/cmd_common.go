package hostcommands

import (
	"context"
	"strings"

	"github.com/openshift/assisted-service/models"
	"github.com/sirupsen/logrus"
)

type CommandGetter interface {
	GetSteps(ctx context.Context, host *models.Host) ([]*models.Step, error)
}

type baseCmd struct {
	CommandGetter
	log logrus.FieldLogger
}

func saveDiskPartitionsIsSet(installerArgs string) bool {
	needToSaveFlags := []string{"--save-partlabel", "--save-partindex"}
	for _, val := range needToSaveFlags {
		if strings.Contains(installerArgs, val) {
			return true
		}
	}
	return false
}