package automations

import (
	"context"
	"time"

	"github.com/guionardo/go-tgbot/pkg/schedules"
	"github.com/guionardo/go-tgbot/tgbot"
)

func AddSetupCommandsAutomation(svc *tgbot.GoTGBotSevice) *tgbot.GoTGBotSevice {
	svc.AddScheduleRunOnce(schedules.CreateSchedule("Setup commands", time.Hour, func(ctx context.Context) error {
		getLogger().Info("Running Setup commands")
		svc := tgbot.GetBotService(ctx)
		cmdMsg, err := svc.Listener().SetupCommandsMessage()
		if err != nil {
			return err
		}
		svc.Publisher().Publish(cmdMsg)
		return nil
	}))
	getLogger().Infof("Added Setup Commands Automation")
	return svc
}