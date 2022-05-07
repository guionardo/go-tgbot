package tgbot

import (
	"net/url"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (pbx *BotPublisher) SendTextMessage(chatID int64, text string) {
	pbx.Publish(tgbotapi.NewMessage(chatID, text))
}

func (pbx *BotPublisher) SendHTMLMessage(chatID int64, htmlMessage string) {
	msg := tgbotapi.NewMessage(chatID, htmlMessage)
	msg.ParseMode = "HTML"
	pbx.Publish(msg)
}

func (pbx *BotPublisher) SendInlineKeyboard(chatID int64, text string, keyboardOptions ...string) {
	if len(keyboardOptions) == 0 {
		return
	}
	tgbotapi.NewInlineKeyboardRow()
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	currentRow := make([]tgbotapi.InlineKeyboardButton, 0)
	for _, option := range keyboardOptions {
		if option == "-" {
			if len(currentRow) > 0 {
				keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, currentRow)
			}
			currentRow = make([]tgbotapi.InlineKeyboardButton, 0)
			continue
		}
		title, value, found := strings.Cut(option, ":")
		if !found {
			title = option
			value = option
		}
		if isValidUrl(value) {
			currentRow = append(currentRow, tgbotapi.NewInlineKeyboardButtonURL(title, value))
		} else {
			currentRow = append(currentRow, tgbotapi.NewInlineKeyboardButtonData(title, value))
		}
	}
	if len(currentRow) > 0 {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, currentRow)
	}
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard
	pbx.Publish(msg)
}

func isValidUrl(toTest string) bool {
	_, err := url.ParseRequestURI(toTest)
	if err != nil {
		return false
	}

	u, err := url.Parse(toTest)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}

	return true
}
