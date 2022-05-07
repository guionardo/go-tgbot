package tgbot

import (
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type BotPublisher struct {
	BotRunner
	publishChannel chan tgbotapi.Chattable
	lastSendTime   time.Time
}

func CreateBotPublisher(bot *tgbotapi.BotAPI) *BotPublisher {
	publisher := &BotPublisher{
		publishChannel: make(chan tgbotapi.Chattable, 10),
		lastSendTime:   time.Now(),
	}
	publisher.Init(bot, "BotPublisher")
	return publisher
}

func (pbx *BotPublisher) Publish(message tgbotapi.Chattable) {
	pbx.publishChannel <- message
}
