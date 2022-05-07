package tgbot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	sch "github.com/guionardo/go-tgbot/pkg/schedules"
)

type BotWorker struct {
	BotRunner
	schedules *sch.ScheduleCollection
}
type BotWorkerAction func(ctx context.Context) error

func createBotWorker(bot *tgbotapi.BotAPI, schedules *sch.ScheduleCollection) *BotWorker {
	worker := &BotWorker{
		schedules: schedules,
	}
	worker.Init(bot, "BotWorker")
	return worker
}

func (wrk *BotWorker) AddSchedule(schedule *sch.Schedule) {
	wrk.schedules.AddSchedule(schedule)
}
