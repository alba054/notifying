package model

import (
	"alba054/kartjis-notify/pkg"
	"sync"
)

type topicSubscriber struct {
	Id       string
	Messages pkg.Queue[string]
	mu       sync.Mutex
}

func (sub *topicSubscriber) Get() string {
	sub.mu.Lock()
	message := *sub.Messages.Dequeue()
	sub.mu.Unlock()
	return message
}

func (sub *topicSubscriber) Set(message string) {
	sub.mu.Lock()
	sub.Messages.Enqueue(message)
	sub.mu.Unlock()
}
