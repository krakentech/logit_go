package logit

import (
	"bytes"
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestSettings(t *testing.T) {
	setSettingsToDefault()
	newWriter := &bytes.Buffer{}

	t.Run("verify defaults", func(t *testing.T) {
		assert.Equal(t, os.Stdout, writer)
		assert.Equal(t, outputDefault, outFormat)
		assert.Equal(t, timeFormatDefault, timeFormat)
		assert.False(t, isDebug)
	})
	t.Run("verify set settings", func(t *testing.T) {
		SetWriter(newWriter)
		SetOutFormat("test out format")
		SetTimeFormat("test time format")
		SetIsDebug(true)
		assert.Equal(t, newWriter, writer)
		assert.Equal(t, "test out format", outFormat)
		assert.Equal(t, "test time format", timeFormat)
		assert.True(t, isDebug)
	})
	setSettingsToDefault()
}

func setSettingsToDefault() {
	isDebug = false
	outFormat = outputDefault
	timeFormat = timeFormatDefault
	writer = os.Stdout
	printLf = fmt.Fprintln
}
