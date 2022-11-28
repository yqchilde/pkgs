package log

import (
	"fmt"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// WriterConsole 控制台打印
	WriterConsole = "console"
	// WriterFile 文件打印
	WriterFile = "file"
)

var zl *zap.Logger

// 日志级别映射
var loggerLevelMap = map[string]zapcore.Level{
	"debug": zapcore.DebugLevel,
	"info":  zapcore.InfoLevel,
	"warn":  zapcore.WarnLevel,
	"error": zapcore.ErrorLevel,
	"panic": zapcore.PanicLevel,
	"fatal": zapcore.FatalLevel,
}

type zapLogger struct {
	sugarLogger *zap.SugaredLogger
}

func buildLogger(conf *Config, skip int) *zap.Logger {
	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(fmt.Sprintf("[%s]", t.Format("2006-01-02 15:04:05")))
	}
	// 自定义日志级别输出格式
	customLevelEncoder := func(lvl zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(fmt.Sprintf("[%s]", lvl.CapitalString()))
	}
	// 自定义文件行号输出格式
	customCallerEncoder := func(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(fmt.Sprintf("[%s]", caller.TrimmedPath()))
	}

	encoderConf := zapcore.EncoderConfig{
		CallerKey:      "caller",
		LevelKey:       "level",
		MessageKey:     "msg",
		TimeKey:        "ts",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,
		EncodeLevel:    customLevelEncoder,
		EncodeCaller:   customCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// console/json 格式打印
	var encoder zapcore.Encoder
	if conf.Encoding == WriterConsole {
		encoder = zapcore.NewConsoleEncoder(encoderConf)
	} else {
		encoder = zapcore.NewJSONEncoder(encoderConf)
	}

	var cores []zapcore.Core
	var options []zap.Option
	writers := strings.Split(conf.Writers, ",")
	for _, w := range writers {
		switch w {
		case WriterConsole:
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), getLoggerLevel(conf)))
		case WriterFile:
			cores = append(cores, getAllCore(encoder, conf))
		default:
			cores = append(cores, zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), getLoggerLevel(conf)))
			cores = append(cores, getAllCore(encoder, conf))
		}
	}
	combineCore := zapcore.NewTee(cores...)
	options = append(options, zap.AddCaller())
	options = append(options, zap.AddCallerSkip(skip))

	return zap.New(combineCore, options...)
}

func getLoggerLevel(conf *Config) zapcore.LevelEnabler {
	level, exist := loggerLevelMap[conf.Level]
	if !exist {
		return zapcore.DebugLevel
	}
	return level
}

func getAllCore(encoder zapcore.Encoder, conf *Config) zapcore.Core {
	logRotatePolicy, err := getLogRollingPolicy(conf)
	if err != nil {
		panic(err)
	}
	allLevel := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl <= zapcore.FatalLevel
	})
	return zapcore.NewCore(encoder, zapcore.AddSync(logRotatePolicy), allLevel)
}

func (l *zapLogger) Println(args ...interface{}) {
	l.sugarLogger.Info(args...)
}

func (l *zapLogger) Printf(format string, args ...interface{}) {
	l.sugarLogger.Infof(format, args...)
}

func (l *zapLogger) Debug(args ...interface{}) {
	l.sugarLogger.Debug(args...)
}

func (l *zapLogger) Debugf(format string, args ...interface{}) {
	l.sugarLogger.Debugf(format, args...)
}

func (l *zapLogger) Warn(args ...interface{}) {
	l.sugarLogger.Warn(args...)
}

func (l *zapLogger) Warnf(format string, args ...interface{}) {
	l.sugarLogger.Warnf(format, args...)
}

func (l *zapLogger) Error(args ...interface{}) {
	l.sugarLogger.Error(args...)
}

func (l *zapLogger) Errorf(format string, args ...interface{}) {
	l.sugarLogger.Errorf(format, args...)
}

func (l *zapLogger) Fatal(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *zapLogger) Fatalf(format string, args ...interface{}) {
	l.sugarLogger.Fatalf(format, args...)
}

func (l *zapLogger) Panic(args ...interface{}) {
	l.sugarLogger.Fatal(args...)
}

func (l *zapLogger) Panicf(format string, args ...interface{}) {
	l.sugarLogger.Panicf(format, args...)
}
func (l *zapLogger) WithFields(fields Fields) Logger {
	var f = make([]interface{}, 0)
	for k, v := range fields {
		f = append(f, k)
		f = append(f, v)
	}
	newLogger := l.sugarLogger.With(f...)
	return &zapLogger{newLogger}
}
