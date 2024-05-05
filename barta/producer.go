package barta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Producer struct {
	ProducerID string
}

func NewProducer(name string) *Producer {
	return &Producer{
		ProducerID: name,
	}
}

func (p *Producer) PutMessage(topic *Topic, m string) {
	message := Message{Value: m}
	stream, err := json.Marshal(message)
	if err != nil {
		log.Println("Unable to deserialize message")
		return
	}
	r, err := http.Post(fmt.Sprintf("http://%s/message/post/%s", topic.ClusterAddr, topic.TopicName), "application/json", bytes.NewReader(stream))
	if err != nil {
		log.Printf("Error in the request %d", r.StatusCode)
		return
	}
}
