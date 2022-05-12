package tgbot

import (
	"github.com/guionardo/go-tgbot/tgbot/infra"
	"github.com/sirupsen/logrus"
)

type (
	BotRunner struct {
		logger *logrus.Entry
		name   string
	}
)

func (runner *BotRunner) Init(name string) {
	runner.logger = infra.GetLogger(name)
	runner.name = name
}

func (runner *BotRunner) String() string {
	return runner.name
}
