package example

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/guionardo/go-tgbot/tgbot"
	"github.com/guionardo/go-tgbot/tgbot/automations"
	"github.com/guionardo/go-tgbot/tgbot/infra"
)

func getConfig() *tgbot.Configuration {
	config, err := tgbot.LoadConfiguration()
	logger := infra.GetLogger("simple_bot")
	if err != nil {
		logger.Panicf("Error loading configuration: %v", err)
	}

	return config
}

func Run() {
	logger := infra.GetLogger("simple_bot")
	logger.Info("Starting Simple Bot")
	config := getConfig()
	svc, err := tgbot.CreateBotService(config)
	if err != nil {
		logger.Panicf("Error creating bot service: %v", err)
	}
	automations.AddHelloWorldAutomation(svc)
	automations.AddSetupCommandsAutomation(svc)
	automations.AddHouseKeepingAutomation(svc)

	// svc.AddSchedule(schedules.CreateSchedule("Every minute", time.Minute, func(ctx context.Context) error {
	// 	svc := tgbot.GetBotService(ctx)
	// 	svc.Publish(tgbotapi.NewMessage(205478553, "PING MINUTE"))
	// 	return nil
	// }))
	// svc.AddSchedule(schedules.CreateSchedule("Every 10 seconds", time.Second*10, func(ctx context.Context) error {
	// 	svc := tgbot.GetBotService(ctx)
	// 	svc.Publish(tgbotapi.NewMessage(205478553, "PING SECONDS"))
	// 	return nil
	// }))

	svc.AddCommandHandler("Menu principal", "menu", func(ctx context.Context, u tgbotapi.Update) error {
		svc := tgbot.GetBotService(ctx)
		svc.Publisher().SendInlineKeyboard(u.Message.Chat.ID, "Menu", "Opção 1:TESTE", "Opção 2:TESTE2", "-", "Opção 3:TESTE3")
		return nil
	})

	svc.AddHandler("all", func(update tgbotapi.Update) bool {
		return true
	}, func(ctx context.Context, update tgbotapi.Update) error {
		logger.Infof("%s : %s", update.Message.From.UserName, update.Message.Text)
		svc := tgbot.GetBotService(ctx)
		msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, "+update.Message.From.UserName+"!")
		msg.ReplyToMessageID = update.Message.MessageID
		svc.Publish(msg)
		return nil
	})
	svc.Start()
}
