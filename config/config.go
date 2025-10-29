package config

import "os"

type Config struct {
	KafkaBrokers []string
	Topic        string
	GroupID      string
}

func Load() *Config {
	brokers := os.Getenv("KAFKA_BROKERS")
	if brokers == "" {
		brokers = "localhost:9092"
	}

	topic := os.Getenv("KAFKA_TOPIC")
	if topic == "" {
		topic = "notifications"
	}

	groupID := os.Getenv("KAFKA_GROUP_ID")
	if groupID == "" {
		groupID = "notification-consumer-group"
	}

	return &Config{
		KafkaBrokers: []string{brokers},
		Topic:        topic,
		GroupID:      groupID,
	}
}
