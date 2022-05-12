package tgbot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/guionardo/go-tgbot/tgbot/helpers"
)

func (pbx *BotPublisher) SendTextMessage(chatID int64, text string) {
	pbx.Publish(helpers.CreateTextMessage(chatID, text))
}

func (pbx *BotPublisher) SendHTMLMessage(chatID int64, htmlMessage string) {
	pbx.Publish(helpers.CreateHTMLMessage(chatID, htmlMessage))
}

func (pbx *BotPublisher) SendInlineKeyboard(chatID int64, text string, keyboardOptions ...string) {
	pbx.Publish(helpers.CreateKeyboardMessage(chatID, text, keyboardOptions...))
}

func (pbx *BotPublisher) ReplyToMessage(update tgbotapi.Update, text string) {
	pbx.Publish(helpers.CreateReplyMessage(update, text))
}
