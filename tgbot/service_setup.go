package tgbot

import (
	"context"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (svc *GoTGBotSevice) AddHandler(title string,
  filter func(update tgbotapi.Update) bool,
  handler func(ctx context.Context, update tgbotapi.Update) error) *GoTGBotSevice {
	svc.listener.AddHandler(title, filter, handler)
	return svc
}
