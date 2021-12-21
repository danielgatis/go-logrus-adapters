package adapters

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/hashicorp/go-hclog"
	"github.com/sirupsen/logrus"
)

var _ hclog.Logger = (*HCLogAdapter)(nil)

// HCLogAdapter represents an hclog.Logger adapter.
type HCLogAdapter struct {
	log  logrus.FieldLogger
	name string
	args []interface{}
}

// NewHCLogAdapter creates a new adapter.
func NewHCLogAdapter(logger logrus.FieldLogger, name string) *HCLogAdapter {
	return &HCLogAdapter{
		log:  logger,
		name: name,
	}
}

// Trace implements hclog.Logger.Trace.
func (adapter *HCLogAdapter) Trace(msg string, args ...interface{}) {
	adapter.CreateEntry(args).Trace(msg)
}

// Debug implements hclog.Logger.Debug.
func (adapter *HCLogAdapter) Debug(msg string, args ...interface{}) {
	adapter.CreateEntry(args).Debug(msg)
}

// Info implements hclog.Logger.Info.
func (adapter *HCLogAdapter) Info(msg string, args ...interface{}) {
	adapter.CreateEntry(args).Info(msg)
}

// Warn implements hclog.Logger.Warn.
func (adapter *HCLogAdapter) Warn(msg string, args ...interface{}) {
	adapter.CreateEntry(args).Warn(msg)
}

// Error implements hclog.Logger.Error.
func (adapter *HCLogAdapter) Error(msg string, args ...interface{}) {
	adapter.CreateEntry(args).Error(msg)
}

// Log implements hclog.Logger.Log.
func (adapter *HCLogAdapter) Log(level hclog.Level, msg string, args ...interface{}) {
	switch level {
	case hclog.Trace:
		adapter.Trace(msg, args...)
	case hclog.Debug:
		adapter.Debug(msg, args...)
	case hclog.Info:
		adapter.Info(msg, args...)
	case hclog.Warn:
		adapter.Warn(msg, args...)
	case hclog.Error:
		adapter.Error(msg, args...)
	}
}

// IsTrace implements hclog.Logger.IsTrace.
func (adapter *HCLogAdapter) IsTrace() bool {
	return false
}

// IsDebug implements hclog.Logger.IsDebug.
func (adapter *HCLogAdapter) IsDebug() bool {
	return adapter.shouldEmit(logrus.DebugLevel)
}

// IsInfo implements hclog.Logger.IsInfo.
func (adapter *HCLogAdapter) IsInfo() bool {
	return adapter.shouldEmit(logrus.InfoLevel)
}

// IsWarn implements hclog.Logger.IsWarn.
func (adapter *HCLogAdapter) IsWarn() bool {
	return adapter.shouldEmit(logrus.WarnLevel)
}

// IsError implements hclog.Logger.IsError.
func (adapter *HCLogAdapter) IsError() bool {
	return adapter.shouldEmit(logrus.ErrorLevel)
}

// SetLevel implements hclog.Logger.SetLevel.
func (adapter *HCLogAdapter) SetLevel(hclog.Level) { /* nope */ }

// With implements hclog.Logger.With.
func (adapter *HCLogAdapter) With(args ...interface{}) hclog.Logger {
	e := adapter.CreateEntry(args)

	return &HCLogAdapter{
		log:  e,
		args: concatFields(adapter.args, args),
	}
}

func concatFields(a, b []interface{}) []interface{} {
	c := make([]interface{}, len(a)+len(b))
	copy(c, a)
	copy(c[len(a):], b)

	return c
}

// ImpliedArgs implements hclog.Logger.ImpliedArgs.
func (adapter *HCLogAdapter) ImpliedArgs() []interface{} {
	return adapter.args
}

// Name implements hclog.Logger.Name.
func (adapter *HCLogAdapter) Name() string {
	return adapter.name
}

// Named implements hclog.Logger.Named.
func (adapter *HCLogAdapter) Named(name string) hclog.Logger {
	var newName bytes.Buffer
	if adapter.name != "" {
		newName.WriteString(adapter.name)
		newName.WriteString(".")
	}

	newName.WriteString(name)
	return adapter.ResetNamed(newName.String())
}

// ResetNamed implements hclog.Logger.ResetNamed.
func (adapter *HCLogAdapter) ResetNamed(name string) hclog.Logger {
	fields := []interface{}{"subsystem_name", name}
	e := adapter.CreateEntry(fields)
	return &HCLogAdapter{log: e, name: name}
}

// StandardLogger implements hclog.Logger.StandardLogger.
func (adapter *HCLogAdapter) StandardLogger(_ *hclog.StandardLoggerOptions) *log.Logger {
	entry := adapter.log.WithFields(logrus.Fields{})
	return log.New(entry.WriterLevel(logrus.InfoLevel), "", 0)
}

// StandardWriter implements hclog.Logger.StandardWriter.
func (adapter *HCLogAdapter) StandardWriter(_ *hclog.StandardLoggerOptions) io.Writer {
	var w io.Writer

	logger, ok := adapter.log.(*logrus.Logger)
	if ok {
		w = logger.Out
	}

	if w == nil {
		w = os.Stderr
	}

	return w
}

func (adapter *HCLogAdapter) shouldEmit(level logrus.Level) bool {
	return adapter.log.WithFields(logrus.Fields{}).Level >= level
}

// CreateEntry implements hclog.Logger.CreateEntry.
func (adapter *HCLogAdapter) CreateEntry(args []interface{}) *logrus.Entry {
	if len(args)%2 != 0 {
		args = append(args, "<unknown>")
	}

	fields := make(logrus.Fields)
	for i := 0; i < len(args); i += 2 {
		k, ok := args[i].(string)
		if !ok {
			continue
		}
		v := args[i+1]
		fields[k] = v
	}

	return adapter.log.WithFields(fields)
}

// Fatal implements hclog.Logger.Fatal.
func (adapter *HCLogAdapter) Fatal(err error) {
	adapter.Error(fmt.Sprintf("%v", err))
	os.Exit(1)
}
