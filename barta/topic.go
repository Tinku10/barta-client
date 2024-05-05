package barta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
