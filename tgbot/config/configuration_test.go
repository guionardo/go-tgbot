package config

import (
	"encoding/json"
	"os"
	"reflect"
	"testing"
	"time"

	"gopkg.in/yaml.v2"
)

func getConfFile(conf Configuration, fileType string) string {

	tmp, err := os.CreateTemp(".", "config.*")
	if err != nil {
		panic(err)
	}
	defer tmp.Close()
	err = saveConfFile(conf, tmp.Name(), fileType)
	// err = os.WriteFile(tmp.Name(), []byte(content), 0644)
	if err != nil {
		panic(err)
	}
	return tmp.Name()
}

func saveConfFile(conf Configuration, filename string, fileType string) (err error) {
	var content []byte
	if fileType == "json" {
		content, err = json.Marshal(conf)
	} else {
		content, err = yaml.Marshal(conf)
	}
	if err != nil {
		return err
	}
	return os.WriteFile(filename, content, 0644)
}

func TestCreateConfigurationFromFile(t *testing.T) {
	type args struct {
		filename string
	}
	cfg := Configuration{
		Bot: BotConfiguration{
			Token:      "token",
			Name:       "bot_name",
			HelloWorld: "hello_world",
		},
		LogLevel: "info",
		Repository: RepositoryConfiguration{
			ConnectionString:   "connection_string",
			HouseKeepingMaxAge: time.Hour,
		},
	}
	// 	jsonContent := `{"bot":{"token":"","name":"","hello_world":""},"repository":{"connection_string":"","house_keeping_max_age":"1h"},"log_level":"info"}`
	// 	yamlContent := `bot: {}
	// repository: {}
	// log_level: info
	// `
	tests := []struct {
		name    string
		args    args
		wantCfg *Configuration
		wantErr bool
	}{
		{
			name: "CreateConfigurationFromFile_with_json",
			args: args{filename: getConfFile(cfg, "json")},
			wantCfg: &Configuration{
				Bot:        BotConfiguration{},
				Repository: RepositoryConfiguration{},
				LogLevel:   "info",
			},
			wantErr: false,
		},
		{
			name: "CreateConfigurationFromFile_with_yaml",
			args: args{filename: getConfFile(cfg, "yaml")},
			wantCfg: &Configuration{
				Bot:        BotConfiguration{},
				Repository: RepositoryConfiguration{},
				LogLevel:   "info",
			},
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
			if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("CreateConfigurationFromFile() = %v, want %v", gotCfg, tt.wantCfg)
			}
			os.Remove(tt.args.filename)
		})
	}
}

func TestCreateConfigurationFromEnv(t *testing.T) {
	type args struct {
		prefix string
	}
	tests := []struct {
		name    string
		args    args
		wantCfg Configuration
		wantErr bool
	}{
		{
			name: "CreateConfigurationFromEnv_with_prefix",
			args: args{prefix: "TG_"},
			wantCfg: Configuration{
				Bot: BotConfiguration{
					Token:      "token",
					Name:       "tgbot_test",
					HelloWorld: "Hello World!",
				},
				Repository: RepositoryConfiguration{
					ConnectionString:   "sqlite:///:memory:",
					HouseKeepingMaxAge: time.Hour,
				},
				LogLevel: "warn",
			},
		},
	}
	os.Setenv("TG_BOT_TOKEN", "token")
	os.Setenv("TG_BOT_NAME", "tgbot_test")
	os.Setenv("TG_BOT_HELLO_WORLD", "Hello World!")
	os.Setenv("TG_REPOSITORY_CONNECTION_STRING", "sqlite:///:memory:")
	os.Setenv("TG_REPOSITORY_HOUSE_KEEPING_MAX_AGE", "1h")
	os.Setenv("TG_LOG_LEVEL", "warn")

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotCfg, err := CreateConfigurationFromEnv(tt.args.prefix)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateConfigurationFromEnv() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(gotCfg, tt.wantCfg) {
				t.Errorf("CreateConfigurationFromEnv() = %v, want %v", gotCfg, tt.wantCfg)
			}
		})
	}
}
