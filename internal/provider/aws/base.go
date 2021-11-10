package aws

import (
	"fmt"

	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/provider"
	"github.com/openshift/assisted-service/models"
	"github.com/sirupsen/logrus"
)

//
type awsProvider struct {
	Log logrus.FieldLogger
}

// NewVsphereProvider creates a new vSphere provider.
func NewAwsProvider(log logrus.FieldLogger) provider.Provider {
	return &awsProvider{
		Log: log,
	}
}

// Name returns the name of the provider
func (p *awsProvider) Name() models.PlatformType {
	return models.
}

func (p *awsProvider) IsHostSupported(host *models.Host) (bool, error) {
	return true, nil
}

func (p *awsProvider) AreHostsSupported(hosts []*models.Host) (bool, error) {
	return true, nil
}
