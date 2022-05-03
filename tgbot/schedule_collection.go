package tgbot

type ScheduleCollection struct {
	schedules []*Schedule
}

func CreateScheduleCollection() *ScheduleCollection {
	return &ScheduleCollection{
		schedules: make([]*Schedule, 0),
	}
}

func (sch *ScheduleCollection) AddSchedule(schedule *Schedule) *ScheduleCollection {
	sch.schedules = append(sch.schedules, schedule)
	return sch
}

func (sch *ScheduleCollection) GetNextSchedule() (nextSchedule *Schedule) {
	for _, schedule := range sch.schedules {
		if nextSchedule == nil || schedule.nextRun.Before(nextSchedule.nextRun) {
			nextSchedule = schedule
			break
		}
	}
	return
}

func (sch *ScheduleCollection) Count() int {
	return len(sch.schedules)
}
