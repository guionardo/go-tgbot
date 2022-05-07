package automations

import (
	"context"
	"time"

	"github.com/guionardo/go-tgbot/pkg/schedules"
	"github.com/guionardo/go-tgbot/tgbot"
)

func AddHouseKeepingAutomation(svc *tgbot.GoTGBotSevice) *tgbot.GoTGBotSevice {
	svc.AddSchedule(schedules.CreateSchedule("House keeping", time.Hour, func(ctx context.Context) error {
		getLogger().Info("Running House keep")
		svc := tgbot.GetBotService(ctx)
		err := svc.Repository().HouseKeeping(svc.Configuration().HouseKeepingMaxAge)

		return err
	}))
	getLogger().Infof("Added Housekeeping Automation")
	return svc
}
