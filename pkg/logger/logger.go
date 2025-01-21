package logger

import (
	"fmt"

	"github.com/pion/logging"
)

// Logger is a simple interface that can be implemented by any logging system
type Logger interface {
	Debug(args ...interface{})
	Debugf(format string, args ...interface{})
	Info(args ...interface{})
	Infof(format string, args ...interface{})
	Warn(args ...interface{})
	Warnf(format string, args ...interface{})
	Error(args ...interface{})
	Errorf(format string, args ...interface{})
	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})
	WithFields(fields map[string]interface{}) Logger
	WithField(key string, value interface{}) Logger
	WithError(err error) Logger
}

type defaultLog struct {
}

func (l *defaultLog) Debug(args ...interface{}) {
	fmt.Println(args...)
}

func (l *defaultLog) Info(args ...interface{}) {
	fmt.Println(args...)
}

func (l *defaultLog) Warn(args ...interface{}) {
	fmt.Println(args...)
}

func (l *defaultLog) Error(args ...interface{}) {
	fmt.Println(args...)
}

func (l *defaultLog) Fatal(args ...interface{}) {
	fmt.Println(args...)
}

func (l *defaultLog) Fatalf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *defaultLog) Debugf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *defaultLog) Infof(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *defaultLog) Warnf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *defaultLog) Errorf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}

func (l *defaultLog) WithFields(fields map[string]interface{}) Logger {
	return l
}

func (l *defaultLog) WithField(key string, value interface{}) Logger {
	return l
}

func (l *defaultLog) WithError(err error) Logger {
	return l
}

// pionLogger implements pion's LeveledLogger interface using our Logger
type pionLogger struct {
	logger Logger
}

func (l *pionLogger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *pionLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

func (l *pionLogger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *pionLogger) Infof(format string, args ...interface{}) {
	l.logger.Infof(format, args...)
}

func (l *pionLogger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *pionLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warnf(format, args...)
}

func (l *pionLogger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *pionLogger) Errorf(format string, args ...interface{}) {
	l.logger.Errorf(format, args...)
}

func (l *pionLogger) Trace(msg string) {
	l.logger.Debug(msg)
}

func (l *pionLogger) Tracef(format string, args ...interface{}) {
	l.logger.Debugf(format, args...)
}

// pionLoggerFactory implements pion's LoggerFactory interface
type pionLoggerFactory struct {
	logger Logger
}

// NewPionLoggerFactory creates a new pion logger factory that uses the provided logger
func NewPionLoggerFactory(logger Logger) logging.LoggerFactory {
	return &pionLoggerFactory{
		logger: logger,
	}
}

// NewLogger implements logging.LoggerFactory
func (f *pionLoggerFactory) NewLogger(scope string) logging.LeveledLogger {
	return &pionLogger{
		logger: f.logger.WithField("scope", scope),
	}
}

var defaultLogger Logger

func SetDefaultLogger(logger Logger) {
	defaultLogger = logger
}

func Info(args ...interface{}) {
	defaultLogger.Info(args...)
}

func Warn(args ...interface{}) {
	defaultLogger.Warn(args...)
}

func Error(args ...interface{}) {
	defaultLogger.Error(args...)
}

func Debug(args ...interface{}) {
	defaultLogger.Debug(args...)
}

func Trace(args ...interface{}) {
	defaultLogger.Debug(args...)
}

func GetLogger() Logger {
	return defaultLogger
}

func NewDefaultLogger() Logger {
	return &defaultLog{}
}

func init() {
	if defaultLogger == nil {
		defaultLogger = NewDefaultLogger()
	}
}
