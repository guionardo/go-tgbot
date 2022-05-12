package tgbot

import (
	"context"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/guionardo/go-tgbot/tgbot/config"
	"github.com/guionardo/go-tgbot/tgbot/infra"
	"github.com/guionardo/go-tgbot/tgbot/runners"
	"gorm.io/gorm"
)

type GoTGBotService struct {
	BotRunner
	bot           *tgbotapi.BotAPI
	context       context.Context
	listener      *BotListener
	worker        *BotWorker
	publisher     *BotPublisher
	cancel        context.CancelFunc
	stop          context.CancelFunc
	repository    IRepository
	configuration *config.Configuration
	setupLevel    SetupLevel
}

// CreateBotService creates a new empty bot service
// This is the first step to create a bot service
//
// After this step, you should:
//
// * LoadConfigurationFromFile or LoadConfigurationFromEnv
//
// * InitBot
//
// * AddRepository (optional, if not set, automatic repository will be created)
func CreateBotService() *GoTGBotService {
	return &GoTGBotService{
		setupLevel: Instance,
	}
}

func (svc *GoTGBotService) LoadConfigurationFromFile(filename string) *GoTGBotService {
	cfg, err := config.CreateConfigurationFromFile(filename)
	if err != nil {
		panic(err)
	}
	svc.configuration = cfg
	svc.setupLevel = Set(svc.setupLevel, Configuration)
	return svc
}

func (svc *GoTGBotService) LoadConfigurationFromEnv(prefix string) *GoTGBotService {
	cfg, err := config.CreateConfigurationFromEnv(prefix)
	if err != nil {
		panic(err)
	}
	svc.configuration = cfg
	svc.setupLevel = Set(svc.setupLevel, Configuration)
	return svc
}

// InitBot setups bot, loggin
func (svc *GoTGBotService) InitBot() *GoTGBotService {
	if !Has(svc.setupLevel, Configuration) || svc.configuration == nil {
		panic("configuration not set")
	}

	// Setup logging
	infra.CreateLoggerFactory(&svc.configuration.Logging)
	svc.logger = infra.GetLogger("GoTGBotService")
	svc.logger.Infof("Init %s", svc.name)

	// Setup bot
	if err := tgbotapi.SetLogger(infra.CreateBotLogger(infra.GetLogger("tgbot"))); err != nil {
		panic(err)
	}
	bot, err := tgbotapi.NewBotAPI(svc.configuration.Bot.Token)
	if err != nil {
		panic(err)
	}
	svc.bot = bot
	svc.logger.Infof("Authorized on account %s", bot.Self.UserName)
	svc.setupLevel = Set(svc.setupLevel, Init)

	// Create listener
	svc.listener = createBotListener(svc.bot)

	// Create worker
	svc.worker = createBotWorker()

	// Create publisher
	svc.publisher = createBotPublisher()
	svc.setupLevel = Set(svc.setupLevel, Workers)
	svc.Init("GoTGBotService")
	return svc
}

func (svc *GoTGBotService) AddRepository(dbs ...*gorm.DB) *GoTGBotService {
	var db *gorm.DB
	var err error
	if len(dbs) == 0 {
		// Create default SQLite repository from connection string
		db, err = infra.GetSQLiteDB(svc.configuration.Repository.ConnectionString)

	} else {
		db = dbs[0]
	}
	if err != nil {
		panic(err)
	}
	svc.repository = infra.CreateMessageRepository(db)
	svc.setupLevel = Set(svc.setupLevel, Repository)
	return svc
}

func (svc *GoTGBotService) SetRepository(repository IRepository) *GoTGBotService {
	svc.repository = repository
	svc.setupLevel = Set(svc.setupLevel, Repository)
	return svc
}

func (svc *GoTGBotService) Start() error {
	if svc.cancel != nil {
		return fmt.Errorf("service already started")
	}

	if !Has(svc.setupLevel, Init) {
		return fmt.Errorf("service not initialized by svc.InitBot()")
	}
	if !Has(svc.setupLevel, Workers) {
		return fmt.Errorf("service not initialized by svc.AddWorkers()")
	}
	if !Has(svc.setupLevel, Repository) {
		svc.logger.Warnf("automatic repository created: %s", svc.configuration.Repository.ConnectionString)
		svc.AddRepository()
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
