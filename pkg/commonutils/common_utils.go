package commonutils

import (
	"strings"
	"time"

	"github.com/openshift/assisted-service/internal/metrics"
	"github.com/sirupsen/logrus"
)

func MeasureOperation(operation string, log logrus.FieldLogger, metricsApi metrics.API) func() {
	start := time.Now()
	log.Debugf("%s started at: %v", operation, start)
	return func() {
		duration := time.Since(start)
		log.Debugf("%s took : %v", operation, duration)
		if metricsApi != nil {
			metricsApi.Duration(operation, duration)
		}
	}
}

func SaveDiskPartitionsIsSet(installerArgs string) bool {
	needToSaveFlags := []string{"--save-partlabel", "--save-partindex"}
	for _, val := range needToSaveFlags {
		if strings.Contains(installerArgs, val) {
			return true
		}
	}
	return false
}
