package models

import "time"

type NotificationType string

const (
	Email NotificationType = "email"
	SMS   NotificationType = "sms"
	Push  NotificationType = "push"
)

type Notification struct {
	ID        string           `json:"id"`
	Type      NotificationType `json:"type"`
	Recipient string           `json:"recipient"`
	Subject   string           `json:"subject,omitempty"`
	Message   string           `json:"message"`
	Timestamp time.Time        `json:"timestamp"`
}
