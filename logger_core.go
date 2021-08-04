package logger

import (
	"github.com/rs/zerolog"
	"os"
)

var log *logger

func init() {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05"
}

type logger struct {
	err    error
	logger zerolog.Logger
}

func Init() *logger {
	return &logger{
		logger: zerolog.New(os.Stdout).With().Caller().Timestamp().Logger(),
	}
}

func (l *logger) SetServer(serverName string) *logger {
	if serverName != "" {
		l.withStr("server-name", serverName)
	}
	return l
}

func (l *logger) SetLevel(level ...string) *logger {
	lev := "debug"
	if len(level) > 0 {
		lev = level[0]
	}
	logLevel, err := zerolog.ParseLevel(lev)
	l.err = err
	if err != nil {
		l.logger = l.logger.Level(logLevel)
	}
	return l
}

func (l *logger) Call() error {
	if l.err != nil {
		return l.err
	}
	log = l
	return nil
}

func (l *logger) withStr(key, value string) {
	l.logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str(key, value)
	})
}

func WithStr(key, value string) {
	log.logger.UpdateContext(func(c zerolog.Context) zerolog.Context {
		return c.Str(key, value)
	})
}

func Debug(k string) {
	debug().Msg(k)
}

func Info(k string) {
	info().Msg(k)
}

func Warn(k string) {
	warn().Msg(k)
}

func Error(k string) {
	err().Msg(k)
}

func Debugf(k string, v ...interface{}) {
	debug().Msgf(k, v...)
}

func Infof(k string, v ...interface{}) {
	info().Msgf(k, v...)
}

func Warnf(k string, v ...interface{}) {
	warn().Msgf(k, v...)
}

func Errorf(k string, v ...interface{}) {
	err().Msgf(k, v...)
}

////////// core //////////
func debug() *zerolog.Event {
	return log.logger.Debug()
}

func info() *zerolog.Event {
	return log.logger.Info()
}

func warn() *zerolog.Event {
	return log.logger.Warn()
}

func err() *zerolog.Event {
	return log.logger.Error()
}
