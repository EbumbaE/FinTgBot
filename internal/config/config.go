package config

import (
	"github.com/spf13/viper"
	client "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/clients/tg"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/currancy"
	server "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/servers/tg"
)

const configFile = "../../data"

type Config struct {
	TgClient client.Config   `mapstructure:"tgClient"`
	TgServer server.Config   `mapstructure:"tgServer"`
	Currency currancy.Config `mapstructure:"currency"`
}

func New() (cfg *Config, err error) {

	viper.AddConfigPath(configFile)
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")

	if err := viper.ReadInConfig(); err != nil {
		return &Config{}, err
	}

	err = viper.Unmarshal(&cfg)

	return cfg, err
}
