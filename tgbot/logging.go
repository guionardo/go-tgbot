package tgbot

import (
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

var (
	loggers      = make(map[string]*log.Entry)
	defaultLevel = log.InfoLevel
)

func init() {
	godotenv.Load()
	var err error
	defaultLevel, err = log.ParseLevel(os.Getenv("TG_LOG_LEVEL"))
	if err != nil {
		defaultLevel = log.InfoLevel
	}
}

func GetLogger(name string) *log.Entry {
	if logger, ok := loggers[name]; ok {
		return logger
	}
	logger := &log.Logger{
		Out: os.Stdout,
		Formatter: &easy.Formatter{
			TimestampFormat: "2006-01-02 15:04:05",
			LogFormat:       "%time% %lvl% [%name%] %msg%\n",
		},
		Level: defaultLevel,
	}

	entry := logger.WithField("name", name)
	loggers[name] = entry
	return entry
}
