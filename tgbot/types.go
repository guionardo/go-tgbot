package tgbot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	MessageHandler = func(context.Context, tgbotapi.Update) error
	MessageFilter  = func(tgbotapi.Update) bool
)
