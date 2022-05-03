package tgbot

import (
	"os"
	"reflect"
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestLoadConfiguration(t *testing.T) {
	os.Setenv("TG_API_KEY", "123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	os.Setenv("TG_LOG_LEVEL", "debug")
	t.Run("LoadConfiguration", func(t *testing.T) {
		wanted := &Configuration{
			BotToken: "123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			LogLevel: log.DebugLevel,
		}
		gotCfg, _ := LoadConfiguration()

		if !reflect.DeepEqual(gotCfg, wanted) {
			t.Errorf("LoadConfiguration() = %v, want %v", gotCfg, wanted)
		}
	})

}

func Test_getConfiguration(t *testing.T) {
	type args struct {
		botToken  string
		sLogLevel string
	}
	tests := []struct {
		name    string
		args    args
		wantCfg *Configuration
		wantErr bool
	}{
		{
			name:    "Valid configuration",
			args:    args{"123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ", "debug"},
			wantCfg: &Configuration{BotToken: "123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ", LogLevel: log.DebugLevel},
			wantErr: false,
		}, {
			name:    "Empty bot token",
			args:    args{"", "debug"},
			wantCfg: nil,
			wantErr: true,
		}, {
			name:    "Invalid log level",
			args:    args{"123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ", "invalid"},
			wantCfg: &Configuration{BotToken: "123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ", LogLevel: log.InfoLevel},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCfg, err := getConfiguration(tt.args.botToken, tt.args.sLogLevel)
			if (err != nil) != tt.wantErr {
				t.Errorf("getConfiguration() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("getConfiguration() = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}

func TestLoadConfigurationFromFile(t *testing.T) {

	tests := []struct {
		name    string
		content string

		wantCfg *Configuration
		wantErr bool
	}{
		{
			name:    "JSON file",
			content: "{\"bot_token\":\"123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ\",\"log_level\":\"debug\"}",
			wantCfg: &Configuration{BotToken: "123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ", LogLevel: log.DebugLevel},
			wantErr: false,
		},
    {
			name:    "YAML file",
			content: "bot_token: 123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ\nlog_level: debug",
			wantCfg: &Configuration{BotToken: "123456789:ABCDEFGHIJKLMNOPQRSTUVWXYZ", LogLevel: log.DebugLevel},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmp, err := os.CreateTemp(".", "test_config.file")
			if err != nil {
				t.Errorf("Failed to create temporary file %v", err)
				return
			}
			err = os.WriteFile(tmp.Name(), []byte(tt.content), 0644)
			if err != nil {
				t.Errorf("Failed to write temporary file %v", err)
				return
			}
			defer os.Remove(tmp.Name())
			gotCfg, err := LoadConfigurationFromFile(tmp.Name())
			if (err != nil) != tt.wantErr {
				t.Errorf("LoadConfigurationFromFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("LoadConfigurationFromFile() = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}
