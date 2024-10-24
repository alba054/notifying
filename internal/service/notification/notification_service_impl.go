package notification

import (
	"alba054/kartjis-notify/internal/exception"
	"alba054/kartjis-notify/internal/model"
	"alba054/kartjis-notify/internal/model/entity"
	"alba054/kartjis-notify/internal/model/request"
	"alba054/kartjis-notify/internal/repository/message"
	"alba054/kartjis-notify/internal/repository/topic"
	"alba054/kartjis-notify/shared"
	"context"
	"database/sql"
)

type NotificationServiceImpl struct {
	topicRepository   topic.TopicRepository
	messageStorage    *model.MessageStorage
	messageRepository message.MessageRepository
	db                *sql.DB
}

func New(
	topicRepo topic.TopicRepository,
	messageRepo message.MessageRepository,
	db *sql.DB,
	messageStorage *model.MessageStorage,
) *NotificationServiceImpl {
	return &NotificationServiceImpl{
		topicRepository:   topicRepo,
		messageRepository: messageRepo,
		messageStorage:    messageStorage,
		db:                db,
	}
}

// deactivate subscriber
// allow another listener to listen to this
func (s *NotificationServiceImpl) DeactivateSubscriber(ctx context.Context, topic, subId string) error {
	s.messageStorage.Set(topic)

	topicStorage := s.messageStorage.Get(topic)

	topicStorage.Set(subId)
	subscriber := topicStorage.Get(subId)

	subscriber.Deactivate()
	return nil
}

// activate subscriber if not yet activated
// this will return an error if more than one listener use this
func (s *NotificationServiceImpl) ActivateSubscriber(ctx context.Context, topic, subId string) error {
	s.messageStorage.Set(topic)

	topicStorage := s.messageStorage.Get(topic)

	topicStorage.Set(subId)
	subscriber := topicStorage.Get(subId)

	if subscriber.IsActive() {
		return exception.NewBadRequestError("subscriber's been active, cannot use the same subId")
	}

	subscriber.Activate()
	return nil
}

func (s *NotificationServiceImpl) GetMessageNotification(ctx context.Context, topic string, subId string) (string, error) {
	topicStorage := s.messageStorage.Get(topic)
	// if topic doesn't exist in local storage
	// there is a posibility that topic hasn't been created in the database
	if topicStorage == nil {
		_, err := s.createTopicToDb(ctx, topic) // we insert topic in the db if not yet

		if err != nil {
			return "", err
		}

		s.messageStorage.Set(topic)
		topicStorage = s.messageStorage.Get(topic)
	}

	topicStorage.Set(subId)
	subscriber := topicStorage.Get(subId)
	message := subscriber.Get()

	return message, nil
}

func (s *NotificationServiceImpl) AddMessageToTopic(ctx context.Context, payload request.PostNotificationMessagePayload) error {
	if shared.IsEmptyString(payload.Id) || shared.IsEmptyString(payload.Topic) {
		return exception.NewBadRequestError("id and topic is not formed correctly")
	}

	if shared.IsEmptyString(payload.Message) {
		return exception.NewBadRequestError("message can't be empty")
	}

	topicId, err := s.createTopicToDb(ctx, payload.Topic)

	if err != nil {
		return err
	}

	// then create a local storage for topic
	s.messageStorage.Set(payload.Topic)

	messageEntity := entity.MessageEntity{
		Id:      payload.Id,
		Message: sql.NullString{String: payload.Message},
		TopicId: topicId,
	}

	// insert message to database first
	err = s.messageRepository.CreateMessage(ctx, s.db, messageEntity)

	if err != nil {
		return err
	}

	// then push message to local storage
	// this will also sync the subscribers's queue
	// all messages that are in master queue will be moved to each subscribers's queue
	s.messageStorage.Get(payload.Topic).PushToMaster(payload.Message)

	return nil
}

func (s *NotificationServiceImpl) createTopicToDb(ctx context.Context, topic string) (int64, error) {
	topicEty, err := s.topicRepository.FindTopicByName(ctx, s.db, topic)

	if err != nil {
		return -1, err
	}

	var topicId int64
	if topicEty == nil {
		// insert topic to database first
		id, err := s.topicRepository.CreateTopic(ctx, s.db, topic)
		if err != nil {
			return -1, err
		}

		topicId = id
	} else {
		topicId = topicEty.Id
	}

	return topicId, nil
}
