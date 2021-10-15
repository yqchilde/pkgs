package log

import (
	"fmt"
)

type Logger interface {
	Info(args ...interface{})
	Infof(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
}

type Fields map[string]interface{}

var log Logger

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

func newLogger(cfg *Config) (Logger, error) {
	return &zapLogger{sugarLogger: buildLogger(cfg).Sugar()}, nil
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
	log, err = newLogger(cfg)
	if err != nil {
		fmt.Errorf("init newLogger err: %v", err)
	}
	return log
}

func GetLogger() Logger {
	return log
}

// Info log
func Info(args ...interface{}) {
	log.Info(args...)
}

// Warn log
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Error log
func Error(args ...interface{}) {
	log.Error(args...)
}

// Infof log
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warnf log
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Errorf log
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// WithFields log
func WithFields(keyValues Fields) Logger {
	return log.WithFields(keyValues)
}
