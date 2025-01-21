package logger

import (
	"fmt"

	"github.com/pion/logging"
)

// Logger is a simple interface that can be implemented by any logging system
type Logger interface {
	Trace(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
}

type defaultLog struct {
}

func (l *defaultLog) Trace(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

func (l *defaultLog) Debug(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

func (l *defaultLog) Info(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

func (l *defaultLog) Warn(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

func (l *defaultLog) Error(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

// pionLogger implements pion's LeveledLogger interface using our Logger
type pionLogger struct {
	logger Logger
}

func (l *pionLogger) Trace(msg string) {
	l.logger.Trace(msg)
}

func (l *pionLogger) Tracef(format string, args ...interface{}) {
	l.logger.Trace(format, args...)
}

func (l *pionLogger) Debug(msg string) {
	l.logger.Debug(msg)
}

func (l *pionLogger) Debugf(format string, args ...interface{}) {
	l.logger.Debug(format, args...)
}

func (l *pionLogger) Info(msg string) {
	l.logger.Info(msg)
}

func (l *pionLogger) Infof(format string, args ...interface{}) {
	l.logger.Info(format, args...)
}

func (l *pionLogger) Warn(msg string) {
	l.logger.Warn(msg)
}

func (l *pionLogger) Warnf(format string, args ...interface{}) {
	l.logger.Warn(format, args...)
}

func (l *pionLogger) Error(msg string) {
	l.logger.Error(msg)
}

func (l *pionLogger) Errorf(format string, args ...interface{}) {
	l.logger.Error(format, args...)
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
		logger: f.logger,
	}
}

var defaultLogger Logger

func SetDefaultLogger(logger Logger) {
	defaultLogger = logger
}

func Info(msg string, args ...interface{}) {
	defaultLogger.Info(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	defaultLogger.Warn(msg, args...)
}

func Error(msg string, args ...interface{}) {
	defaultLogger.Error(msg, args...)
}

func Debug(msg string, args ...interface{}) {
	defaultLogger.Debug(msg, args...)
}

func Trace(msg string, args ...interface{}) {
	defaultLogger.Trace(msg, args...)
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
