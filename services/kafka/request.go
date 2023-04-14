package kafka

import (
	"time"
)

type MessageWrapper[T any] struct {
	Timestamp time.Time  `json:"@timestamp"`
	UserAgent string     `json:"user_agent"`
	Message   Message[T] `json:"message"`
}

type Message[T any] struct {
	Application string `json:"application"`
	Action      string `json:"action"`
	Body        T      `json:"body"`
}
