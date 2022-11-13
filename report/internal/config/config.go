package config

import (
	"github.com/spf13/viper"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/clients/sender"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/servers/consumer"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/storage/psql"
)

const configFile = "../../data"

type Config struct {
	PsqlDatabase psql.Config     `mapstructure:"Psql"`
	Consumer     consumer.Config `mapstructure:"Consumer"`
	SenderClient sender.Config   `mapstructure:"Sender"`
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
