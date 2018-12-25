package common

import (
	"github.com/spn/go/support/log"
)

const SpnAmountPrecision = 7

func CreateLogger(serviceName string) *log.Entry {
	return log.DefaultLogger.WithField("service", serviceName)
}
