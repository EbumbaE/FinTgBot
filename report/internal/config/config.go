package config

import (
	"github.com/EbumbaE/FinTgBot/report/internal/clients/sender"
	"github.com/EbumbaE/FinTgBot/report/internal/servers/consumer"
	"github.com/EbumbaE/FinTgBot/report/internal/storage/psql"
	"github.com/spf13/viper"
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
