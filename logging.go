package logit

import (
	"bufio"
	"encoding/json"
	"fmt"
	"strings"
)

type logType string

var (
	logTypeDebug       logType = "🐛"
	logTypeInfo        logType = "🧠"
	logTypeWarn        logType = "⚠️"
	logTypeError       logType = "💥"
	logTypeData        logType = "✨"
	logTypePerformance logType = "⏰"
)

func Debug(format string, a ...interface{}) {
	if isDebug {
		printLine(logTypeDebug, fmt.Sprintf(format, a...))
	}
}

func Info(format string, a ...interface{}) {
	printLine(logTypeInfo, fmt.Sprintf(format, a...))
}

func Warn(format string, a ...interface{}) {
	printLine(logTypeWarn, fmt.Sprintf(format, a...))
}

func Error(format string, a ...interface{}) {
	printLine(logTypeError, fmt.Sprintf(format, a...))
}

func Err(err error, format string, a ...interface{}) {
	msg := fmt.Sprintf(format, a...)
	if err == nil {
		err = fmt.Errorf("no error to log")
	}
	printLine(logTypeError, fmt.Sprintf("%s: %s", msg, err.Error()))
}

func DebugData(data any, formated bool, format string, a ...interface{}) {
	if isDebug {
		var b = make([]byte, 0)
		var err error = nil

		if formated {
			b, err = json.MarshalIndent(data, "", "\t")
		} else {
			b, err = json.Marshal(data)
		}

		if err != nil {
			Error("failed to marshal data object: %s", err.Error())
			return
		}

		printLine(logTypeData, fmt.Sprintf("--> %s", fmt.Sprintf(format, a...)))
		scanner := bufio.NewScanner(strings.NewReader(string(b)))
		for scanner.Scan() {
			printLine(logTypeData, scanner.Text())
		}
	}
}

func printLine(lType logType, msg string) {
	out := outFormat

	logTime := now().Format(timeFormat)

	out = strings.Replace(out, "{{time}}", logTime, 1)
	out = strings.Replace(out, "{{type}}", string(lType), 1)
	out = strings.Replace(out, "{{message}}", msg, 1)

	_, err := printLf(writer, out)
	if err != nil {
		fmt.Printf("failed to print line: %s", err.Error())
	}
}
