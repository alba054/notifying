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

// get topic subscribers by id
func (storage *topicStorage) Get(subId string) *topicSubscriber {
	return storage.subscribers[subId]
}

// add new subscriber to the topic
func (storage *topicStorage) Set(subId string) error {
	if _, exists := storage.subscribers[subId]; exists {
		return errors.New("subscriber has been set")
	}

	storage.subscribers[subId] = newTopicSubscriber(subId)

	return nil
}

// push message to master queue
// then syncronize it to all subscribers's queue
func (storage *topicStorage) PushToMaster(message string) {
	storage.mu.Lock()
	storage.masterQueue.Enqueue(message)
	storage.mu.Unlock()

	storage.syncSubscriberQ()
}

// this will run after pushing to master queue
// this will naively dequeue the messages from master queue as if there is subscriber of the topic
func (storage *topicStorage) syncSubscriberQ() {
	var message *string
	wg := sync.WaitGroup{}
	for {
		storage.mu.Lock()
		message = storage.masterQueue.Dequeue() // * i'm still not sure if this should be locked
		storage.mu.Unlock()
		if message == nil {
			break
		}

		for _, subscriber := range storage.subscribers {
			wg.Add(1)
			go func() {
				subscriber.Set(*message)
				wg.Done()
			}()
		}
		wg.Wait()
	}
}
