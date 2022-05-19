package example

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/guionardo/go-tgbot/tgbot"
	"github.com/guionardo/go-tgbot/tgbot/automations"
)

func Run() {

	// Create bot runner from environment variables
	svc := tgbot.CreateBotService().
		LoadConfigurationFromEnv("TG_").
		InitBot()

	// Setup automations
	automations.AddStartupGreetingsAutomation(svc)
	automations.AddSetupCommandsAutomation(svc)
	automations.AddHouseKeepingAutomation(svc)

	// svc.AddSchedule(schedules.CreateSchedule("Every minute", time.Minute, func(ctx context.Context) error {
	// 	svc := tgbot.GetBotService(ctx)
	// 	svc.Publish(tgbotapi.NewMessage(205478553, "PING MINUTE"))
	// 	return nil
	// }))
	// svc.AddSchedule(schedules.CreateSchedule("Every 10 seconds", time.Second*10, func(ctx context.Context) error {
	// 	svc := tgbot.GetBotService(ctx)
	// 	svc.Publish(tgbotapi.NewMessage(205478553, "PING SECONDS"))
	// 	return nil
	// }))

	svc.AddCommandHandlers(&tgbot.ListenerCommandHandler{
		Command: "menu",
		Title:   "Menu principal",
		Func: func(ctx context.Context, u tgbotapi.Update) error {
			svc := tgbot.GetBotService(ctx)
			svc.Publisher().SendMenuKeyboard(u.Message.Chat.ID, "Menu principal", "menu", "Opção 1:OPCAO1", "Opção 2:OPCAO2", "-", "Opção 3:OPCAO3", "Opção 4:OPCAO4")

			return nil
		}}, &tgbot.ListenerCommandHandler{
		Command: "hello",
		Title:   "Dizer olá",
		Func: func(ctx context.Context, u tgbotapi.Update) error {
			svc := tgbot.GetBotService(ctx)
			svc.Publisher().ReplyToMessage(u, fmt.Sprintf("Olá %s", u.Message.From.UserName))
			return nil
		}},
	)

	svc.AddCallbackHandlers(tgbot.CreateListenerCallbackHandler("menu", "menu", func(ctx context.Context, u tgbotapi.Update) error {
		svc := tgbot.GetBotService(ctx)
		svc.Publisher().SendHTMLMessage(u.CallbackQuery.Message.Chat.ID, fmt.Sprintf("MENU: <b>%s</b>", u.CallbackQuery.Data))
		return nil
	}))

	svc.AddHandlers(&tgbot.ListenerFilteredHandler{
		Title:  "all",
		Filter: func(update tgbotapi.Update) bool { return true },
		Func: func(ctx context.Context, update tgbotapi.Update) error {
			svc := tgbot.GetBotService(ctx)
			msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Hello, "+update.Message.From.UserName+"!")
			msg.ReplyToMessageID = update.Message.MessageID
			svc.Publisher().Publish(msg)
			return nil
		}})

	svc.Start()
}
