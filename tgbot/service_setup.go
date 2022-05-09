package tgbot

import "github.com/guionardo/go-tgbot/pkg/schedules"

func (svc *GoTGBotService) AddHandler(title string,
	filter MessageFilter,
	handler MessageHandler) *GoTGBotService {
	svc.listener.AddHandler(title, filter, handler)
	return svc
}

func (svc *GoTGBotService) AddSchedule(schedule *schedules.Schedule) {
	svc.logger.Infof("Added schedule %s every %v", schedule.Title, schedule.Interval)
	svc.worker.schedules.AddSchedule(schedule)
}

func (svc *GoTGBotService) AddScheduleRunOnce(schedule *schedules.Schedule) {
	svc.logger.Infof("Added schedule (run once) %s", schedule.Title)
	svc.worker.schedules.AddRunOnceSchedule(schedule)
}

func (svc *GoTGBotService) AddRepository(repository IRepository) {
	svc.repository = repository
}

func (svc *GoTGBotService) AddCommandHandler(title string, command string, handler MessageHandler) *GoTGBotService {
	svc.listener.AddCommandHandler(title, command, handler)
	return svc
}
