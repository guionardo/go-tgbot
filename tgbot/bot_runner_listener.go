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

	updates := lst.bot.GetUpdatesChan(u)
	for {
		select {
		case <-ctx.Done():
			lst.logger.Info("context done")
			return nil
		case <-runner.StopChannel:
			runner.Logger.Info("stopped by channel")
			return nil
		case update := <-updates:
			var handler *BotListenerHandler
			if update.Message != nil {
				svc.repository.SaveChat(update.Message.Chat)
				svc.repository.SaveMessage(update.Message)
				if update.Message.IsCommand() {
					if commandHandler, ok := lst.commands[update.Message.Command()]; ok {
						handler = commandHandler
					} else {
						lst.logger.Warningf("no command handler for %v", update.Message.Command())
					}
				} else {
					for _, messageHandler := range lst.handlers {
						if messageHandler.filter(update) {
							handler = messageHandler
							break
						}
					}
				}
			} else if update.CallbackQuery != nil {
				callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
				if _, err := lst.bot.Request(callback); err != nil {
					lst.logger.Errorf("cannot get request - %v", err)
				}

				// And finally, send a message containing the data received.
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
				svc.publisher.Publish(msg)
				continue
			}
			if handler == nil {
				lst.logger.Warningf("[UNHANDLED] %s : %s", update.Message.From.UserName, update.Message.Text)
				continue
			}
			err := handler.handler(ctx, update)
			if err == nil {
				lst.logger.Infof("[%s] %s : %s", handler.title, update.Message.From.UserName, update.Message.Text)
			} else {
				lst.logger.Errorf("[%s] %v - %s : %s", handler.title, err, update.Message.From.UserName, update.Message.Text)
			}
		}
	}
}
