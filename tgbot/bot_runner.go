package tgbot

import (
	"sync"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/guionardo/go-tgbot/tgbot/infra"
	"github.com/sirupsen/logrus"
)

type (
	BotRunner struct {
		bot             *tgbotapi.BotAPI
		isRunning       bool
		logger          *logrus.Entry
		name            string
		lock            sync.RWMutex
		internalChannel InternalChannel
	}
)

func (runner *BotRunner) IsRunning() bool {
	runner.lock.Lock()
	defer runner.lock.Unlock()
	return runner.isRunning
}

func (runner *BotRunner) IsRunningInt() int {
	if runner.IsRunning() {
		return 1
	}
	return 0
}

func (runner *BotRunner) GetName() string {
	return runner.name
}

func (runner *BotRunner) Start() {
	runner.lock.Lock()
	defer runner.lock.Unlock()
	runner.isRunning = true
	runner.logger.Info("started")
}

func (runner *BotRunner) Stop() {
	runner.lock.Lock()
	defer runner.lock.Unlock()
	runner.isRunning = false
	if runner.internalChannel != nil {
		runner.internalChannel.Stop(runner.name)
	}
	runner.logger.Info("stopped")
}

func (runner *BotRunner) SetInternalChannel(channel chan InternalMessage) {
	runner.internalChannel = channel
}

func (runner *BotRunner) Init(bot *tgbotapi.BotAPI, name string) {
	runner.bot = bot
	runner.logger = infra.GetLogger(name)
	runner.name = name
}

func (runner *BotRunner) String() string {
	return runner.name
}

func Wait(runners ...IContextRunner) {
	running := make(map[string]IContextRunner)
	logger := infra.GetLogger("Wait")
	for _, runner := range runners {
		running[runner.GetName()] = runner
	}
	for len(running) > 0 {
		for name, runner := range running {
			if runner == nil || !runner.IsRunning() {
				delete(running, name)
			}
		}
		logger.Debugf("%d runners running = %v", len(running), running)
		time.Sleep(time.Second)
	}
}
