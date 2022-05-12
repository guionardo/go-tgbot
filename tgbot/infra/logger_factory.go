package infra

import (
	"os"

	"github.com/guionardo/go-tgbot/tgbot/config"
	log "github.com/sirupsen/logrus"
	easy "github.com/t-tomalak/logrus-easy-formatter"
)

type (
	LoggerFactory struct {
		instances map[string]*log.Entry
		config    *config.LoggerConfiguration
	}
)

var loggerFactory *LoggerFactory

func CreateLoggerFactory(config *config.LoggerConfiguration) *LoggerFactory {
	if loggerFactory == nil {
		config.FixDefaults()
		factory := &LoggerFactory{
			instances: make(map[string]*log.Entry),
			config:    config,
		}
		loggerFactory = factory
	}
	return loggerFactory
}

func GetLoggerFactory() *LoggerFactory {
	if loggerFactory == nil {
		panic("LoggerFactory not initialized")
	}
	return loggerFactory
}

func (factory *LoggerFactory) GetLogger(name string) *log.Entry {
	if logger, ok := factory.instances[name]; ok {
		return logger
	}
	logger := &log.Logger{
		Out: os.Stdout,
		Formatter: &easy.Formatter{TimestampFormat: "2006-01-02 15:04:05",
			LogFormat: "%time% %lvl% [%name%] %msg%\n",
		},
		Level: factory.config.LogLevel,
	}
	entry := logger.WithField("name", name)
	factory.instances[name] = entry
	return entry
}

func GetLogger(name string) *log.Entry {
	return GetLoggerFactory().GetLogger(name)
}
