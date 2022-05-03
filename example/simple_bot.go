package example

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/guionardo/go-tgbot/tgbot"
)

func getAPIKey() string {
	config, err := tgbot.LoadConfiguration()
	logger := tgbot.GetLogger("simple_bot")
	if err != nil {
		logger.Panicf("Error loading configuration: %v", err)
	}

	return config.BotToken
}

func Run() {
	logger := tgbot.GetLogger("simple_bot")
	logger.Info("Starting Simple Bot")
	apiKey := getAPIKey()
	svc, err := tgbot.CreateBotService(apiKey)
	if err != nil {
		logger.Panicf("Error creating bot service: %v", err)
	}

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
