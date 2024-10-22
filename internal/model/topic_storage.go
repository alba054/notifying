package model

import (
	"alba054/kartjis-notify/pkg"
	"errors"
	"sync"
)

type topicStorage struct {
	TopicName   string
	masterQueue pkg.Queue[string]
	subscribers map[string]*topicSubscriber
	mu          sync.Mutex
}

func newTopicStorage(topicName string) *topicStorage {
	return &topicStorage{
		TopicName:   topicName,
		masterQueue: pkg.Queue[string]{},
		subscribers: make(map[string]*topicSubscriber),
		mu:          sync.Mutex{},
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
	storage.mu.Lock()
	storage.masterQueue.Enqueue(message)
	storage.mu.Unlock()

	storage.syncSubscriberQ()
}

func (storage *topicStorage) syncSubscriberQ() {
	if len(storage.subscribers) < 1 {
		return
	}

	var message *string
	wg := sync.WaitGroup{}
	for {
		storage.mu.Lock()
		message = storage.masterQueue.Dequeue()
		storage.mu.Unlock()
		if message == nil {
			break
		}

		for key := range storage.subscribers {
			wg.Add(1)
			go func() {
				storage.subscribers[key].Set(*message)
				wg.Done()
			}()
		}
	}
	wg.Wait()
}
