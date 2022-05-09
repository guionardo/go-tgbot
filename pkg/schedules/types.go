package schedules

import (
	"context"
	"time"
)

type (
	Schedule struct {
		Title        string
		lastRun      time.Time
		nextRun      time.Time
		Interval     time.Duration
		action       ScheduledAction
		lastWasRound bool
	}
	ScheduledAction    func(ctx context.Context) error
	ScheduleCollection struct {
		schedules []*Schedule
		runOnce   []*Schedule
	}
)
