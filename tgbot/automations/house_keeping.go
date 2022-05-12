package automations

import (
	"context"
	"time"

	"github.com/guionardo/go-tgbot/pkg/schedules"
	"github.com/guionardo/go-tgbot/tgbot"
)

// AddHouseKeepingAutomation setups hourly house keeping (cleaning up old messages, etc)
func AddHouseKeepingAutomation(svc *tgbot.GoTGBotService) *tgbot.GoTGBotService {
	svc.SetupBackgroundSchedules(schedules.CreateSchedule("House keeping", time.Hour, func(ctx context.Context) error {
		getLogger().Info("Running House keep")
		svc := tgbot.GetBotService(ctx)
		err := svc.Repository().HouseKeeping(svc.Configuration().Repository.HouseKeepingMaxAge)

		return err
	}))	
	return svc
}
