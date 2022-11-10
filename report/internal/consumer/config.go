package consumer

type Config struct {
	KafkaTopic         string
	KafkaConsumerGroup string
	BrokersList        []string
}
