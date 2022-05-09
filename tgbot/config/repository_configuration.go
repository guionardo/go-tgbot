package config

import "time"

type RepositoryConfiguration struct {
	ConnectionString   string        `yaml:"connection_string" json:"connection_string" env:"CONNECTION_STRING"`
	HouseKeepingMaxAge time.Duration `json:"house_keeping_max_age" yaml:"house_keeping_max_age" env:"HOUSE_KEEPING_MAX_AGE"`
}

// fixDefaults checks and fixes the configuration of repository defaults values.
func (cfg *RepositoryConfiguration) fixDefaults() {
	if cfg.HouseKeepingMaxAge == 0 {
		cfg.HouseKeepingMaxAge = time.Hour
	}
}
