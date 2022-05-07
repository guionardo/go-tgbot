package schedules

import (
	"context"
	"time"
)

type (
	Schedule struct {
		title        string
		lastRun      time.Time
		nextRun      time.Time
		interval     time.Duration
		action       ScheduledAction
		lastWasRound bool
	}
	ScheduledAction    func(ctx context.Context) error
	ScheduleCollection struct {
		schedules []*Schedule
		runOnce   []*Schedule
	}
)
