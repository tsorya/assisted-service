package aws

import (
	"github.com/openshift/assisted-service/internal/common"
	"github.com/openshift/assisted-service/internal/installcfg"
	"github.com/openshift/assisted-service/models"
)

func setPlatformValues(platform *installcfg.AwsInstallConfigPlatform, clusterPlatform *models.AwsPlatform) {
	if clusterPlatform != nil {
		if clusterPlatform.Region != nil {
			platform.Region = *clusterPlatform.Region
		}
	}
}

func (p awsProvider) AddPlatformToInstallConfig(cfg *installcfg.InstallerConfigBaremetal, cluster *common.Cluster) error {
	if cluster.Platform.Aws != nil {
		awsPlatform := &installcfg.AwsInstallConfigPlatform{}
		setPlatformValues(awsPlatform, cluster.Platform.Aws)
		cfg.Platform = installcfg.Platform{
			Aws: awsPlatform,
		}
		cfg.ControlPlane.Replicas = 3

		if cluster.Platform.Aws.ComputeReplicas != nil && *cluster.Platform.Aws.ComputeReplicas > 0 {
			cfg.Compute[0].Replicas = int(*cluster.Platform.Aws.ComputeReplicas)
		}
	}
	return nil
}
