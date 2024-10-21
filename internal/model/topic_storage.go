package model

import (
	"alba054/kartjis-notify/pkg"
	"errors"
)

type topicStorage struct {
	TopicName   string
	masterQueue pkg.Queue[string]
	subscribers map[string]*topicSubscriber
}

func newTopicStorage(topicName string) *topicStorage {
	return &topicStorage{
		TopicName:   topicName,
		masterQueue: pkg.Queue[string]{},
		subscribers: make(map[string]*topicSubscriber),
	}
}

func (storage *topicStorage) Get(subId string) *topicSubscriber {
	return storage.subscribers[subId]
}

func (storage *topicStorage) Set(subId string, value *topicSubscriber) error {
	if _, exists := storage.subscribers[subId]; exists {
		return errors.New("subscriber has been set")
	}

	storage.subscribers[subId] = value

	return nil
}

func (storage *topicStorage) PushToMaster(message string) {
	storage.masterQueue.Enqueue(message)
}
