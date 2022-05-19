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

// SendMenuKeyboard sends a menu keyboard to the given chatID.
// caption is the caption of the menu.
// baseCommand is the command handler name that will be called when the user selects an option.
// keyboardOptions is a list of options. Each option is in format "Title:data". You can use "-" to separate options in a row.
func (pbx *BotPublisher) SendMenuKeyboard(chatID int64, caption string, baseCommand string, keyboardOptions ...string) {
	pbx.Publish(helpers.CreateKeyboardMessageWithCommand(chatID, caption, baseCommand, keyboardOptions...))
}

func (pbx *BotPublisher) ReplyToMessage(update tgbotapi.Update, text string) {
	pbx.Publish(helpers.CreateReplyMessage(update, text))
}
