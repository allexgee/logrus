package logrus

import (
	"io"
)

var (
	// std is the name of the standard Logger in stdlib `log`
	std = New(InfoLevel)
)

// StandardLogger returns a new Logger which writes to os.Stderr
//with the default TextFormatter
func StandardLogger() *Logger {
	return std
}

// SetOutput sets the standard Logger output.
func SetOutput(out io.Writer) {
	std.mux.Lock()
	defer std.mux.Unlock()
	std.out = out
}

// SetFormatter sets the standard Logger formatter.
func SetFormatter(formatter Formatter) {
	std.mux.Lock()
	defer std.mux.Unlock()
	std.formatter = formatter
}

// SetLevel sets the standard Logger level.
func SetLevel(level Level) {
	std.mux.Lock()
	defer std.mux.Unlock()
	std.level = level
}

// AddHook adds a hook to the standard Logger hooks.
func AddHook(hook Hook) {
	std.mux.Lock()
	defer std.mux.Unlock()
	std.hooks.Add(hook)
}

// WithError creates an entry from the standard Logger and adds an error to the entry.
func WithError(err error) *Entry {
	return std.WithField(errorKey, err)
}

// WithField creates an entry from the standard Logger and adds a field to
// the entry. If you want multiple fields, use `WithFields`.
//
// Note that it doesn't log until you call Debug, Print, Info, Warn, Fatal
// or Panic on the Entry it returns.
func WithField(key string, value interface{}) *Entry {
	return std.WithField(key, value)
}

// WithFields creates an entry from the standard Logger and adds multiple
// fields to the entry.
//
// Note that it doesn't log until you call Write method on the entry object it returns.
func WithFields(fields Fields) *Entry {
	return std.WithFields(fields)
}

// Debug logs a Message at level Debug on the standard Logger.
func Debug(args ...interface{}) {
	std.Debug(args...)
}

// Info logs a Message at level Info on the standard Logger.
func Info(args ...interface{}) {
	std.Info(args...)
}

// Warning logs a Message at level Warn on the standard Logger.
func Warning(args ...interface{}) {
	std.Warning(args...)
}

// Error logs a Message at level Error on the standard Logger.
func Error(args ...interface{}) {
	std.Error(args...)
}

// Panic logs a Message at level Panic on the standard Logger.
func Panic(args ...interface{}) {
	std.Panic(args...)
}

// Fatal logs a Message at level Fatal on the standard Logger.
func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

// Debugf logs a Message at level Debug on the standard Logger.
func Debugf(format string, args ...interface{}) {
	std.Debugf(format, args...)
}

// Infof logs a Message at level Info on the standard Logger.
func Infof(format string, args ...interface{}) {
	std.Infof(format, args...)
}

// Warningf logs a Message at level Warn on the standard Logger.
func Warningf(format string, args ...interface{}) {
	std.Warningf(format, args...)
}

// Errorf logs a Message at level Error on the standard Logger.
func Errorf(format string, args ...interface{}) {
	std.Errorf(format, args...)
}

// Panicf logs a Message at level Panic on the standard Logger.
func Panicf(format string, args ...interface{}) {
	std.Panicf(format, args...)
}

// Fatalf logs a Message at level Fatal on the standard Logger.
func Fatalf(format string, args ...interface{}) {
	std.Fatalf(format, args...)
}

// Debugln logs a Message at level Debug on the standard Logger.
func Debugln(args ...interface{}) {
	std.Debugln(args...)
}

// Infoln logs a Message at level Info on the standard Logger.
func Infoln(args ...interface{}) {
	std.Infoln(args...)
}

// Warningln logs a Message at level Warn on the standard Logger.
func Warningln(args ...interface{}) {
	std.Warningln(args...)
}

// Errorln logs a Message at level Error on the standard Logger.
func Errorln(args ...interface{}) {
	std.Errorln(args...)
}

// Panicln logs a Message at level Panic on the standard Logger.
func Panicln(args ...interface{}) {
	std.Panicln(args...)
}

// Fatalln logs a Message at level Fatal on the standard Logger.
func Fatalln(args ...interface{}) {
	std.Fatalln(args...)
}
