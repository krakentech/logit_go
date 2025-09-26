package logit

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
	"time"
)

func TestDebug(t *testing.T) {
	defer setSettingsToDefault()
	buff := setupForLoggingTest()

	tests := []struct {
		name    string
		isDebug bool
		message string
		values  []any
		want    string
	}{
		{
			name:    "Debug False",
			isDebug: false,
			message: "Test Debug Message",
			values:  []any{},
			want:    "",
		},
		{
			name:    "Debug Message",
			message: "This is a debug message",
			isDebug: true,
			values:  []any{},
			want:    "06.01.02-15:04:05 🐛 - This is a debug message\n",
		},
		{
			name:    "Debug Message and Values",
			message: "This is a debug message [%s, %d, %t]",
			isDebug: true,
			values:  []any{"A", 123, true},
			want:    "06.01.02-15:04:05 🐛 - This is a debug message [A, 123, true]\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetIsDebug(tt.isDebug)
			Debug(tt.message, tt.values...)
			assert.Equal(t, tt.want, buff.String())
			buff.Reset()
		})
	}

	// buf := &bytes.Buffer{}
	// SetWriter(buf)
	// now = func() time.Time { return time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC) }
	//
	// Debug("This is a debug message: %d", 42)
	// assert.Equal(t, "", buf.String())
	//
	// SetIsDebug(true)
	// Debug("This is a debug message: %d", 42)
	// assert.Equal(t, "06.01.02-15:04:05 🐛 - This is a debug message: 42\n", buf.String())
	//
	// setSettingsToDefault()
}

func TestInfo(t *testing.T) {
	defer setSettingsToDefault()
	buff := setupForLoggingTest()

	tests := []struct {
		name    string
		message string
		values  []any
		want    string
	}{
		{
			name:    "Info Message",
			message: "This is a info message",
			values:  []any{},
			want:    "06.01.02-15:04:05 🧠 - This is a info message\n",
		},
		{
			name:    "Info Message and Values",
			message: "This is a info message [%s, %d, %t]",
			values:  []any{"A", 123, true},
			want:    "06.01.02-15:04:05 🧠 - This is a info message [A, 123, true]\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Info(tt.message, tt.values...)
			assert.Equal(t, tt.want, buff.String())
			buff.Reset()
		})
	}
}

func TestWarn(t *testing.T) {
	defer setSettingsToDefault()
	buff := setupForLoggingTest()

	tests := []struct {
		name    string
		message string
		values  []any
		want    string
	}{
		{
			name:    "Warn Message",
			message: "This is a warn message",
			values:  []any{},
			want:    "06.01.02-15:04:05 ⚠️ - This is a warn message\n",
		},
		{
			name:    "Warn Message and Values",
			message: "This is a warn message [%s, %d, %t]",
			values:  []any{"A", 123, true},
			want:    "06.01.02-15:04:05 ⚠️ - This is a warn message [A, 123, true]\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Warn(tt.message, tt.values...)
			assert.Equal(t, tt.want, buff.String())
			buff.Reset()
		})
	}
}

func TestError(t *testing.T) {
	defer setSettingsToDefault()
	buff := setupForLoggingTest()

	tests := []struct {
		name    string
		message string
		values  []any
		want    string
	}{
		{
			name:    "Error Message",
			message: "This is a error message",
			values:  []any{},
			want:    "06.01.02-15:04:05 💥 - This is a error message\n",
		},
		{
			name:    "Error Message and Values",
			message: "This is a error message [%s, %d, %t]",
			values:  []any{"A", 123, true},
			want:    "06.01.02-15:04:05 💥 - This is a error message [A, 123, true]\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Error(tt.message, tt.values...)
			assert.Equal(t, tt.want, buff.String())
			buff.Reset()
		})
	}
}

func TestErr(t *testing.T) {
	defer setSettingsToDefault()
	buff := setupForLoggingTest()

	tests := []struct {
		name    string
		message string
		values  []any
		err     error
		want    string
	}{
		{
			name:    "Err Message",
			message: "This is a error message",
			values:  []any{},
			err:     errors.New("test error"),
			want:    "06.01.02-15:04:05 💥 - This is a error message: test error\n",
		},
		{
			name:    "Err Message and Values",
			message: "This is a error message [%s, %d, %t]",
			values:  []any{"A", 123, true},
			err:     errors.New("test error"),
			want:    "06.01.02-15:04:05 💥 - This is a error message [A, 123, true]: test error\n",
		},
		{
			name:    "Err Message and Values No Error",
			message: "This is a error message [%s, %d, %t]",
			values:  []any{"A", 123, true},
			err:     nil,
			want:    "06.01.02-15:04:05 💥 - This is a error message [A, 123, true]: no error to log\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Err(tt.err, tt.message, tt.values...)
			assert.Equal(t, tt.want, buff.String())
			buff.Reset()
		})
	}
}

