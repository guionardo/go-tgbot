package infra

import "github.com/sirupsen/logrus"

type BotLogger struct {
	logger *logrus.Entry
}

func CreateBotLogger(logger *logrus.Entry) *BotLogger {
	return &BotLogger{
		logger: logger,
	}
}

func (logger *BotLogger) Println(v ...interface{}) {
	logger.logger.Infoln(v...)
}

func (logger *BotLogger) Printf(format string, v ...interface{}) {
	logger.logger.Infof(format, v...)
}
