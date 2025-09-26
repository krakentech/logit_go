package logit

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	// version specifies the current version of the log-it library.
	version = "v0.0.0"

	// outputDefault defines the default format for log messages.
	outputDefault = "{{time}} {{type}} - {{message}}"

	// timeFormatDefault specifies the default time format for log messages.
	timeFormatDefault = "06.01.02-15:04:05"
)

var (
	isDebug    bool
	outFormat  string    = outputDefault
	timeFormat string    = timeFormatDefault
	writer     io.Writer = os.Stdout
	printLf              = fmt.Fprintln
	now                  = time.Now
)

// SetWriter will update where the logging is written
func SetWriter(newWriter io.Writer) {
	writer = newWriter
}

// SetOutFormat will change the format at which things are logged ( please check the readme for better understanding )
func SetOutFormat(format string) {
	outFormat = format
}

// SetTimeFormat changes the time format used when logging a message
func SetTimeFormat(format string) {
	timeFormat = format
}

// SetIsDebug will set log-it to debug mode which is the only way to print debug messages
func SetIsDebug(debugMode bool) {
	isDebug = debugMode
}
