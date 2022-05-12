package tgbot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type (
	ListenerFilteredHandler struct {
		Command string
		Title   string
		Func    ListenerHandlerFunc
		Filter  ListenerFilter
	}

	ListenerCommandHandler struct {
		Command string
		Title   string
		Func    ListenerHandlerFunc
	}

	// ListenerHandlerFunc is the handler for messages received by the bot.
	ListenerHandlerFunc = func(context.Context, tgbotapi.Update) error

	// ListenerFilter is the filter for messages received by the bot. If the filter returns true, the message will be handled by the handler.
	ListenerFilter = func(tgbotapi.Update) bool
)
