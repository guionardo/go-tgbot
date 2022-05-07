package automations

import (
	"context"
	"fmt"
	"time"

	"github.com/guionardo/go-tgbot/pkg/schedules"
	"github.com/guionardo/go-tgbot/tgbot"
)

func AddHelloWorldAutomation(svc *tgbot.GoTGBotSevice) *tgbot.GoTGBotSevice {
	svc.AddScheduleRunOnce(schedules.CreateSchedule("Chats greetings", time.Hour, func(ctx context.Context) error {
		getLogger().Infof("Running Hello World")
		svc := tgbot.GetBotService(ctx)
		chats, err := svc.Repository().GetChats()
		if err != nil {
			return err
		}
		for _, chat := range chats {
			svc.Publisher().SendTextMessage(int64(chat.ID), fmt.Sprintf("🤖 %s - %s", svc.Configuration().BotName, svc.Configuration().BotHelloWorld))
		}
		return nil
	}))
	getLogger().Infof("Added Hello World Automation")
	return svc
}
