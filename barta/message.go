package barta

import (
	"time"
)

type Message struct {
	Key       string      `json:"key,omitempty"`
	Value     interface{} `json:"value"`
	TopicName string      `json:"topic"`
	Timestamp time.Time   `json:"timestamp"`
}
