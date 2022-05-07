package tgbot

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type (
	Configuration struct {
		BotToken                   string
		LogLevel                   log.Level
		RepositoryConnectionString string
		BotName                    string
		BotHelloWorld              string
		HouseKeepingMaxAge         time.Duration
	}
	fileConfiguration struct {
		BotToken                   string `json:"bot_token" yaml:"bot_token"`
		LogLevel                   string `json:"log_level" yaml:"log_level"`
		RepositoryConnectionString string `json:"repository_connection_string" yaml:"repository_connection_string"`
		BotName                    string `json:"bot_name" yaml:"bot_name"`
		BotHelloWorld              string `json:"bot_hello_world" yaml:"bot_hello_world"`
		HouseKeepingMaxAge         string `json:"house_keeping_max_age" yaml:"house_keeping_max_age"`
	}
)

const (
	TG_API_KEY                      = "TG_API_KEY"
	TG_LOG_LEVEL                    = "TG_LOG_LEVEL"
	TG_REPOSITORY_CONNECTION_STRING = "TG_REPOSITORY_CONNECTION_STRING"
	TG_BOT_NAME                     = "TG_BOT_NAME"
	TG_BOT_HELLO_WORLD              = "TG_BOT_HELLO_WORLD"
	TG_BOT_HOUSEKEEPING_MAX_AGE     = "TG_BOT_HOUSEKEEPING_MAX_AGE"
)

func LoadConfiguration() (cfg *Configuration, err error) {
	godotenv.Load()
	return getConfiguration(
		os.Getenv(TG_API_KEY),
		os.Getenv(TG_LOG_LEVEL),
		os.Getenv(TG_REPOSITORY_CONNECTION_STRING),
		os.Getenv(TG_BOT_NAME),
		os.Getenv(TG_BOT_HELLO_WORLD),
		os.Getenv(TG_BOT_HOUSEKEEPING_MAX_AGE),
	)
}

func LoadConfigurationFromFile(filename string) (cfg *Configuration, err error) {
	fileContent, err := os.ReadFile(filename)

	if err != nil {
		return
	}

	var fCfg fileConfiguration
	// Try unmarshal as json
	err = json.Unmarshal(fileContent, &fCfg)
	if err != nil {
		// Try unmarshal as yaml
		err = yaml.Unmarshal(fileContent, &fCfg)
	}
	if err != nil {
		return
	}

	return getConfiguration(fCfg.BotToken, fCfg.LogLevel, fCfg.RepositoryConnectionString, fCfg.BotName, fCfg.BotHelloWorld, fCfg.HouseKeepingMaxAge)
}

func getConfiguration(botToken string, sLogLevel string, sConnectionString string, botName string, botHelloWorld string, houseKeepingMaxAge string) (cfg *Configuration, err error) {
	if len(botToken) == 0 {
		err = fmt.Errorf("TG_API_KEY not found")
		return
	}
	logLevel, err := log.ParseLevel(sLogLevel)
	if err != nil {
		log.Warningf("Invalid TG_LOG_LEVEL: %v, using INFO", sLogLevel)
		logLevel = log.InfoLevel
	}
	log.SetLevel(logLevel)

	if len(botName) == 0 {
		botName = "tgbot"
	}
	if len(botHelloWorld) == 0 {
		botHelloWorld = "Hello World!"
	}
	var hk time.Duration
	if hk, err = time.ParseDuration(houseKeepingMaxAge); err != nil {
		log.Warningf("Invalid TG_BOT_HOUSEKEEPING_MAX_AGE: %v, using 24h", houseKeepingMaxAge)
		hk = time.Duration(24 * time.Hour)
	}

	cfg = &Configuration{
		BotToken:                   botToken,
		LogLevel:                   logLevel,
		RepositoryConnectionString: sConnectionString,
		BotName:                    botName,
		BotHelloWorld:              botHelloWorld,
		HouseKeepingMaxAge:         hk,
	}
	return
}
