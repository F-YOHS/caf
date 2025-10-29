package main

import (
	"log"
	"time"

	"test/config"
	"test/internal/kafka"
	"test/internal/models"

	"github.com/google/uuid"
)

func main() {
	cfg := config.Load()
	log.Printf("Starting producer with brokers: %v, topic: %s\n", cfg.KafkaBrokers, cfg.Topic)

	producer, err := kafka.NewProducer(cfg.KafkaBrokers, cfg.Topic)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	// Отправляем несколько тестовых нотификаций
	notifications := []models.Notification{
		{
			ID:        uuid.New().String(),
			Type:      models.Email,
			Recipient: "user@example.com",
			Subject:   "Welcome!",
			Message:   "Welcome to our service!",
			Timestamp: time.Now(),
		},
		{
			ID:        uuid.New().String(),
			Type:      models.SMS,
			Recipient: "+1234567890",
			Message:   "Your verification code is 123456",
			Timestamp: time.Now(),
		},
		{
			ID:        uuid.New().String(),
			Type:      models.Push,
			Recipient: "device-token-123",
			Message:   "You have a new message",
			Timestamp: time.Now(),
		},
	}

	for _, notification := range notifications {
		if err := producer.SendNotification(&notification); err != nil {
			log.Printf("Error sending notification: %v\n", err)
		} else {
			log.Printf("Notification sent: %s\n", notification.ID)
		}
		time.Sleep(time.Second)
	}

	log.Println("All notifications sent successfully")
}
