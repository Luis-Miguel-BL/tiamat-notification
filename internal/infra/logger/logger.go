package logger

type Fields map[string]interface{}

type LogLevel string

const (
	Trace LogLevel = "trace"
	Info  LogLevel = "info"
	Debug LogLevel = "debug"
	Warn  LogLevel = "warn"
	Error LogLevel = "error"
	Fatal LogLevel = "fatal"
)

type Logger interface {
	SetLevel(level LogLevel)
	WithFields(fields Fields)
	NewWithFields(fields Fields)

	Info(msg string)
	Debug(msg string)
	Trace(msg string)
	Warn(msg string)
	Error(msg string)
	Fatal(msg string)
	Panic(msg string)

	Infof(formatMsg string, args ...interface{})
	Debugf(formatMsg string, args ...interface{})
	Tracef(formatMsg string, args ...interface{})
	Warnf(formatMsg string, args ...interface{})
	Errorf(formatMsg string, args ...interface{})
	Fatalf(formatMsg string, args ...interface{})
	Panicf(formatMsg string, args ...interface{})

	Infoln(args ...interface{})
	Debugln(args ...interface{})
	Traceln(args ...interface{})
	Warnln(args ...interface{})
	Errorln(args ...interface{})
	Fatalln(args ...interface{})
	Panicln(args ...interface{})
}
