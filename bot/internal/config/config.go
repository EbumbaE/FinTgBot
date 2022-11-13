package config

import (
	"github.com/spf13/viper"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/clients/sender"
	client "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/clients/tg"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/currency"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/producer"
	server "gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/servers/tg"
	"gitlab.ozon.dev/ivan.hom.200/telegram-bot/internal/storage/psql"
)

const configFile = "../../data"

type Config struct {
	TgClient     client.Config   `mapstructure:"TgClient"`
	TgServer     server.Config   `mapstructure:"TgServer"`
	Currency     currency.Config `mapstructure:"Currency"`
	PsqlDatabase psql.Config     `mapstructure:"Psql"`
	Producer     producer.Config `mapstructure:"Producer"`
	SenderServer sender.Config   `mapstructure:"Sender"`
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
