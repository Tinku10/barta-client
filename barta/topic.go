package barta

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
  "errors"
)

type Topic struct {
	TopicName         string
	ReplicationFactor int
  ClusterAddr       string
}

func NewTopic(topicName string, clusterAddr string) (*Topic, error) {
  t := &Topic{
    TopicName:  topicName,
    ClusterAddr: clusterAddr,
  }

  stream, err := json.Marshal(t)
  if err != nil {
    log.Println("Unable to serialize topic")
    return new(Topic), errors.New("Unable to serialize topic")
  }
  r, err := http.Post(fmt.Sprintf("http://%s/topic/post", clusterAddr), "application/json", bytes.NewReader(stream))
  if err != nil {
    return new(Topic), err
  }

  if r.StatusCode != 201 {
    log.Println(r.StatusCode)
    return new(Topic), errors.New("Unable to create a topic")
  }


  return t, nil
}

func GetTopic(topicName string, clusterAddr string) (*Topic, error) {
  t := &Topic{
    TopicName:  topicName,
    ClusterAddr: clusterAddr,
  }

  _, err := json.Marshal(t)
  if err != nil {
    log.Println("Unable to serialize topic")
    return new(Topic), errors.New("Unable to serialize topic")
  }
  r, err := http.Get(fmt.Sprintf("http://%s/topic/get/%s", clusterAddr, topicName))
  if err != nil {
    return new(Topic), err
  }

  if r.StatusCode != 200 {
    log.Println("Topic does not exist")
    log.Println("Creating a new topic...")

    return NewTopic(topicName, clusterAddr)
  }

  log.Println("Getting the topic")
  json.NewDecoder(r.Body).Decode(&t)

  return t, nil
}
