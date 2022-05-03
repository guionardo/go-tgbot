package tgbot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type IContextRunner interface {
	Run(ctx context.Context)
	IsRunning() bool
	GetName() string
}

type IScheduleCollection interface {
	AddSchedule(schedule *Schedule) *ScheduleCollection
	GetNextSchedule() (nextSchedule *Schedule)
	Count() int
}

type IBotPublisher interface {
	IContextRunner
	Publish(messages tgbotapi.Chattable)
}
