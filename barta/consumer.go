package barta

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Consumer struct {
	ConsumerID    string
	Subscriptions []*Topic
}

func NewConsumer(name string) *Consumer {
	return &Consumer{
		ConsumerID: name,
	}
}

func (c *Consumer) SubscribeToTopic(topic *Topic) {
  c.Subscriptions = append(c.Subscriptions, topic)
}

func (c *Consumer) GetMessages() ([]Message, error) {
  var messages []Message
  for _, s := range c.Subscriptions {
    r, err := http.Get(fmt.Sprintf("http://%s/message/get/%s/%s", s.ClusterAddr, s.TopicName, c.ConsumerID))
    if err != nil {
      log.Println(r)
      log.Printf("Error in the request %d", r.StatusCode)
      return []Message{}, err
    }

    var m Message
    json.NewDecoder(r.Body).Decode(&m)

    messages = append(messages, m)
  }

  return messages, nil
}
