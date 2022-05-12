package tgbot

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotListener struct {
	BotRunner
	handlers map[string]*ListenerFilteredHandler
	commands map[string]*ListenerCommandHandler
}

func createBotListener(bot *tgbotapi.BotAPI) *BotListener {
	listener := &BotListener{
		handlers: make(map[string]*ListenerFilteredHandler),
		commands: make(map[string]*ListenerCommandHandler),
	}
	listener.Init("BotListener")
	return listener
}

func (lst *BotListener) AddHandler(title string, filter func(update tgbotapi.Update) bool, handler func(ctx context.Context, update tgbotapi.Update) error) {
	lst.handlers[title] = &ListenerFilteredHandler{
		Filter: filter,
		Func:   handler,
		Title:  title,
	}
}

func (lst *BotListener) AddCommandHandler(title string, command string, handler ListenerHandlerFunc) {
	lst.commands[command] = &ListenerCommandHandler{
		Command: command,
		Func:    handler,
		Title:   title,
	}
}

func (lst *BotListener) SetupCommandsMessage() (msg tgbotapi.SetMyCommandsConfig, err error) {
	if len(lst.commands) == 0 {
		err = fmt.Errorf("no commands")
		return
	}
	commands := make([]tgbotapi.BotCommand, len(lst.commands))
	index := 0
	for name, cmd := range lst.commands {
		commands[index] = tgbotapi.BotCommand{
			Command:     name,
			Description: cmd.Title,
		}
		index++
	}
	scope := tgbotapi.NewBotCommandScopeDefault()
	return tgbotapi.NewSetMyCommandsWithScope(scope, commands...), nil
}
