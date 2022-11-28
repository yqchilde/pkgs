package log

// Logger 接口定义
type Logger interface {
	Println(args ...interface{})
	Printf(format string, args ...interface{})

	Debug(args ...interface{})
	Debugf(format string, args ...interface{})

	Warn(args ...interface{})
	Warnf(format string, args ...interface{})

	Error(args ...interface{})
	Errorf(format string, args ...interface{})

	Fatal(args ...interface{})
	Fatalf(format string, args ...interface{})

	Panic(args ...interface{})
	Panicf(format string, args ...interface{})

	WithFields(keyValues Fields) Logger
}

var (
	log Logger
	_   Logger = &zapLogger{}
)

type Fields map[string]interface{}

func Default(skip ...int) Logger {
	if len(skip) == 0 {
		skip = []int{1}
	}
	conf := &Config{
		Encoding: "console",
		Writers:  "console",
	}
	log = &zapLogger{sugarLogger: buildLogger(conf, skip[0]).Sugar()}
	return log
}

func Init(conf *Config, skip ...int) Logger {
	if len(skip) == 0 {
		skip = []int{1}
	}
	log = &zapLogger{sugarLogger: buildLogger(conf, skip[0]).Sugar()}
	return log
}

func GetLogger() Logger {
	return log
}

func Println(args ...interface{}) {
	log.Println(args...)
}

func Printf(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func Panic(args ...interface{}) {
	log.Panic(args...)
}

func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

func WithFields(keyValues Fields) Logger {
	return log.WithFields(keyValues)
}
