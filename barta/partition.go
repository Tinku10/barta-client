package barta

import (
	"sync"
)

type Partition struct {
	PartitionId int
	Messages    []Message
  Topic       *Topic
	mutex       *sync.Mutex
}

func NewPartition(id int) *Partition {
	return &Partition{
		PartitionId: id,
		Messages:    []Message{},
		mutex:       &sync.Mutex{},
	}
}

func (p *Partition) WriteMessage(message *Message) {
	p.mutex.Lock()
	p.Messages = append(p.Messages, *message)
	p.mutex.Unlock()
}

func (p *Partition) ReadMessage(offset int) (Message, error) {
	return p.Messages[offset], nil
}
