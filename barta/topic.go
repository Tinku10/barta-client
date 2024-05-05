package barta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

const (
	PartitionsPerTopic = 4
)

type Topic struct {
	TopicName         string
	ReplicationFactor int
  ClusterAddr       string
}

func NewTopic(topicName string, clusterAddr string) *Topic {
  t := &Topic{
    TopicName:  topicName,
    ClusterAddr: clusterAddr,
  }

  stream, err := json.Marshal(t)
  if err != nil {
    log.Println("Unable to serialize topic")
    return new(Topic)
  }
  _, err = http.Post(fmt.Sprintf("http://%s/topic/post", clusterAddr), "application/json", bytes.NewReader(stream))
  if err != nil {
    log.Println(err.Error())
  }

  return t
}

func (t *Topic) PutMessage(m string) {
  message := Message{Value: m}
  stream, err := json.Marshal(message)
  if err != nil {
    log.Println("Unable to deserialize message")
    return
  }
  r, err := http.Post(fmt.Sprintf("http://%s/message/post/%s", t.ClusterAddr, t.TopicName), "application/json", bytes.NewReader(stream))
  if err != nil {
    log.Printf("Error in the request %d", r.StatusCode)
    return
  }
}

func (t *Topic) GetMessage() (Message, error) {
  r, err := http.Get(fmt.Sprintf("http://%s/message/get/%s", t.ClusterAddr, t.TopicName))
  if err != nil {
    log.Println(r)
    log.Printf("Error in the request %d", r.StatusCode)
    return Message{}, err
  }

  var m Message
  json.NewDecoder(r.Body).Decode(&m)

  return m, nil
}
