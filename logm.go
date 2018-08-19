package logm

import (
	"fmt"
	"io"
	golog "log"
	"os"
)

var (
	colReset   = string([]byte{27, 91, 48, 109})
	colBlue    = string([]byte{27, 91, 57, 55, 59, 52, 52, 109}) // LvlInfo
	colGreen   = string([]byte{27, 91, 57, 55, 59, 52, 50, 109}) // LvlOk
	colCyan    = string([]byte{27, 91, 57, 55, 59, 52, 54, 109}) // LvlNotice
	colYellow  = string([]byte{27, 91, 57, 55, 59, 52, 51, 109}) // LvlWarning
	colRed     = string([]byte{27, 91, 57, 55, 59, 52, 49, 109}) // LvlError
	colMagenta = string([]byte{27, 91, 57, 55, 59, 52, 53, 109}) // LvlPanic
	colWhite   = string([]byte{27, 91, 57, 48, 59, 52, 55, 109}) // LvlHigh
)

// Log Level constants used for log-level filtering and coloring
const (
	LvlVerbose = 0
	LvlInfo    = 1
	LvlOk      = 2
	LvlNotice  = 3
	LvlWarning = 4
	LvlError   = 5
	LvlPanic   = 6

	LvlMute = 100
)

// Logm is an instance of the logger
type Logm struct {
	logger    *golog.Logger
	logPrefix string

	// DisableColor specifies if log output in console should use colors
	DisableColor bool // = false

	/*
		Log level to be output. Verbose level (0) enables all, info level (1) enables info and more
		important, warning level (2) enables warning, error and higher level, etc.
	*/
	LogLevel int
}

// New instantiates a new logger with a given log prefix
func New(logPrefix string) Logm {
	return Logm{
		logger:    golog.New(os.Stderr, "", golog.LstdFlags),
		logPrefix: logPrefix,
	}
}

// NewWithOutput instantiates a new logger with a given output writer and log prefix
func NewWithOutput(out io.Writer, logPrefix string) Logm {
	return Logm{
		logger:    golog.New(out, "", golog.LstdFlags),
		logPrefix: logPrefix,
	}
}

// Verbose logs a verbose message
func (l Logm) Verbose(message string, args ...interface{}) {
	l.Log(LvlVerbose, message, args...)
}

// Info logs an info message
func (l Logm) Info(message string, args ...interface{}) {
	l.Log(LvlInfo, message, args...)
}

// Warning logs a warning message
func (l Logm) Warning(message string, args ...interface{}) {
	l.Log(LvlWarning, message, args...)
}

// Error logs an error message
func (l Logm) Error(message string, args ...interface{}) {
	l.Log(LvlError, message, args...)
}

// Log logs a message with given log level. Loglevel should be one of the Lvl* constants from this
// package.
func (l Logm) Log(logLevel int, logMessage string, logArgs ...interface{}) {

	color, text := l.getProps(logLevel)

	if l.LogLevel <= logLevel {
		l.logger.Printf("[%s] |%s %s %s| %s\n",
			l.logPrefix,
			color, text, colReset,
			fmt.Sprintf(logMessage, logArgs...),
		)
	}

}

func (l Logm) getProps(lvl int) (color string, text string) {

	switch lvl {
	case LvlVerbose:
		color = colReset
		text = "VERBOSE"
		break
	case LvlInfo:
		color = colBlue
		text = "INFO   "
		break
	case LvlOk:
		color = colGreen
		text = "OK     "
		break
	case LvlNotice:
		color = colCyan
		text = "NOTICE "
		break
	case LvlWarning:
		color = colYellow
		text = "WARNING"
		break
	case LvlError:
		color = colRed
		text = "ERROR  "
		break
	case LvlPanic:
		color = colMagenta
		text = "PANIC  "
		break
	default:
		color = colWhite
		text = fmt.Sprintf("!%2d    ", lvl)
	}

	if l.DisableColor {
		color = colReset
	}

	return color, text
}
