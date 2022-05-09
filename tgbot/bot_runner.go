package tgbot

import (
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/guionardo/go-tgbot/tgbot/infra"
	"github.com/sirupsen/logrus"
)

type (
	BotRunner struct {
		bot       *tgbotapi.BotAPI
		isRunning bool
		logger    *logrus.Entry
		name      string
		lock      sync.RWMutex
	}
)

func (runner *BotRunner) GetName() string {
	return runner.name
}

func (runner *BotRunner) Init(bot *tgbotapi.BotAPI, name string) {
	runner.bot = bot
	runner.logger = infra.GetLogger(name)
	runner.name = name
}

func (runner *BotRunner) String() string {
	return runner.name
}
