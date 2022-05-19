package helpers

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func UpdateToString(update tgbotapi.Update) string {
	message := ""
	if update.Message != nil && update.Message.From != nil {
		message = update.Message.From.UserName + " : " + update.Message.Text
	} else {
		message = "no message"
	}

	return message
}
