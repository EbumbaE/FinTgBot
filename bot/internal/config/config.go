package config

import (
	"github.com/EbumbaE/FinTgBot/bot/internal/clients/producer"
	client "github.com/EbumbaE/FinTgBot/bot/internal/clients/tg"
	"github.com/EbumbaE/FinTgBot/bot/internal/currency"
	"github.com/EbumbaE/FinTgBot/bot/internal/servers/sender"
	server "github.com/EbumbaE/FinTgBot/bot/internal/servers/tg"
	"github.com/EbumbaE/FinTgBot/bot/internal/storage/psql"
	"github.com/spf13/viper"
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
