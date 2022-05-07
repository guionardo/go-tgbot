package tgbot

import (
	"context"
	"time"

	"github.com/guionardo/go-tgbot/tgbot/runners"
)

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
		case messages := <-pbx.publishChannel:
			waitTime := time.Now().Sub(pbx.lastSendTime)
			if waitTime < time.Second {
				time.Sleep(time.Second - waitTime)
			}
			_, err := pbx.bot.Send(messages)
			if err != nil {
				pbx.logger.Errorf("error sending messages: %v", err)
			}
		}
	}
}
