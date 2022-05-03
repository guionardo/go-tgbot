package tgbot

import (
	"context"
	"time"
)

type (
	Schedule struct {
		title    string
		lastRun  time.Time
		nextRun  time.Time
		interval time.Duration
		action   ScheduledAction
	}
	ScheduledAction func(ctx context.Context) error
)

func CreateSchedule(title string, interval time.Duration, action ScheduledAction) *Schedule {
	return &Schedule{
		title:    title,
		interval: interval,
		action:   action,
	}
}

func (sch *Schedule) RunNow() *Schedule {
	sch.nextRun = time.Now()
	return sch
}

func (sch *Schedule) RoundNextRun() *Schedule {
	sch.nextRun = time.Now().Round(sch.interval)
	return sch
}

func (sch *Schedule) CanRun() bool {
	return sch.nextRun.Before(time.Now())
}
