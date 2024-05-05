package main

import (
  "fmt"
  "flag"
  "barta-client/barta"
)

func main() {
  clusterAddr := flag.String("caddr", "localhost:8080", "Address of the cluster")
  flag.Parse()

  topic := barta.NewTopic("topic1", *clusterAddr)

  fmt.Println("-----------------------")
  // topic.PutMessage("Hello")
  // topic.PutMessage("Hello2")
  // topic.PutMessage("Hello3")
  fmt.Println(topic.GetMessage())
  // fmt.Println(topic.GetMessage())
  // fmt.Println(topic.GetMessage())
  fmt.Println("-----------------------")
}
