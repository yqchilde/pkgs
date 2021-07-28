package logger

import (
	"fmt"

	"go.uber.org/zap"
)

var logger Logger
var zl *zap.Logger

type Fields map[string]interface{}

type Config struct {
	Development       bool   `mapstructure:"development"`
	DisableCaller     bool   `mapstructure:"disable-caller"`
	DisableStacktrace bool   `mapstructure:"disable-stacktrace"`
	Encoding          string `mapstructure:"encoding"`
	Level             string `mapstructure:"level"`
	Name              string `mapstructure:"name"`
	Writers           string `mapstructure:"writers"`
	LoggerFile        string `mapstructure:"logger-file"`
	LoggerWarnFile    string `mapstructure:"logger-warn-file"`
	LoggerErrorFile   string `mapstructure:"logger-error-file"`
	LogFormatText     bool   `mapstructure:"log-format-text"`
	LogRollingPolicy  string `mapstructure:"log-rolling-policy"`
	LogRotateDate     int    `mapstructure:"log-rotate-date"`
	LogRotateSize     int    `mapstructure:"log-rotate-size"`
	LogBackupCount    uint   `mapstructure:"log-backup-count"`
}

func Init(cfg *Config) Logger {
	var err error
	// new zap logger
	zl, err = newZapLogger(cfg)
	if err != nil {
		fmt.Errorf("init newZapLogger err: %v", err)
	}
	_ = zl

	// new sugar logger
	logger, err = newLogger(cfg)
	if err != nil {
		fmt.Errorf("init newLogger err: %v", err)
	}

	return logger
}

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

	Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
}

func GetLogger() Logger {
	return logger
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	logger.Debugf(format, args...)
}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Infof(format string, args ...interface{}) {
	logger.Infof(format, args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	logger.Warnf(format, args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	logger.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	logger.Fatalf(format, args...)
}

// Panicf logger
func Panicf(format string, args ...interface{}) {
	logger.Panicf(format, args...)
}

// WithFields logger
// output more field, eg:
// 		contextLogger := log.WithFields(log.Fields{"key1": "value1"})
// 		contextLogger.Info("print multi field")
// or more sample to use:
// 	    log.WithFields(log.Fields{"key1": "value1"}).Info("this is a test log")
// 	    log.WithFields(log.Fields{"key1": "value1"}).Infof("this is a test log, user_id: %d", userID)
func WithFields(keyValues Fields) Logger {
	return logger.WithFields(keyValues)
}
