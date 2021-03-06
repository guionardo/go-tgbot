package schedules

import (
	"context"
	"time"

	"github.com/guionardo/go-tgbot/tgbot/infra"
	"github.com/sirupsen/logrus"
)

var schLogger *logrus.Entry

func schGetLogger() *logrus.Entry {
	if schLogger == nil {
		schLogger = infra.GetLogger("schedule")
	}
	return schLogger
}

func CreateSchedule(title string, interval time.Duration, action ScheduledAction) *Schedule {
	return &Schedule{
		Title:    title,
		Interval: interval,
		action:   action,
	}

}

func (sch *Schedule) WaitUntilNextRun() {
	sch.lastWasRound = false
	sch.nextRun = sch.lastRun.Add(sch.Interval)
	sch.waitUntilNextRun()
}

func (sch *Schedule) WaitUntilNextRunRound() {
	sch.lastWasRound = true
	sch.nextRun = sch.lastRun.Add(sch.Interval).Round(sch.Interval)
	if sch.nextRun.Before(time.Now()) {
		sch.nextRun = sch.nextRun.Add(sch.Interval)
	}
	sch.waitUntilNextRun()
}

func (sch *Schedule) waitUntilNextRun() {
	waitTime := sch.nextRun.Sub(time.Now())
	if waitTime.Seconds() > 0 {
		schGetLogger().Infof("wait %v for schedule %s", waitTime, sch.Title)
		time.Sleep(waitTime)
	}
}
func (sch *Schedule) RunNow() *Schedule {
	sch.nextRun = time.Now()

	return sch
}

func (sch *Schedule) RoundNextRun() *Schedule {
	sch.nextRun = sch.lastRun.Add(sch.Interval).Round(sch.Interval)
	return sch
}

func (sch *Schedule) CanRun() bool {
	return sch.nextRun.Before(time.Now())
}

func (sch *Schedule) DoAction(ctx context.Context) {
	sch.action(ctx)
	sch.lastRun = time.Now()
	if sch.lastWasRound {
		sch.nextRun = sch.lastRun.Add(sch.Interval).Round(sch.Interval)
	} else {
		sch.nextRun = sch.lastRun.Add(sch.Interval)
	}
}
