package logm

import (
	golog "log"
	"time"
)

var (
	blue   = string([]byte{27, 91, 57, 55, 59, 52, 52, 109}) // INF
	yellow = string([]byte{27, 91, 57, 55, 59, 52, 51, 109}) // WRN
	red    = string([]byte{27, 91, 57, 55, 59, 52, 49, 109}) // ERR
	cyan   = string([]byte{27, 91, 57, 55, 59, 52, 54, 109}) // ???
	reset  = string([]byte{27, 91, 48, 109})
)

// Logm is an instance of the logger
type Logm struct {
	// DisableColor specifies if log output in console should use colors
	DisableColor bool // = false
	// LogLevel specifies what logs should be output. Level 0 enables all output, level 1 disables VRB
	// output, level 2 disables VRN and INF, and so on... Level 4 will disable all known output types
	// (VRB, INF, WAR, ERR), and level 5 or higher will disable all output (including output with
	// custom log type).
	LogLevel  int // = 0
	logPrefix string
}

// New instantiates a new logger with a given log prefix
func New(logPrefix string) Logm {
	return Logm{
		logPrefix: logPrefix,
	}
}

// LogV logs a verbose message
func (l Logm) LogV(logMessage string) {
	if l.LogLevel <= 0 {
		l.Log("VRB", logMessage)
	}
}

// LogI logs an info message
func (l Logm) LogI(logMessage string) {
	if l.LogLevel <= 1 {
		l.Log("INF", logMessage)
	}
}

// LogW logs a warning message
func (l Logm) LogW(logMessage string) {
	if l.LogLevel <= 2 {
		l.Log("WRN", logMessage)
	}
}

// LogE logs an error message
func (l Logm) LogE(logMessage string) {
	if l.LogLevel <= 3 {
		l.Log("ERR", logMessage)
	}
}

// Log logs a message with custom log level
func (l Logm) Log(logLevel string, logMessage string) {

	var color string
	if !l.DisableColor {
		switch logLevel {
		case "VRB":
			color = reset
			break
		case "INF":
			color = blue
			break
		case "WRN":
			color = yellow
			break
		case "ERR":
			color = red
			break
		default:
			color = cyan
		}
	} else {
		color = reset
	}

	if len(logLevel) > 3 {
		logLevel = logLevel[:3]
	}

	if l.LogLevel <= 4 {
		golog.Printf("[%s] %v |%s %s %s| %s\n",
			l.logPrefix,
			time.Now().Format("2006/01/02 15:04:05"),
			color, logLevel, reset,
			logMessage,
		)
	}

}
