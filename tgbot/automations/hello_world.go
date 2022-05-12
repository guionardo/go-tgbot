package automations

import (
	"context"
	"fmt"
	"time"

	"github.com/guionardo/go-tgbot/pkg/schedules"
	"github.com/guionardo/go-tgbot/tgbot"
)

// AddStartupGreetingsAutomation On startup, send a greeting message to the chats
func AddStartupGreetingsAutomation(svc *tgbot.GoTGBotService) *tgbot.GoTGBotService {
	svc.SetupBackgroundRunOnceSchedules(schedules.CreateSchedule("Chats greetings", time.Hour, func(ctx context.Context) error {
		getLogger().Infof("Running Hello World")
		svc := tgbot.GetBotService(ctx)
		chats, err := svc.Repository().GetChats()
		if err != nil {
			return err
		}
		for _, chat := range chats {
			svc.Publisher().SendTextMessage(int64(chat.ID), fmt.Sprintf("🤖 %s - %s", svc.Configuration().Bot.Name, svc.Configuration().Bot.HelloWorld))
		}
		return nil
	}))	
	return svc
}
