package logit

import (
	"fmt"
	"time"
)

type Tracker struct {
	title string
	start time.Time
}

func NewTracker(title string) *Tracker {
	start := now()
	return &Tracker{
		title: title,
		start: start,
	}
}

func (p *Tracker) Log() {
	diff := now().Sub(p.start)
	printLine(logTypePerformance, fmt.Sprintf("%s [%d ns] %s", p.title, diff.Nanoseconds(), diff.String()))
}
