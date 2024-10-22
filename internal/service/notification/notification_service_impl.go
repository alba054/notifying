package notification

import (
	"alba054/kartjis-notify/internal/exception"
	"alba054/kartjis-notify/internal/model"
	"alba054/kartjis-notify/internal/model/entity"
	"alba054/kartjis-notify/internal/model/request"
	"alba054/kartjis-notify/internal/repository/message"
	"alba054/kartjis-notify/internal/repository/topic"
	"context"
	"database/sql"
	"strings"
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

func (s *NotificationServiceImpl) AddMessageToTopic(ctx context.Context, payload request.PostNotificationMessagePayload) error {
	if strings.Trim(payload.Id, " ") == "" || strings.Trim(payload.Topic, " ") == "" {
		return exception.NewBadRequestError("id and topic is not formed correctly")
	}

	if strings.Trim(payload.Message, " ") == "" {
		return exception.NewBadRequestError("message can't be empty")
	}

	topicEty, err := s.topicRepository.FindTopicByName(ctx, s.db, payload.Topic)

	if err != nil {
		return err
	}

	if topicEty == nil {
		// insert topic to database first
		err = s.topicRepository.CreateTopic(ctx, s.db, payload.Topic)
		if err != nil {
			return err
		}

		// then create a local storage for topic
		err = s.messageStorage.Set(payload.Topic)
		if err != nil {
			return err
		}
	}

	messageEntity := entity.MessageEntity{
		Id:      payload.Id,
		Message: sql.NullString{String: payload.Message},
		TopicId: topicEty.Id,
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
