package adapters

import (
	cron "github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

var _ cron.Logger = (*CronLogAdapter)(nil)

// CronLogAdapter represents an cron.Logger adapter.
type CronLogAdapter struct {
	logger logrus.FieldLogger
}

// NewCronLogAdapter creates a new adapter.
func NewCronLogAdapter(logger logrus.FieldLogger) *CronLogAdapter {
	return &CronLogAdapter{
		logger: logger,
	}
}

// Info implements cron.Logger.Info.
func (adapter *CronLogAdapter) Info(msg string, keysAndValues ...interface{}) {
	adapter.logger.Info(msg, keysAndValues)
}

// Errorf implements cron.Logger.Errorf.
func (adapter *CronLogAdapter) Error(_ error, msg string, keysAndValues ...interface{}) {
	adapter.logger.Error(msg, keysAndValues)
}
