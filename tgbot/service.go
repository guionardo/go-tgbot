package tgbot

import (
	"context"
	"fmt"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type GoTGBotSevice struct {
	BotRunner
	context         context.Context
	listener        *BotListener
	worker          *BotWorker
	publisher       *BotPublisher
	cancel          context.CancelFunc
	stop            context.CancelFunc
	internalChannel chan InternalMessage
}

func CreateBotService(telegramToken string, injections ...interface{}) (service *GoTGBotSevice, err error) {
	var bot *tgbotapi.BotAPI
	bot, err = tgbotapi.NewBotAPI(telegramToken)
	if err != nil {
		return
	}

	logger := GetLogger("GoTGBotSevice")
	user, err := bot.GetMe()
	logger.Infof("Authorized on account %s", user.UserName)
	var listener *BotListener
	var worker *BotWorker
	var publisher *BotPublisher
	var schedules *ScheduleCollection

	for _, injection := range injections {
		switch injection.(type) {
		case *BotListener:
			listener = injection.(*BotListener)
		case *BotWorker:
			worker = injection.(*BotWorker)
		case *BotPublisher:
			publisher = injection.(*BotPublisher)
		case *ScheduleCollection:
			schedules = injection.(*ScheduleCollection)
		}

	}

	internalChannel := make(chan InternalMessage, 10)
	if listener == nil {
		listener = createBotListener(bot)
	}
	listener.SetInternalChannel(internalChannel)

	if schedules == nil {
		schedules = CreateScheduleCollection()
	}
	if worker == nil {
		worker = createBotWorker(bot, schedules)
	}
	worker.SetInternalChannel(internalChannel)

	if publisher == nil {
		publisher = CreateBotPublisher(bot)
	}
	publisher.SetInternalChannel(internalChannel)

	service = &GoTGBotSevice{
		listener:        listener,
		worker:          worker,
		publisher:       publisher,
		internalChannel: internalChannel,
	}

	service.Init(bot, "GoTGBotSevice")
	return service, nil
}

func (svc *GoTGBotSevice) Start() error {
	if svc.cancel != nil {
		return fmt.Errorf("service already started")
	}
	botContext, cancel := CreateBotContext(svc)

	svc.cancel = cancel

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		svc.listener.Run(botContext)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		svc.worker.Run(botContext)
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		svc.publisher.Run(botContext)
	}()
	svc.logger.Info("started")
	wg.Wait()
	return nil
}

func (svc *GoTGBotSevice) Stop() error {
	if svc.cancel == nil {
		return fmt.Errorf("service not started")
	}
	svc.logger.Info("stopping")
	svc.cancel()
	svc.logger.Info("stopped")
	return nil
}

func (svc *GoTGBotSevice) RunnersRunning() int {
	return svc.listener.IsRunningInt() + svc.worker.IsRunningInt() + svc.publisher.IsRunningInt()
}

func (svc *GoTGBotSevice) Publish(message tgbotapi.Chattable) {
	svc.publisher.Publish(message)
}

func (svc *GoTGBotSevice) AddSchedule(schedule *Schedule) {
	svc.worker.AddSchedule(schedule)
}
