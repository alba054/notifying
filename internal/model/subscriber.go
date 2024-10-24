package model

import (
	"alba054/kartjis-notify/pkg"
	"sync"
)

type topicSubscriber struct {
	id       string
	messages pkg.Queue[string]
	mu       sync.Mutex
	isActive bool
}

func newTopicSubscriber(subId string) *topicSubscriber {
	return &topicSubscriber{
		id:       subId,
		messages: pkg.Queue[string]{},
		mu:       sync.Mutex{},
		isActive: false,
	}
}

func (sub *topicSubscriber) IsActive() bool {
	return sub.isActive
}

func (sub *topicSubscriber) Activate() {
	sub.mu.Lock()
	sub.isActive = true
	sub.mu.Unlock()
}

func (sub *topicSubscriber) Deactivate() {
	sub.mu.Lock()
	sub.isActive = false
	sub.mu.Unlock()
}

// get message from queue
func (sub *topicSubscriber) Get() string {
	sub.mu.Lock()
	message := sub.messages.Dequeue()
	sub.mu.Unlock()

	if message == nil {
		return ""
	}

	return *message
}

// enqueue new message
func (sub *topicSubscriber) Set(message string) {
	sub.mu.Lock()
	sub.messages.Enqueue(message)
	sub.mu.Unlock()
}
