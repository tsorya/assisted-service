package log

import (
	"context"
	"runtime"
	"strings"
	"time"

	"github.com/openshift/assisted-service/internal/metrics"

	"github.com/openshift/assisted-service/pkg/requestid"
	"github.com/sirupsen/logrus"
)

//log formats as defined by LOG_FORMAT env variable
const (
	LogFormatText = "text"
	LogFormatJson = "json"
)

type Config struct {
	LogLevel  string `envconfig:"LOG_LEVEL" default:"info"`
	LogFormat string `envconfig:"LOG_FORMAT" default:"text"`
}

// FromContext equip a given logger with values from the given context
func FromContext(ctx context.Context, inner logrus.FieldLogger) logrus.FieldLogger {
	requestID := requestid.FromContext(ctx)
	return requestid.RequestIDLogger(inner, requestID).WithField("go-id", goid())
}

// get the low-level gorouting id
// This has been taken from:
// https://groups.google.com/d/msg/golang-nuts/Nt0hVV_nqHE/bwndAYvxAAAJ
// This is hacky and should not be used for anything but logging
func goid() string {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	return idField
}

func MeasureOperation(operation string, log logrus.FieldLogger, metricsApi metrics.API) func() {
	start := time.Now()
	return func() {
		duration := time.Since(start)
		log.Infof("%s took : %v", operation, duration)
		if metricsApi != nil {
			metricsApi.Duration(operation, duration)
		}
	}
}
