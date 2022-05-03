package tgbot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotListener struct {
	BotRunner
	handlers map[string]*BotListenerHandler
}
type BotListenerHandler struct {
	filter  func(update tgbotapi.Update) bool
	handler func(ctx context.Context, update tgbotapi.Update) error
}

func createBotListener(bot *tgbotapi.BotAPI) *BotListener {
	listener := &BotListener{
		handlers: make(map[string]*BotListenerHandler),
	}
	listener.Init(bot, "BotListener")
	return listener
}

func (wrk *BotListener) Run(ctx context.Context) {
	defer wrk.Stop()
	if len(wrk.handlers) == 0 {
		wrk.logger.Warning("No handlers for listener")
		return
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := wrk.bot.GetUpdatesChan(u)
	for update := range updates {
		select {
		case <-ctx.Done():
			wrk.logger.Info("context done")
			return
		default:
			handled := false
			for title, handler := range wrk.handlers {
				if handler.filter(update) {
					err := handler.handler(ctx, update)
					if err == nil {
						wrk.logger.Infof("[%s] %s : %s", title, update.Message.From.UserName, update.Message.Text)
					} else {
						wrk.logger.Errorf("[%s] %v - %s : %s", title, err, update.Message.From.UserName, update.Message.Text)
					}
					handled = true
					break
				}
			}
			if !handled {
				wrk.logger.Warningf("[UNHANDLED] %s : %s", update.Message.From.UserName, update.Message.Text)
			}
		}
	}
}

func (wrk *BotListener) AddHandler(title string, filter func(update tgbotapi.Update) bool, handler func(ctx context.Context, update tgbotapi.Update) error) {
	wrk.handlers[title] = &BotListenerHandler{
		filter:  filter,
		handler: handler,
	}
}
