package barta

import (
	"time"
)

type MessageType int

const (
	REGULAR_MESSAGE MessageType = iota
	SENTINEL_MESSAGE
)

type Message struct {
	Key       string      `json:"key,omitempty"`
	Value     interface{} `json:"value"`
	TopicName string      `json:"topic"`
	Timestamp time.Time   `json:"timestamp"`
}
