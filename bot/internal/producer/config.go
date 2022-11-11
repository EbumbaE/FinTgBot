package producer

type Config struct {
	KafkaTopic  string   `mapstructure:"kafkaTopic"`
	BrokersList []string `mapstructure:"brokersList"`
}
