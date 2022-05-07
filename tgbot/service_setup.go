package tgbot

import "github.com/guionardo/go-tgbot/pkg/schedules"

func (svc *GoTGBotSevice) AddHandler(title string,
	filter MessageFilter,
	handler MessageHandler) *GoTGBotSevice {
	svc.listener.AddHandler(title, filter, handler)
	return svc
}

func (svc *GoTGBotSevice) AddSchedule(schedule *schedules.Schedule) {
	svc.worker.schedules.AddSchedule(schedule)
}

func (svc *GoTGBotSevice) AddScheduleRunOnce(schedule *schedules.Schedule) {
	svc.worker.schedules.AddRunOnceSchedule(schedule)
}

func (svc *GoTGBotSevice) AddRepository(repository IRepository) {
	svc.repository = repository
}

func (svc *GoTGBotSevice) AddCommandHandler(title string, command string, handler MessageHandler) *GoTGBotSevice {
	svc.listener.AddCommandHandler(title, command, handler)
	return svc
}
