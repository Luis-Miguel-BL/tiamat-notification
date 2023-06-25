package logger

import (
	"fmt"
	"os"
	"reflect"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

type ZeroLogger struct {
	log     zerolog.Logger
	service string
}

func NewZerologger(service string) Logger {
	log.Logger = zerolog.New(os.Stderr).With().Str("service", service).Timestamp().Logger()
	return &ZeroLogger{
		log:     log.Logger,
		service: service,
	}
}

func (l *ZeroLogger) SetLevel(level LogLevel) {
	mapLevel := map[LogLevel]zerolog.Level{
		Trace: zerolog.TraceLevel,
		Info:  zerolog.InfoLevel,
		Debug: zerolog.DebugLevel,
		Warn:  zerolog.WarnLevel,
		Error: zerolog.ErrorLevel,
		Fatal: zerolog.FatalLevel,
	}

	zerolog.SetGlobalLevel(mapLevel[level])
}

func (l *ZeroLogger) parseFields(fields Fields) {
	for k, v := range fields {
		switch t := v.(type) {
		case string:
			l.log = l.log.With().Str(k, t).Logger()
		case int:
			l.log = l.log.With().Int(k, t).Logger()
		case int64:
			l.log = l.log.With().Int64(k, t).Logger()
		case float64:
			l.log = l.log.With().Float64(k, t).Logger()
		case bool:
			l.log = l.log.With().Bool(k, t).Logger()
		default:
			if v != nil {
				fmt.Printf("unknow field type %s for variable %v\n", reflect.TypeOf(v).Name(), k)
			} else {
				fmt.Printf("variable %s is nil\n", k)
			}
		}
	}
}

func (l *ZeroLogger) WithFields(fields Fields) {
	l.parseFields(fields)
}

func (l *ZeroLogger) NewWithFields(fields Fields) {
	l.log = zerolog.New(os.Stderr).With().Str("service", l.service).Timestamp().Logger()
	l.parseFields(fields)
}

func (l *ZeroLogger) Info(msg string) {
	l.log.Info().Msg(msg)
}

func (l *ZeroLogger) Debug(msg string) {
	l.log.Debug().Msg(msg)
}

func (l *ZeroLogger) Trace(msg string) {
	l.log.Trace().Msg(msg)
}

func (l *ZeroLogger) Warn(msg string) {
	l.log.Warn().Msg(msg)
}

func (l *ZeroLogger) Error(msg string) {
	l.log.Error().Msg(msg)
}

func (l *ZeroLogger) Fatal(msg string) {
	l.log.Fatal().Msg(msg)
}

func (l *ZeroLogger) Panic(msg string) {
	l.log.Panic().Msg(msg)
}

func (l *ZeroLogger) Infof(formatMsg string, args ...interface{}) {
	l.log.Info().Msgf(formatMsg, args...)
}

func (l *ZeroLogger) Debugf(formatMsg string, args ...interface{}) {
	l.log.Debug().Msgf(formatMsg, args...)
}

func (l *ZeroLogger) Tracef(formatMsg string, args ...interface{}) {
	l.log.Trace().Msgf(formatMsg, args...)
}

func (l *ZeroLogger) Warnf(formatMsg string, args ...interface{}) {
	l.log.Warn().Msgf(formatMsg, args...)
}

func (l *ZeroLogger) Errorf(formatMsg string, args ...interface{}) {
	l.log.Error().Msgf(formatMsg, args...)
}

func (l *ZeroLogger) Fatalf(formatMsg string, args ...interface{}) {
	l.log.Fatal().Msgf(formatMsg, args...)
}

func (l *ZeroLogger) Panicf(formatMsg string, args ...interface{}) {
	l.log.Panic().Msgf(formatMsg, args...)
}

func (l *ZeroLogger) Infoln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	l.log.Info().Msg(msg)
}

func (l *ZeroLogger) Debugln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	l.log.Debug().Msg(msg)
}

func (l *ZeroLogger) Traceln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	l.log.Trace().Msg(msg)
}

func (l *ZeroLogger) Warnln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	l.log.Warn().Msg(msg)
}

func (l *ZeroLogger) Errorln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	l.log.Error().Msg(msg)
}

func (l *ZeroLogger) Fatalln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	l.log.Fatal().Msg(msg)
}

func (l *ZeroLogger) Panicln(args ...interface{}) {
	msg := fmt.Sprintln(args...)
	l.log.Panic().Msg(msg)
}
