package logit

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestTracker(t *testing.T) {
	setSettingsToDefault()
	buf := setupForLoggingTest()
	tracker := NewTracker("test-tracker")

	tests := []struct {
		name   string
		newNow func() time.Time
		want   string
	}{
		{
			name:   "1 ns",
			newNow: func() time.Time { return time.Date(2006, time.January, 2, 15, 4, 5, 1, time.UTC) },
			want:   "06.01.02-15:04:05 ⏰ - test-tracker [1 ns] 1ns\n",
		},
		{
			name:   "1 sec",
			newNow: func() time.Time { return time.Date(2006, time.January, 2, 15, 4, 6, 0, time.UTC) },
			want:   "06.01.02-15:04:06 ⏰ - test-tracker [1000000000 ns] 1s\n",
		},
		{
			name:   "1 min",
			newNow: func() time.Time { return time.Date(2006, time.January, 2, 15, 5, 5, 0, time.UTC) },
			want:   "06.01.02-15:05:05 ⏰ - test-tracker [60000000000 ns] 1m0s\n",
		},
		{
			name:   "1 hour",
			newNow: func() time.Time { return time.Date(2006, time.January, 2, 16, 4, 5, 0, time.UTC) },
			want:   "06.01.02-16:04:05 ⏰ - test-tracker [3600000000000 ns] 1h0m0s\n",
		},
		{
			name:   "1 day",
			newNow: func() time.Time { return time.Date(2006, time.January, 3, 15, 4, 5, 0, time.UTC) },
			want:   "06.01.03-15:04:05 ⏰ - test-tracker [86400000000000 ns] 24h0m0s\n",
		},
		{
			name:   "1 month",
			newNow: func() time.Time { return time.Date(2006, time.February, 2, 15, 4, 5, 0, time.UTC) },
			want:   "06.02.02-15:04:05 ⏰ - test-tracker [2678400000000000 ns] 744h0m0s\n",
		},
		{
			name:   "1 year",
			newNow: func() time.Time { return time.Date(2007, time.January, 2, 15, 4, 5, 0, time.UTC) },
			want:   "07.01.02-15:04:05 ⏰ - test-tracker [31536000000000000 ns] 8760h0m0s\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			now = tt.newNow
			tracker.Log()
			assert.Equal(t, tt.want, buf.String())
			buf.Reset()
		})
	}
}
