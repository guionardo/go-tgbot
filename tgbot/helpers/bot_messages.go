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

// getMenuOption -> caption:command|value
func GetMenuOption(keyboardOption string) (command string, caption string, value string, isLineBreak bool) {
	if strings.TrimSpace(keyboardOption) == "-" {
		isLineBreak = true
		return
	}

	caption, value, found := strings.Cut(keyboardOption, ":")
	if !found {
		caption = keyboardOption
		value = keyboardOption
	}
	command, newValue, found := strings.Cut(value, "|")
	if !found {
		command = ""
	} else {
		value = newValue
	}

	return
}

func CreateKeyboardMessageWithCommand(chatID int64, caption string, baseCommand string, keyboardOptions ...string) tgbotapi.MessageConfig {
	if len(keyboardOptions) == 0 {
		return CreateTextMessage(chatID, "[no keyboard options]")
	}

	for index, option := range keyboardOptions {
		menuOption := ParseBotMenuOption(option)

		if menuOption.IsLineBreak {
			continue
		}
		menuOption.Command = baseCommand

		keyboardOptions[index] = menuOption.String()

	}
	return CreateKeyboardMessage(chatID, caption, keyboardOptions...)
}

func CreateKeyboardMessage(chatID int64, text string, keyboardOptions ...string) tgbotapi.MessageConfig {
	if len(keyboardOptions) == 0 {
		return CreateTextMessage(chatID, "[no keyboard options]")
	}
	tgbotapi.NewInlineKeyboardRow()
	keyboard := tgbotapi.InlineKeyboardMarkup{}
	currentRow := make([]tgbotapi.InlineKeyboardButton, 0)
	for _, option := range keyboardOptions {
		menuOption := ParseBotMenuOption(option)

		if menuOption.IsLineBreak {
			if len(currentRow) > 0 {
				keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, currentRow)
			}
			currentRow = make([]tgbotapi.InlineKeyboardButton, 0)
			continue
		}

		if IsValidUrl(menuOption.Value) {
			currentRow = append(currentRow, tgbotapi.NewInlineKeyboardButtonURL(menuOption.Caption, menuOption.Value))
		} else {
			currentRow = append(currentRow, tgbotapi.NewInlineKeyboardButtonData(menuOption.Caption, menuOption.MessageValue()))
		}
	}
	if len(currentRow) > 0 {
		keyboard.InlineKeyboard = append(keyboard.InlineKeyboard, currentRow)
	}
	msg := tgbotapi.NewMessage(chatID, text)
	msg.ReplyMarkup = keyboard
	return msg
}

func CreateReplyMessage(update tgbotapi.Update, text string) tgbotapi.MessageConfig {
	msg := tgbotapi.NewMessage(update.Message.Chat.ID, text)
	msg.ReplyToMessageID = update.Message.MessageID
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
