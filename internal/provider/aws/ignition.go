package aws

import "github.com/openshift/assisted-service/internal/common"

func (p awsProvider) PreCreateManifestsHook(cluster *common.Cluster, envVars *[]string, workDir string) error {
	return nil
}
func (p awsProvider) PostCreateManifestsHook(cluster *common.Cluster, envVars *[]string, workDir string) error {
	return nil
}
