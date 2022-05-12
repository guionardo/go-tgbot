package tgbot

import "github.com/guionardo/go-tgbot/pkg/schedules"

func (svc *GoTGBotService) SetupBackgroundSchedules(schedules ...*schedules.Schedule) *GoTGBotService {
	for _, schedule := range schedules {
		svc.worker.schedules.AddSchedule(schedule)
		svc.logger.Infof("Added schedule %s every %v", schedule.Title, schedule.Interval)
	}
	return svc
}

func (svc *GoTGBotService) SetupBackgroundRunOnceSchedules(schedules ...*schedules.Schedule) *GoTGBotService {
	for _, schedule := range schedules {
		svc.worker.schedules.AddRunOnceSchedule(schedule)
		svc.logger.Infof("Added schedule (run once) %s", schedule.Title)
	}
	return svc
}
