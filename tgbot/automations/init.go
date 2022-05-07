package automations

import (
	"github.com/guionardo/go-tgbot/tgbot/infra"
	"github.com/sirupsen/logrus"
)

var logger *logrus.Entry

func getLogger() *logrus.Entry {
	if logger == nil {
		logger = infra.GetLogger("automations")
	}
	return logger
}
