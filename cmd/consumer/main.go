package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"test/config"
	"test/internal/kafka"
)

func main() {
	cfg := config.Load()
	log.Printf("Starting consumer with brokers: %v, topic: %s, group: %s\n",
		cfg.KafkaBrokers, cfg.Topic, cfg.GroupID)

	consumer, err := kafka.NewConsumer(cfg.KafkaBrokers, cfg.GroupID, cfg.Topic)
	if err != nil {
		log.Fatalf("Failed to create consumer: %v", err)
	}
	defer consumer.Close()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Обработка сигналов для graceful shutdown
	sigterm := make(chan os.Signal, 1)
	signal.Notify(sigterm, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigterm
		log.Println("Received termination signal, shutting down...")
		cancel()
	}()

	log.Println("Consumer started, waiting for messages...")
	if err := consumer.Start(ctx); err != nil {
		log.Fatalf("Error consuming messages: %v", err)
	}

	log.Println("Consumer stopped")
}
