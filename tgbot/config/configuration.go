package config

import (
	"context"
	"encoding/json"
	"os"

	"github.com/sethvargo/go-envconfig"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

// Configuration is the configuration of the bot.
type Configuration struct {
	Bot        BotConfiguration        `yaml:"bot" json:"bot" env:"BOT,prefix=BOT_"`
	Repository RepositoryConfiguration `yaml:"repository" json:"repository" env:"REPOSITORY,prefix=REPOSITORY_"`
	Logging    LoggerConfiguration     `yaml:"logging" json:"logging" env:"LOGGING,prefix=LOGGING_"`
	LogLevel   string                  `yaml:"log_level" json:"log_level" env:"LOG_LEVEL,default=info"`
	logLevel   logrus.Level
}

// fixDefaults checks and fixes the configuration of defaults values.
func (cfg *Configuration) fixDefaults() {
	cfg.Bot.fixDefaults()
	cfg.Repository.fixDefaults()
	cfg.Logging.FixDefaults()
}

func CreateConfigurationFromFile(filename string) (cfg *Configuration, err error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	err = json.Unmarshal(content, cfg)
	if err != nil {
		err = yaml.Unmarshal(content, cfg)
		if err != nil {
			return
		}
	}
	cfg.fixDefaults()
	return cfg, nil
}

func CreateConfigurationFromEnv(prefix string) (cfg *Configuration, err error) {
	ctx := context.Background()
	cfg = &Configuration{}
	if len(prefix) > 0 {
		l := envconfig.PrefixLookuper(prefix, envconfig.OsLookuper())
		err = envconfig.ProcessWith(ctx, cfg, l)

	} else {
		err = envconfig.Process(ctx, cfg)
	}
	cfg.fixDefaults()
	return
}
