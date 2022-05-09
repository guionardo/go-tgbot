package automations

import (
	"github.com/guionardo/go-tgbot/tgbot/infra"
	"github.com/sirupsen/logrus"
)

const authLoggerName = "automations"

func getLogger() *logrus.Entry {
	return infra.GetLogger(authLoggerName)
}
