package config

import (
	"strings"

	"github.com/sirupsen/logrus"
)

// LoggerConfiguration is the configuration for the logger
type LoggerConfiguration struct {
	Level           string       `yaml:"level" json:"level" env:"LEVEL,default=info"`
	FormatTimeStamp string       `yaml:"format_time_stamp" json:"format_time_stamp" env:"FORMAT_TIME_STAMP,default=2006-01-02 15:04:05"`
	Format          string       `yaml:"log_format" json:"log_format" env:"LOG_FORMAT,default=%time% %lvl% [%name%] %msg%\n"`
	LogLevel        logrus.Level `yaml:"-" json:"-"`
}

func (cfg *LoggerConfiguration) FixDefaults() LoggerConfiguration {
	cfg.Level = strings.ToLower(cfg.Level)
	logLevel, err := logrus.ParseLevel(cfg.Level)
	if err != nil {
		logLevel = logrus.InfoLevel
	}
	cfg.LogLevel = logLevel
	cfg.Level = logLevel.String()
	if len(cfg.FormatTimeStamp) == 0 {
		cfg.FormatTimeStamp = "2006-01-02 15:04:05"
	}
	if len(cfg.Format) == 0 {
		cfg.Format = "%time% %lvl% [%name%] %msg%\n"
	}
	return LoggerConfiguration{
		Level:           cfg.Level,
		FormatTimeStamp: cfg.FormatTimeStamp,
		Format:          cfg.Format,
		LogLevel:        cfg.LogLevel,
	}
}
