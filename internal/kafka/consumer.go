package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"test/internal/models"

	"github.com/IBM/sarama"
)

type Consumer struct {
	consumerGroup sarama.ConsumerGroup
	topic         string
}

func NewConsumer(brokers []string, groupID, topic string) (*Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Group.Rebalance.Strategy = sarama.NewBalanceStrategyRoundRobin()
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup(brokers, groupID, config)
	if err != nil {
		return nil, fmt.Errorf("failed to create consumer group: %w", err)
	}

	return &Consumer{
		consumerGroup: consumerGroup,
		topic:         topic,
	}, nil
}

func (c *Consumer) Start(ctx context.Context) error {
	handler := &consumerGroupHandler{}

	for {
		if err := c.consumerGroup.Consume(ctx, []string{c.topic}, handler); err != nil {
			return fmt.Errorf("error from consumer: %w", err)
		}

		if ctx.Err() != nil {
			return ctx.Err()
		}
	}
}

func (c *Consumer) Close() error {
	return c.consumerGroup.Close()
}

type consumerGroupHandler struct{}

func (h *consumerGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	log.Println("Consumer group session setup")
	return nil
}

func (h *consumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	log.Println("Consumer group session cleanup")
	return nil
}

func (h *consumerGroupHandler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		var notification models.Notification
		if err := json.Unmarshal(message.Value, &notification); err != nil {
			log.Printf("Error unmarshaling message: %v\n", err)
			continue
		}

		log.Printf("Received notification: ID=%s, Type=%s, Recipient=%s, Message=%s\n",
			notification.ID, notification.Type, notification.Recipient, notification.Message)

		// Здесь должна быть логика отправки нотификации
		if err := processNotification(&notification); err != nil {
			log.Printf("Error processing notification: %v\n", err)
		}

		session.MarkMessage(message, "")
	}

	return nil
}

func processNotification(notification *models.Notification) error {
	// Имитация обработки нотификации
	log.Printf("Processing %s notification for %s\n", notification.Type, notification.Recipient)

	switch notification.Type {
	case models.Email:
		log.Printf("Sending email to: %s\n", notification.Recipient)
	case models.SMS:
		log.Printf("Sending SMS to: %s\n", notification.Recipient)
	case models.Push:
		log.Printf("Sending push notification to: %s\n", notification.Recipient)
	}

	return nil
}
