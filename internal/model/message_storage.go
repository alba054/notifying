package model

import (
	"errors"
)

type MessageStorage struct {
	topicStorages map[string]*topicStorage
}

func New() *MessageStorage {
	return &MessageStorage{
		topicStorages: make(map[string]*topicStorage),
	}
}

// get topic
func (storage *MessageStorage) Get(topicId string) *topicStorage {
	return storage.topicStorages[topicId]
}

// set topic to local storage
func (storage *MessageStorage) Set(topicId string) error {
	if _, exists := storage.topicStorages[topicId]; exists {
		return errors.New("topic has been set")
	}

	storage.topicStorages[topicId] = newTopicStorage(topicId)

	return nil
}
