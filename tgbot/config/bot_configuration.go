package config

type BotConfiguration struct {
	Token      string `yaml:"token" json:"token" env:"TOKEN,required"`
	Name       string `yaml:"name" json:"name" env:"NAME,default=tgbot"`
	HelloWorld string `yaml:"hello_world" json:"hello_world" env:"HELLO_WORLD,default=Hello World"`
}

// fixDefaults checks and fixes the configuration of defaults values of bot.
func (cfg *BotConfiguration) fixDefaults() {
	if len(cfg.Token) == 0 {
		panic("BotConfiguration.Token is required")
	}
	if len(cfg.Name) == 0 {
		cfg.Name = "tgbot"
	}
	if len(cfg.HelloWorld) == 0 {
		cfg.HelloWorld = "Hello World"
	}
}
