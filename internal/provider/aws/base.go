package aws

import (
	"github.com/openshift/assisted-service/internal/provider"
	"github.com/openshift/assisted-service/models"
	"github.com/sirupsen/logrus"
)

//
type awsProvider struct {
	Log logrus.FieldLogger
}

// NewVsphereProvider creates a new aws provider.
func NewAwsProvider(log logrus.FieldLogger) provider.Provider {
	return &awsProvider{
		Log: log,
	}
}

// Name returns the name of the provider
func (p *awsProvider) Name() models.PlatformType {
	return models.PlatformTypeAws
}

func (p *awsProvider) IsHostSupported(host *models.Host) (bool, error) {
	return true, nil
}

func (p *awsProvider) AreHostsSupported(hosts []*models.Host) (bool, error) {
	return true, nil
}
