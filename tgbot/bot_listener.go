package tgbot

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotListener struct {
	BotRunner
	handlers map[string]*BotListenerHandler
	commands map[string]*BotListenerHandler
}
type BotListenerHandler struct {
	title   string
	filter  MessageFilter
	handler MessageHandler
}

func createBotListener(bot *tgbotapi.BotAPI) *BotListener {
	listener := &BotListener{
		handlers: make(map[string]*BotListenerHandler),
		commands: make(map[string]*BotListenerHandler),
	}
	listener.Init(bot, "BotListener")
	return listener
}

func (lst *BotListener) AddHandler(title string, filter func(update tgbotapi.Update) bool, handler func(ctx context.Context, update tgbotapi.Update) error) {
	lst.handlers[title] = &BotListenerHandler{
		filter:  filter,
		handler: handler,
		title:   title,
	}
}

func (lst *BotListener) AddCommandHandler(title string, command string, handler MessageHandler) {
	lst.commands[command] = &BotListenerHandler{
		filter: func(update tgbotapi.Update) bool {
			return update.Message.IsCommand() && update.Message.Command() == command
		},
		handler: handler,
		title:   title,
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
			Description: cmd.title,
		}
		index++
	}
	return tgbotapi.NewSetMyCommands(commands...), nil
}
