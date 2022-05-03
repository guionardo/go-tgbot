package tgbot

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

type (
	Configuration struct {
		BotToken string
		LogLevel log.Level
	}
	fileConfiguration struct {
		BotToken string `json:"bot_token" yaml:"bot_token"`
		LogLevel string `json:"log_level" yaml:"log_level"`
	}
)

func LoadConfiguration() (cfg *Configuration, err error) {
	godotenv.Load()
	return getConfiguration(os.Getenv("TG_API_KEY"), os.Getenv("TG_LOG_LEVEL"))
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

	return getConfiguration(fCfg.BotToken, fCfg.LogLevel)
}

func getConfiguration(botToken string, sLogLevel string) (cfg *Configuration, err error) {
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

	cfg = &Configuration{
		BotToken: botToken,
		LogLevel: logLevel,
	}
	return
}
