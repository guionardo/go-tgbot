package config

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

var expectedConfig = &Configuration{
	Bot: BotConfiguration{
		Token:      "token",
		Name:       "tgbot_test",
		HelloWorld: "Hello World!",
	},
	Repository: RepositoryConfiguration{
		ConnectionString:   "sqlite:///:memory:",
		HouseKeepingMaxAge: time.Hour,
	}, Logging: LoggerConfiguration{
		Level: "warn",
	},
}

func init() {
	expectedConfig.FixDefaults()
}

func TestCreateConfigurationFromFile(t *testing.T) {
	type args struct {
		filename string
	}

	tests := []struct {
		name    string
		args    args
		wantCfg *Configuration
		wantErr bool
	}{
		{
			name:    "CreateConfigurationFromFile_with_json",
			args:    args{filename: "sample_config.json"},
			wantCfg: expectedConfig,
			wantErr: false,
		},
		{
			name:    "CreateConfigurationFromFile_with_yaml",
			args:    args{filename: "sample_config.yaml"},
			wantCfg: expectedConfig,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCfg, err := CreateConfigurationFromFile(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateConfigurationFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !cmp.Equal(gotCfg, tt.wantCfg) {
				t.Errorf("CreateConfigurationFromFile() = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}

func TestCreateConfigurationFromEnv(t *testing.T) {

	t.Run("CreateConfigurationFromEnv", func(t *testing.T) {
		gotCfg, err := CreateConfigurationFromEnvFile("TG_", "sample.env")
		if err != nil {
			t.Errorf("CreateConfigurationFromEnv() error = %v", err)
			return
		}
		if !cmp.Equal(gotCfg, expectedConfig) {		
			t.Errorf("CreateConfigurationFromEnv() = %v, want %v", gotCfg, expectedConfig)
		}
	})

}
