package tgbot

import (
	"context"
	"strings"

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
		Command     string
		Title       string
		Func        ListenerHandlerFunc
		SubCommands []*ListenerCommandHandler
	}

	ListenerCallbackHandler struct {
		Title  string
		Filter ListenerFilter
		Func   ListenerHandlerFunc
	}

	// ListenerHandlerFunc is the handler for messages received by the bot.
	ListenerHandlerFunc = func(context.Context, tgbotapi.Update) error

	// ListenerFilter is the filter for messages received by the bot. If the filter returns true, the message will be handled by the handler.
	ListenerFilter = func(tgbotapi.Update) bool
)

func CreateListenerCallbackHandler(title string, command string, handlerFunc ListenerHandlerFunc) *ListenerCallbackHandler {
	return &ListenerCallbackHandler{
		Title: title,
		Filter: func(update tgbotapi.Update) bool {
			return update.CallbackQuery != nil && strings.HasPrefix(update.CallbackQuery.Data,command+"|")
		},
		Func: handlerFunc,
	}
}
