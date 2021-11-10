package aws

import (
	"github.com/openshift/assisted-service/internal/usage"
	"github.com/openshift/assisted-service/models"
)

func (p *awsProvider) CleanPlatformValuesFromDBUpdates(updates map[string]interface{}) error {
	updates[DbFieldRegion] = nil
	updates[DbFieldAccess] = nil
	updates[DbFieldSecret] = nil
	return nil
}

func (p *awsProvider) SetPlatformValuesInDBUpdates(
	platformParams *models.Platform, updates map[string]interface{}) error {
	if platformParams.Aws != nil {
		updates[DbFieldRegion] = platformParams.Aws.Region
		updates[DbFieldAccess] = platformParams.Aws.AccessKey
		updates[DbFieldSecret] = platformParams.Aws.Secret
	}
	return nil
}

func (p *awsProvider) SetPlatformUsages(
	platformParams *models.Platform,
	usages map[string]models.Usage,
	usageApi usage.API) error {

	withCredentials := platformParams.Aws != nil &&
		platformParams.Aws.Region != nil &&
		platformParams.Aws.AccessKey != nil &&
		platformParams.Aws.Secret != nil

	props := &map[string]interface{}{
		"platform_type":    p.Name(),
		"with_credentials": withCredentials}
	usageApi.Add(usages, usage.PlatformSelectionUsage, props)
	return nil
}
