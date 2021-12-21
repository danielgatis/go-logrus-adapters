package adapters

import (
	"encoding/json"
	"io"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"github.com/sirupsen/logrus"
)

var _ echo.Logger = (*EchoLogAdapter)(nil)

// EchoLogAdapter represents an echo.Logger adapter.
type EchoLogAdapter struct {
	logger logrus.FieldLogger
}

// NewEchoLogAdapter creates a new adapter.
func NewEchoLogAdapter(logger logrus.FieldLogger) *EchoLogAdapter {
	return &EchoLogAdapter{
		logger: logger,
	}
}

// Output implements cron.Logger.Output.
func (adapter *EchoLogAdapter) Output() io.Writer {
	var w io.Writer

	logger, ok := adapter.logger.(*logrus.Logger)
	if ok {
		w = logger.Out
	}

	if w == nil {
		w = os.Stderr
	}

	return w
}

// SetOutput implements cron.Logger.SetOutput.
func (adapter *EchoLogAdapter) SetOutput(_ io.Writer) { /* nope */ }

// Level implements cron.Logger.Level.
func (adapter *EchoLogAdapter) Level() log.Lvl {
	w := logrus.InfoLevel

	logger, ok := adapter.logger.(*logrus.Logger)
	if ok {
		w = logger.GetLevel()
	}

	return toEchoLevel(w)
}

// SetLevel implements cron.Logger.SetLevel.
func (adapter *EchoLogAdapter) SetLevel(_ log.Lvl) { /* nope */ }

// SetHeader implements cron.Logger.SetHeader.
func (adapter *EchoLogAdapter) SetHeader(_ string) { /* nope */ }

// Formatter implements cron.Logger.Formatter.
func (adapter *EchoLogAdapter) Formatter() logrus.Formatter {
	var w logrus.Formatter

	logger, ok := adapter.logger.(*logrus.Logger)
	if ok {
		w = logger.Formatter
	}

	return w
}

// SetFormatter implements cron.Logger.SetFormatter.
func (adapter *EchoLogAdapter) SetFormatter(_ logrus.Formatter) { /* nope */ }

// Prefix implements cron.Logger.Prefix.
func (adapter *EchoLogAdapter) Prefix() string {
	return ""
}

// SetPrefix implements cron.Logger.SetPrefix.
func (adapter *EchoLogAdapter) SetPrefix(_ string) {

}

// Print implements cron.Logger.Print.
func (adapter *EchoLogAdapter) Print(i ...interface{}) {
	adapter.logger.Print(i...)
}

// Printf implements cron.Logger.Printf.
func (adapter *EchoLogAdapter) Printf(format string, args ...interface{}) {
	adapter.logger.Printf(format, args...)
}

// Printj implements cron.Logger.Printj.
func (adapter *EchoLogAdapter) Printj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	adapter.logger.Println(string(b))
}

// Debug implements cron.Logger.Debug.
func (adapter *EchoLogAdapter) Debug(i ...interface{}) {
	adapter.logger.Debug(i...)
}

// Debugf implements cron.Logger.Debugf.
func (adapter *EchoLogAdapter) Debugf(format string, args ...interface{}) {
	adapter.logger.Debugf(format, args...)
}

// Debugj implements cron.Logger.Debugj.
func (adapter *EchoLogAdapter) Debugj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	adapter.logger.Debugln(string(b))
}

// Info implements cron.Logger.Info.
func (adapter *EchoLogAdapter) Info(i ...interface{}) {
	adapter.logger.Info(i...)
}

// Infof implements cron.Logger.Infof.
func (adapter *EchoLogAdapter) Infof(format string, args ...interface{}) {
	adapter.logger.Infof(format, args...)
}

// Infoj implements cron.Logger.Infoj.
func (adapter *EchoLogAdapter) Infoj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	adapter.logger.Infoln(string(b))
}

// Warn implements cron.Logger.Warn.
func (adapter *EchoLogAdapter) Warn(i ...interface{}) {
	adapter.logger.Warn(i...)
}

// Warnf implements cron.Logger.Warnf.
func (adapter *EchoLogAdapter) Warnf(format string, args ...interface{}) {
	adapter.logger.Warnf(format, args...)
}

// Warnj implements cron.Logger.Warnj.
func (adapter *EchoLogAdapter) Warnj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	adapter.logger.Warnln(string(b))
}

// Error implements cron.Logger.Error.
func (adapter *EchoLogAdapter) Error(i ...interface{}) {
	adapter.logger.Error(i...)
}

// Errorf implements cron.Logger.Errorf.
func (adapter *EchoLogAdapter) Errorf(format string, args ...interface{}) {
	adapter.logger.Errorf(format, args...)
}

// Errorj implements cron.Logger.Errorj.
func (adapter *EchoLogAdapter) Errorj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	adapter.logger.Errorln(string(b))
}

// Fatal implements cron.Logger.Fatal.
func (adapter *EchoLogAdapter) Fatal(i ...interface{}) {
	adapter.logger.Fatal(i...)
}

// Fatalf implements cron.Logger.Fatalf.
func (adapter *EchoLogAdapter) Fatalf(format string, args ...interface{}) {
	adapter.logger.Fatalf(format, args...)
}

// Fatalj implements cron.Logger.Fatalj.
func (adapter *EchoLogAdapter) Fatalj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	adapter.logger.Fatalln(string(b))
}

// Panic implements cron.Logger.Panic.
func (adapter *EchoLogAdapter) Panic(i ...interface{}) {
	adapter.logger.Panic(i...)
}

// Panicf implements cron.Logger.Panicf.
func (adapter *EchoLogAdapter) Panicf(format string, args ...interface{}) {
	adapter.logger.Panicf(format, args...)
}

// Panicj implements cron.Logger.Panicj.
func (adapter *EchoLogAdapter) Panicj(j log.JSON) {
	b, err := json.Marshal(j)
	if err != nil {
		panic(err)
	}

	adapter.logger.Panicln(string(b))
}

func toEchoLevel(level logrus.Level) log.Lvl {
	switch level {
	case logrus.DebugLevel:
		return log.DEBUG
	case logrus.InfoLevel:
		return log.INFO
	case logrus.WarnLevel:
		return log.WARN
	case logrus.ErrorLevel:
		return log.ERROR
	}

	return log.OFF
}
