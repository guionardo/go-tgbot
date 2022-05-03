package tgbot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotWorker struct {
	BotRunner
	schedules IScheduleCollection
}
type BotWorkerAction func(ctx context.Context) error

func createBotWorker(bot *tgbotapi.BotAPI, schedules IScheduleCollection) *BotWorker {
	worker := &BotWorker{
		schedules: schedules,
	}
	worker.Init(bot, "BotWorker")
	return worker
}

func (wrk *BotWorker) AddSchedule(schedule *Schedule) {
	wrk.schedules.AddSchedule(schedule)
}

func (wrk *BotWorker) Run(ctx context.Context) {
	defer wrk.Stop()
	wrk.Start()
	if wrk.schedules.Count() == 0 {
		wrk.logger.Info("no schedules to run")
		return
	}

	for {
		select {
		case <-ctx.Done():
			wrk.logger.Info("context done")
			return
		default:
			nextSchedule := wrk.schedules.GetNextSchedule()
      nextSchedule.WaitUntilNextRunRound()
			nextSchedule.DoAction(ctx)
		}
	}
}
