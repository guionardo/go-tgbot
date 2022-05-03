package tgbot

import (
	"context"
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

func (pbx *BotPublisher) Run(ctx context.Context) {
	defer pbx.Stop()
	pbx.Start()
	stopped := make(map[string]bool)
	for {
		select {
		case <-ctx.Done():
			pbx.logger.Info("context done")
			return
		case messages := <-pbx.publishChannel:
			waitTime := time.Now().Sub(pbx.lastSendTime)
			if waitTime < time.Second {
				time.Sleep(time.Second - waitTime)
			}
			_, err := pbx.bot.Send(messages)
			if err != nil {
				pbx.logger.Errorf("error sending messages: %v", err)
			}
		case message := <-pbx.internalChannel:
			if message.message == "stop" {
        svc:=GetBotService(ctx)
        if svc.RunnersRunning()==1 {
					pbx.logger.Warningf("all stopped - stopping - %v", stopped)
          // svc.Stop()
				}
			}
		}
	}
}
