package main

import (
	"barta-client/barta"
	"flag"
	"log"
)

func main() {
  clusterAddr := flag.String("caddr", "localhost:8080", "Address of the cluster")
  flag.Parse()

  topic := barta.NewTopic("topic1", *clusterAddr)
  // producer := barta.NewProducer("producer1")
  consumer := barta.NewConsumer("consumer1")
  consumer.SubscribeToTopic(topic)

  // producer.PutMessage(topic, "hello-----------")
  // producer.PutMessage(topic, "hello2-----------")
  messages, err := consumer.GetMessages()
  if err != nil {
    log.Println(err)
  }

  for _, m := range messages {
    log.Println(m)
  }
}
