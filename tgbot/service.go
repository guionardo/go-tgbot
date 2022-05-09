package tgbot

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	sch "github.com/guionardo/go-tgbot/pkg/schedules"
	"github.com/guionardo/go-tgbot/tgbot/config"
	"github.com/guionardo/go-tgbot/tgbot/infra"
	"github.com/guionardo/go-tgbot/tgbot/runners"
)

type GoTGBotService struct {
	BotRunner
	context         context.Context
	listener        *BotListener
	worker          *BotWorker
	publisher       *BotPublisher
	cancel          context.CancelFunc
	stop            context.CancelFunc
	repository      IRepository
	configuration   *config.Configuration
}

// CreateBotService creates a new empty bot service
// This is the first step to create a bot service
func CreateBotService() *GoTGBotService {
	return &GoTGBotService{}
}

func (svc *GoTGBotService) LoadConfigurationFromFile(filename string) *GoTGBotService {
	cfg, err := config.CreateConfigurationFromFile(filename)
	if err != nil {
		panic(err)
	}
	svc.configuration = cfg
	return svc
}

func (svc *GoTGBotService) LoadConfigurationFromEnv(prefix string) *GoTGBotService {
	cfg, err := config.CreateConfigurationFromEnv(prefix)
	if err != nil {
		panic(err)
	}
	svc.configuration = cfg
	return svc
}

// InitBot setups bot, loggin
func (svc *GoTGBotService) InitBot() *GoTGBotService {
	if svc.configuration == nil {
		panic("configuration not loaded")
	}
	// Setup logging
	infra.CreateLoggerFactory(&svc.configuration.Logging)
	svc.logger = infra.GetLogger("GoTGBotService")
	svc.logger.Infof("Init %s", svc.name)

	// Setup bot
	bot, err := tgbotapi.NewBotAPI(svc.configuration.Bot.Token)
	if err != nil {
		panic(err)
	}
	svc.bot = bot
	svc.logger.Infof("Authorized on account %s", bot.Self.UserName)

	return svc
}

func (svc *GoTGBotService) AddWorkers() *GoTGBotService {

	// Create listener
	svc.listener = createBotListener(svc.bot)

	// Create schedules and worker
	schedules := sch.CreateScheduleCollection()
	svc.worker = createBotWorker(svc.bot, schedules)

	// Create publisher
	svc.publisher = CreateBotPublisher(svc.bot)

	// Create repository
	db, err := infra.GetSQLiteDB(svc.configuration.Repository.ConnectionString)
	if err != nil {
		panic(err)
	}
	svc.repository = infra.CreateMessageRepository(db)

	svc.Init(svc.bot, "GoTGBotService")
	return svc
}

func (svc *GoTGBotService) Start() error {
	if svc.cancel != nil {
		return fmt.Errorf("service already started")
	}
	botContext, cancel := CreateBotContext(svc)
	svc.cancel = cancel

	runnerCollection := runners.NewRunnerCollection()
	runnerCollection.CreateRunnerCustomLoop("listener", BotListenerAction, svc)
	runnerCollection.CreateRunnerCustomLoop("publisher", BotPublisherAction, svc)
	runnerCollection.CreateRunner("worker", BotWorkerRunnerAction, svc)
	runnerCollection.RunAll(botContext)

	return nil
}

func (svc *GoTGBotService) Stop() error {
	if svc.cancel == nil {
		return fmt.Errorf("service not started")
	}
	svc.logger.Info("stopping")
	svc.cancel()
	svc.logger.Info("stopped")
	return nil
}

func (svc *GoTGBotService) Publish(message tgbotapi.Chattable) {
	svc.publisher.Publish(message)
}

func (svc *GoTGBotService) Publisher() *BotPublisher {
	return svc.publisher
}

func (svc *GoTGBotService) Repository() IRepository {
	return svc.repository
}

func (svc *GoTGBotService) Listener() *BotListener {
	return svc.listener
}

func (svc *GoTGBotService) Configuration() *config.Configuration {
	return svc.configuration
}
