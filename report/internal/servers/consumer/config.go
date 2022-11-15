package consumer

type Config struct {
	KafkaTopic         string   `mapstructure:"kafkaTopic"`
	KafkaConsumerGroup string   `mapstructure:"kafkaConsumerGroup"`
	BrokersList        []string `mapstructure:"brokersList"`
}
