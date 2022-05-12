package config

import (
	"context"
	"encoding/json"
	"os"

	"github.com/joho/godotenv"
	"github.com/sethvargo/go-envconfig"
	"gopkg.in/yaml.v3"
)

// Configuration is the configuration of the bot.
type Configuration struct {
	Bot        BotConfiguration        `yaml:"bot" json:"bot" env:"BOT,prefix=BOT_"`
	Repository RepositoryConfiguration `yaml:"repository" json:"repository" env:"REPOSITORY,prefix=REPOSITORY_"`
	Logging    LoggerConfiguration     `yaml:"logging" json:"logging" env:"LOGGING,prefix=LOGGING_"`
}

// FixDefaults checks and fixes the configuration of defaults values.
func (cfg *Configuration) FixDefaults() *Configuration {
	cfg.Bot.fixDefaults()
	cfg.Repository.fixDefaults()
	cfg.Logging.FixDefaults()
	return cfg
}

func CreateConfigurationFromFile(filename string) (cfg *Configuration, err error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return
	}
	var config Configuration
	err = json.Unmarshal(content, &config)
	if err != nil {
		err = yaml.Unmarshal(content, &config)
		if err != nil {
			return
		}
	}
	config.FixDefaults()
	return &config, nil
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
	cfg.FixDefaults()
	return
}

func CreateConfigurationFromEnvFile(prefix string, filenames ...string) (cfg *Configuration, err error) {
	err = godotenv.Load(filenames...)
	if err != nil {
		return
	}
	return CreateConfigurationFromEnv(prefix)
}
