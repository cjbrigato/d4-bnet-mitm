package log

import (
	"bytes"
	"fmt"
	"os"
	"sync"

	"github.com/pterm/pterm"
)

var Wmutex sync.Mutex

type logFunc func(msg string, args ...[]pterm.LoggerArgument)

type Loglevel int64

var (
	Writer = bytes.NewBufferString("")
	logger = pterm.Logger{
		Formatter:  pterm.LogFormatterColorful,
		Writer:     os.Stdout,
		Level:      pterm.LogLevelTrace,
		ShowTime:   true,
		TimeFormat: "2006-01-02 15:04:05",
		MaxWidth:   80,
		KeyStyles: map[string]pterm.Style{
			"error":  *pterm.NewStyle(pterm.FgRed, pterm.Bold),
			"err":    *pterm.NewStyle(pterm.FgRed, pterm.Bold),
			"caller": *pterm.NewStyle(pterm.FgGray, pterm.Bold),
		},
	}

	//logger = pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)
)

//go:generate stringer -type=Loglevel
const (
	LOG_ERROR Loglevel = iota
	LOG_INFO
	LOG_DEBUG
	LOG_TRACE
	LOG_WARN
	LOG_FATAL
)

type Logentry struct {
	Level   Loglevel
	m       *map[string]any
	Message string
	args    []interface{}
}

var LogEntries []*Logentry

func LogEntry(level Loglevel, m *map[string]any, s string, args ...interface{}) *Logentry {
	return &Logentry{Level: level, m: m, Message: s, args: args}
}

func Error(m *map[string]any, s string, args ...interface{}) {
	AddLogEntry(LOG_ERROR, m, s, args...)
}
func Info(m *map[string]any, s string, args ...interface{}) {
	AddLogEntry(LOG_INFO, m, s, args...)
}
func Warn(m *map[string]any, s string, args ...interface{}) {
	AddLogEntry(LOG_WARN, m, s, args...)
}

func Fatal(m *map[string]any, s string, args ...interface{}) {
	AddLogEntry(LOG_FATAL, m, s, args...)
}
func Debug(m *map[string]any, s string, args ...interface{}) {
	AddLogEntry(LOG_DEBUG, m, s, args...)
}
func Trace(m *map[string]any, s string, args ...interface{}) {
	AddLogEntry(LOG_TRACE, m, s, args...)
}
func AddLogEntry(level Loglevel, m *map[string]any, s string, args ...interface{}) {
	LogEntries = append(LogEntries, LogEntry(level, m, s, args...))
}

func PrintLogEntry(e *Logentry) {
	leftovers := 0
	maxm := 6
	n := make(map[string]any)
	i := 0
	if e.m != nil {

		if len(*e.m) >= 6 {
			leftovers = len(*e.m) - 6
		} else {
			maxm = len(*e.m) - 1
		}

		for k, v := range *e.m {
			if i == maxm {
				break
			}
			n[k] = v
			i++
		}
		pt := logger.ArgsFromMap(n)
		if leftovers > 0 {
			pt = append(pt, pterm.LoggerArgument{Key: "...", Value: fmt.Sprintf("and %d others...", leftovers)})
		}
		TruncLog(e.Level, pt, e.Message, e.args...)
	} else {
		Log(e.Level, nil, e.Message, e.args...)
	}
	//Log(e.Level, e.m, e.Message, e.args...)
}

func TruncLog(level Loglevel, ptargs []pterm.LoggerArgument, s string, args ...interface{}) {
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
	f(fmt.Sprintf(s, args...), ptargs)
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
