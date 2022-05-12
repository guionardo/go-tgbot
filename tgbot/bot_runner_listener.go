package tgbot

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/guionardo/go-tgbot/tgbot/runners"
)

func BotListenerAction(ctx context.Context, runner *runners.Runner) error {
	svc := GetBotService(ctx)
	lst := svc.listener
	if len(lst.handlers) == 0 {
		runner.Stop()
		return fmt.Errorf("no handlers for listener")
	}
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := svc.bot.GetUpdatesChan(u)
	for {
		select {
		case <-ctx.Done():
			lst.logger.Info("context done")
			return nil
		case <-runner.StopChannel:
			runner.Logger.Info("stopped by channel")
			return nil
		case update := <-updates:
			var handlerFunc ListenerHandlerFunc
			handlerTitle := "UNHANDLED"
			if update.Message != nil {
				svc.repository.SaveChat(update.Message.Chat)
				svc.repository.SaveMessage(update.Message)
				if update.Message.IsCommand() {
					if commandHandler, ok := lst.commands[update.Message.Command()]; ok {
						handlerFunc = commandHandler.Func
						handlerTitle = commandHandler.Title
					} else {
						lst.logger.Warningf("no command handler for %v", update.Message.Command())
					}
				} else {
					for _, messageHandler := range lst.handlers {
						if messageHandler.Filter(update) {
							handlerFunc = messageHandler.Func
							handlerTitle = messageHandler.Title
							break
						}
					}
				}
			} else if update.CallbackQuery != nil {
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
				if _, err := svc.bot.Request(callback); err != nil {
					lst.logger.Errorf("cannot get request - %v", err)
				}

				// And finally, send a message containing the data received.
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
				svc.publisher.Publish(msg)
				continue
			}
			if handlerFunc == nil {
				lst.logger.Warningf("[UNHANDLED] %s : %s", update.Message.From.UserName, update.Message.Text)
				continue
			}
			err := handlerFunc(ctx, update)
			if err == nil {
				lst.logger.Infof("[%s] %s : %s", handlerTitle, update.Message.From.UserName, update.Message.Text)
			} else {
				lst.logger.Errorf("[%s] %v - %s : %s", handlerTitle, err, update.Message.From.UserName, update.Message.Text)
			}
		}
	}
}
