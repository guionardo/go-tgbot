package tgbot

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	sch "github.com/guionardo/go-tgbot/pkg/schedules"
	"github.com/guionardo/go-tgbot/tgbot/infra"
	"github.com/guionardo/go-tgbot/tgbot/runners"
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
	repository      IRepository
	configuration   *Configuration
}

func CreateBotService(config *Configuration, injections ...interface{}) (service *GoTGBotSevice, err error) {
	var bot *tgbotapi.BotAPI
	bot, err = tgbotapi.NewBotAPI(config.BotToken)
	if err != nil {
		return
	}

	logger := infra.GetLogger("GoTGBotSevice")
	user, err := bot.GetMe()
	logger.Infof("Authorized on account %s", user.UserName)
	var listener *BotListener
	var worker *BotWorker
	var publisher *BotPublisher
	var schedules *sch.ScheduleCollection

	for _, injection := range injections {
		switch injection.(type) {
		case *BotListener:
			listener = injection.(*BotListener)
		case *BotWorker:
			worker = injection.(*BotWorker)
		case *BotPublisher:
			publisher = injection.(*BotPublisher)
		case *sch.ScheduleCollection:
			schedules = injection.(*sch.ScheduleCollection)
		}

	}

	internalChannel := make(chan InternalMessage, 10)
	if listener == nil {
		listener = createBotListener(bot)
	}
	listener.SetInternalChannel(internalChannel)

	if schedules == nil {
		schedules = sch.CreateScheduleCollection()
	}
	if worker == nil {
		worker = createBotWorker(bot, schedules)
	}
	worker.SetInternalChannel(internalChannel)

	if publisher == nil {
		publisher = CreateBotPublisher(bot)
	}
	publisher.SetInternalChannel(internalChannel)

	db, err := infra.GetSQLiteDB(config.RepositoryConnectionString)
	if err != nil {
		panic(err)
	}
	service = &GoTGBotSevice{
		listener:        listener,
		worker:          worker,
		publisher:       publisher,
		internalChannel: internalChannel,
		repository:      infra.CreateMessageRepository(db),
		configuration:   config,
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

	runners := runners.NewRunnerCollection()
	runners.CreateRunnerCustomLoop("listener", BotListenerAction, svc)
	runners.CreateRunnerCustomLoop("publisher", BotPublisherAction, svc)
	runners.CreateRunner("worker", BotWorkerRunnerAction, svc)
	runners.RunAll(botContext)

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

func (svc *GoTGBotSevice) Publish(message tgbotapi.Chattable) {
	svc.publisher.Publish(message)
}

func (svc *GoTGBotSevice) Publisher() *BotPublisher {
	return svc.publisher
}

func (svc *GoTGBotSevice) Repository() IRepository {
	return svc.repository
}

func (svc *GoTGBotSevice) Listener() *BotListener {
	return svc.listener
}

func (svc *GoTGBotSevice) Configuration() *Configuration {
	return svc.configuration
}
