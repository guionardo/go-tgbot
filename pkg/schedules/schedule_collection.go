package schedules

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
	if nextSchedule = sch.GetRunOnceSchedule(); nextSchedule != nil {
		return
	}
	for _, schedule := range sch.schedules {
		if nextSchedule == nil || schedule.nextRun.Before(nextSchedule.nextRun) {
			nextSchedule = schedule
		}
	}
	return
}

func (sch *ScheduleCollection) GetRunOnceSchedule() (nextSchedule *Schedule) {
	if len(sch.runOnce) == 0 {
		return nil
	}
	nextSchedule = sch.runOnce[0]
	sch.runOnce = sch.runOnce[1:]
	return
}

func (sch *ScheduleCollection) AddRunOnceSchedule(schedule *Schedule) *ScheduleCollection {
	sch.runOnce = append(sch.runOnce, schedule)
	return sch
}

func (sch *ScheduleCollection) Count() int {
	return len(sch.schedules) + len(sch.runOnce)
}
