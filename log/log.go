package log

import (
	"fmt"

	"github.com/pterm/pterm"
)

type logFunc func(msg string, args ...[]pterm.LoggerArgument)

type Loglevel int64

var (
	logger = pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)
)

const (
	LOG_ERROR Loglevel = iota
	LOG_INFO
	LOG_DEBUG
	LOG_TRACE
	LOG_WARN
	LOG_FATAL
)

func Error(m *map[string]any, s string, args ...interface{}) {
	Log(LOG_ERROR, m, s, args...)
}
func Info(m *map[string]any, s string, args ...interface{}) {
	Log(LOG_INFO, m, s, args...)
}
func Debug(m *map[string]any, s string, args ...interface{}) {
	Log(LOG_DEBUG, m, s, args...)
}
func Trace(m *map[string]any, s string, args ...interface{}) {
	Log(LOG_TRACE, m, s, args...)
}
func Warn(m *map[string]any, s string, args ...interface{}) {
	Log(LOG_WARN, m, s, args...)
}
func Fatal(m *map[string]any, s string, args ...interface{}) {
	Log(LOG_FATAL, m, s, args...)
}

func Log(level Loglevel, m *map[string]any, s string, args ...interface{}) {
	var f logFunc
	switch level {
	case LOG_ERROR:
		f = logger.Error
	case LOG_INFO:
		f = logger.Info
	case LOG_DEBUG:
		f = logger.Debug
	case LOG_TRACE:
		f = logger.Trace
	case LOG_WARN:
		f = logger.Warn
	case LOG_FATAL:
		f = logger.Fatal
	}
	if m != nil {
		f(fmt.Sprintf(s, args...), logger.ArgsFromMap(*m))
	} else {
		f(fmt.Sprintf(s, args...))
	}
}
