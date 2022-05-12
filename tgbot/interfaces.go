package tgbot

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/guionardo/go-tgbot/tgbot/infra"
)


type IRepository interface {
	Save(*infra.Message) error
	GetChats() ([]*infra.Chat, error)
	SaveChat(chat *tgbotapi.Chat) error
	SaveMessage(message *tgbotapi.Message) error
	HouseKeeping(maxAge time.Duration) error
}
