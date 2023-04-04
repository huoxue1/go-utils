// Package log
// this file is copy from https://github.com/sirupsen/logrus/blob/master/exported.go
package log

import (
	"context"
	"github.com/sirupsen/logrus"
	"io"
	"time"
)

var (
	globalLogger = New()
)

func StandardLogger() *Logger {
	return globalLogger
}

// SetOutput sets the standard logger output.
func SetOutput(out io.Writer) {
	globalLogger.SetOutput(out)
}

// SetFormatter sets the standard logger formatter.
func SetFormatter(formatter logrus.Formatter) {
	globalLogger.SetFormatter(formatter)
}

// SetReportCaller sets whether the standard logger will include the calling
// method as a field.
func SetReportCaller(include bool) {
	globalLogger.SetReportCaller(include)
}

// SetLevel sets the standard logger level.
func SetLevel(level logrus.Level) {
	globalLogger.SetLevel(level)
}

// GetLevel returns the standard logger level.
func GetLevel() logrus.Level {
	return globalLogger.GetLevel()
}

// IsLevelEnabled checks if the log level of the standard logger is greater than the level param
func IsLevelEnabled(level logrus.Level) bool {
	return globalLogger.IsLevelEnabled(level)
}

// AddHook adds a hook to the standard logger hooks.
func AddHook(hook logrus.Hook) {
	globalLogger.AddHook(hook)
}

// WithError creates an entry from the standard logger and adds an error to it, using the value defined in ErrorKey as key.
func WithError(err error) *logrus.Entry {
	return globalLogger.WithField(logrus.ErrorKey, err)
}

// WithContext creates an entry from the standard logger and adds a context to it.
func WithContext(ctx context.Context) *logrus.Entry {
	return globalLogger.WithContext(ctx)
}

// WithField creates an entry from the standard logger and adds a field to
// it. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *logrus.Entry {
	return globalLogger.WithField(key, value)
}

// WithFields creates an entry from the standard logger and adds multiple
// fields to it. This is simply a helper for `WithField`, invoking it
// once for each field.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithFields(fields logrus.Fields) *logrus.Entry {
	return globalLogger.WithFields(fields)
}

// WithTime creates an entry from the standard logger and overrides the time of
// logs generated with it.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithTime(t time.Time) *logrus.Entry {
	return globalLogger.WithTime(t)
}

// Trace logs a message at level Trace on the standard logger.
func Trace(args ...interface{}) {
	globalLogger.Trace(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	globalLogger.Debug(args...)
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	globalLogger.Print(args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	globalLogger.Info(args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	globalLogger.Warn(args...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(args ...interface{}) {
	globalLogger.Warning(args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	globalLogger.Error(args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	globalLogger.Panic(args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	globalLogger.Fatal(args...)
}

// TraceFn logs a message from a func at level Trace on the standard logger.
func TraceFn(fn logrus.LogFunction) {
	globalLogger.TraceFn(fn)
}

// DebugFn logs a message from a func at level Debug on the standard logger.
func DebugFn(fn logrus.LogFunction) {
	globalLogger.DebugFn(fn)
}

// PrintFn logs a message from a func at level Info on the standard logger.
func PrintFn(fn logrus.LogFunction) {
	globalLogger.PrintFn(fn)
}

// InfoFn logs a message from a func at level Info on the standard logger.
func InfoFn(fn logrus.LogFunction) {
	globalLogger.InfoFn(fn)
}

// WarnFn logs a message from a func at level Warn on the standard logger.
func WarnFn(fn logrus.LogFunction) {
	globalLogger.WarnFn(fn)
}

// WarningFn logs a message from a func at level Warn on the standard logger.
func WarningFn(fn logrus.LogFunction) {
	globalLogger.WarningFn(fn)
}

// ErrorFn logs a message from a func at level Error on the standard logger.
func ErrorFn(fn logrus.LogFunction) {
	globalLogger.ErrorFn(fn)
}

// PanicFn logs a message from a func at level Panic on the standard logger.
func PanicFn(fn logrus.LogFunction) {
	globalLogger.PanicFn(fn)
}

// FatalFn logs a message from a func at level Fatal on the standard logger then the process will exit with status set to 1.
func FatalFn(fn logrus.LogFunction) {
	globalLogger.FatalFn(fn)
}

// Tracef logs a message at level Trace on the standard logger.
func Tracef(format string, args ...interface{}) {
	globalLogger.Tracef(format, args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	globalLogger.Debugf(format, args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	globalLogger.Printf(format, args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	globalLogger.Infof(format, args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	globalLogger.Warnf(format, args...)
}

// Warningf logs a message at level Warn on the standard logger.
func Warningf(format string, args ...interface{}) {
	globalLogger.Warningf(format, args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	globalLogger.Errorf(format, args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	globalLogger.Panicf(format, args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	globalLogger.Fatalf(format, args...)
}

// Traceln logs a message at level Trace on the standard logger.
func Traceln(args ...interface{}) {
	globalLogger.Traceln(args...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(args ...interface{}) {
	globalLogger.Debugln(args...)
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	globalLogger.Println(args...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(args ...interface{}) {
	globalLogger.Infoln(args...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(args ...interface{}) {
	globalLogger.Warnln(args...)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(args ...interface{}) {
	globalLogger.Warningln(args...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(args ...interface{}) {
	globalLogger.Errorln(args...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(args ...interface{}) {
	globalLogger.Panicln(args...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(args ...interface{}) {
	globalLogger.Fatalln(args...)
}
