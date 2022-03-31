package commonutils

import (
	"fmt"
	"github.com/openshift/assisted-service/models"
)

func GetDeviceIdentifier(installationDisk *models.Disk) string {
	// We changed the host.installationDiskPath to contain the disk id instead of the disk path.
	// Here we updates the old installationDiskPath to the disk id.
	// (That's the reason we return the disk.ID instead of the previousInstallationDisk if exist)
	if installationDisk.ID != "" {
		return installationDisk.ID
	}

	// Old inventory or a bug
	return GetDeviceFullName(installationDisk)
}

func GetDeviceFullName(installationDisk *models.Disk) string {
	return fmt.Sprintf("/dev/%s", installationDisk.Name)
}
