package tgbot

import (
	"context"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/guionardo/go-tgbot/tgbot/runners"
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

func BotPublisherAction(ctx context.Context, runner *runners.Runner) error {
	svc := GetBotService(ctx)
	pbx := svc.Publisher()
	for {
		select {
		case <-ctx.Done():
			pbx.logger.Info("context done")
			return nil
		case <-runner.StopChannel:
			runner.Logger.Info("stopped by channel")
			return nil
		case msg := <-pbx.publishChannel:
			waitTime := time.Now().Sub(pbx.lastSendTime)
			if waitTime < time.Second {
				time.Sleep(time.Second - waitTime)
			}
			_, err := pbx.bot.Send(msg)
			if err != nil {
				pbx.logger.Errorf("error sending message: %v -  %v", msg, err)
			}
		}
	}
}
