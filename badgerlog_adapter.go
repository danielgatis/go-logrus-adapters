package adapters

import (
	"fmt"

	"github.com/dgraph-io/badger/v3"
	"github.com/sirupsen/logrus"
)

var _ = (*badger.Logger)(nil)

// BadgerLogAdapter represents an badger.Logger adapter.
type BadgerLogAdapter struct {
	logger logrus.FieldLogger
}

// NewBadgerLogAdapter creates a new adapter.
func NewBadgerLogAdapter(logger logrus.FieldLogger) *BadgerLogAdapter {
	return &BadgerLogAdapter{
		logger: logger,
	}
}

// Errorf implements badger.Logger.Errorf.
func (adapter *BadgerLogAdapter) Errorf(msg string, args ...interface{}) {
	adapter.logger.Error(fmt.Sprintf(msg, args...))
}

// Warningf implements badger.Logger.Warningf.
func (adapter *BadgerLogAdapter) Warningf(msg string, args ...interface{}) {
	adapter.logger.Warn(fmt.Sprintf(msg, args...))
}

// Infof implements badger.Logger.Infof.
func (adapter *BadgerLogAdapter) Infof(msg string, args ...interface{}) {
	adapter.logger.Info(fmt.Sprintf(msg, args...))
}

// Debugf implements badger.Logger.Debugf.
func (adapter *BadgerLogAdapter) Debugf(msg string, args ...interface{}) {
	adapter.logger.Debug(fmt.Sprintf(msg, args...))
}