func TestDebugData(t *testing.T) {
	defer setSettingsToDefault()
	buff := setupForLoggingTest()

	tests := []struct {
		name     string
		isDebug  bool
		data     any
		formated bool
		message  string
		values   []any
		want     string
	}{
		{
			name:     "Debug False",
			isDebug:  false,
			data:     map[string]any{},
			formated: true,
			message:  "Test Debug Data Message",
			values:   []any{},
			want:     "",
		},
		{
			name:     "Debug Message [un-formated]",
			isDebug:  true,
			data:     []string{"A", "B", "C"},
			formated: false,
			message:  "Test Debug Data Message",
			values:   []any{},
			want:     "06.01.02-15:04:05 ✨ - --> Test Debug Data Message\n06.01.02-15:04:05 ✨ - [\"A\",\"B\",\"C\"]\n",
		},
		{
			name:     "Debug Message [formated]",
			isDebug:  true,
			data:     []string{"A", "B", "C"},
			formated: true,
			message:  "Test Debug Data Message",
			values:   []any{},
			want:     "06.01.02-15:04:05 ✨ - --> Test Debug Data Message\n06.01.02-15:04:05 ✨ - [\n06.01.02-15:04:05 ✨ - \t\"A\",\n06.01.02-15:04:05 ✨ - \t\"B\",\n06.01.02-15:04:05 ✨ - \t\"C\"\n06.01.02-15:04:05 ✨ - ]\n",
		},
		{
			name:     "Debug Bad Data",
			isDebug:  true,
			data:     make(chan int),
			formated: false,
			message:  "Test Debug Data Message",
			values:   []any{},
			want:     "06.01.02-15:04:05 💥 - failed to marshal data object: json: unsupported type: chan int\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetIsDebug(tt.isDebug)
			DebugData(tt.data, tt.formated, tt.message, tt.values...)
			assert.Equal(t, tt.want, buff.String())
			buff.Reset()
		})
	}

}

func TestPrintLine(t *testing.T) {
	defer setSettingsToDefault()
	buff := setupForLoggingTest()

	tests := []struct {
		name    string
		printLf func(w io.Writer, a ...any) (n int, err error)
		logType logType
		format  string
		message string
		want    string
	}{
		{
			name:    "Basic",
			printLf: fmt.Fprintln,
			logType: logTypeInfo,
			format:  outputDefault,
			message: "This is a debug message",
			want:    "06.01.02-15:04:05 🧠 - This is a debug message\n",
		},
		{
			name:    "Basic No Time",
			printLf: fmt.Fprintln,
			logType: logTypeInfo,
			format:  "{{type}} - {{message}}",
			message: "This is a debug message",
			want:    "🧠 - This is a debug message\n",
		},
		{
			name:    "Basic No Type",
			printLf: fmt.Fprintln,
			logType: logTypeInfo,
			format:  "{{message}}",
			message: "This is a debug message",
			want:    "This is a debug message\n",
		},
		{
			name:    "Basic Nothing",
			printLf: fmt.Fprintln,
			logType: logTypeInfo,
			format:  "NO MESSAGE",
			message: "This is a debug message",
			want:    "NO MESSAGE\n",
		},
		{
			name: "ERROR",
			printLf: func(w io.Writer, a ...any) (n int, err error) {
				return 0, fmt.Errorf("ERROR")
			},
			logType: logTypeInfo,
			format:  "NO MESSAGE",
			message: "This is a debug message",
			want:    "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printLf = tt.printLf
			SetOutFormat(tt.format)

			printLine(tt.logType, tt.message)
			assert.Equal(t, tt.want, buff.String())

			buff.Reset()
		})
	}
}

func setupForLoggingTest() *bytes.Buffer {
	setSettingsToDefault()
	buf := &bytes.Buffer{}
	SetWriter(buf)
	now = func() time.Time { return time.Date(2006, time.January, 2, 15, 4, 5, 0, time.UTC) }
	return buf
}
