package model

import "alba054/kartjis-notify/pkg"

type topicSubscriber struct {
	Id       string
	Messages pkg.Queue[string]
}

func (sub *topicSubscriber) Get() string {
	return *sub.Messages.Dequeue()
}

func (sub *topicSubscriber) Set(message string) {
	sub.Messages.Enqueue(message)
}
