package tgbot

import (
	"context"

	sch "github.com/guionardo/go-tgbot/pkg/schedules"
)

type BotWorker struct {
	BotRunner
	schedules *sch.ScheduleCollection
}
type BotWorkerAction func(ctx context.Context) error

func createBotWorker() *BotWorker {
	worker := &BotWorker{
		schedules: sch.CreateScheduleCollection(),
	}
	worker.Init("BotWorker")
	return worker
}

func (wrk *BotWorker) AddSchedule(schedule *sch.Schedule) {
	wrk.schedules.AddSchedule(schedule)
}
