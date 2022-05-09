package helpers

import (
	"net/url"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// CreateTextMessage creates a new message with the given chatID and text.
func CreateTextMessage(chatID int64, text string) tgbotapi.MessageConfig {
	return tgbotapi.NewMessage(chatID, text)
}

func CreateHTMLMessage(chatID int64, htmlMessage string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(chatID, htmlMessage)
	msg.ParseMode = "HTML"
	return msg
}

func CreateKeyboardMessage(chatID int64, text string, keyboardOptions ...string) tgbotapi.MessageConfig {
	if len(keyboardOptions) == 0 {
		return CreateTextMessage(chatID, "[no keyboard options]")
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
		if IsValidUrl(value) {
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
	return msg
}

func IsValidUrl(toTest string) bool {
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
