package notification

import (
	"alba054/kartjis-notify/internal/model/request"
	"context"
)

type NotificationService interface {
	AddMessageToTopic(ctx context.Context, payload request.PostNotificationMessagePayload) error
	GetMessageNotification(ctx context.Context, topic string, subId string) (string, error)
	ActivateSubscriber(ctx context.Context, topic, subId string) error
	DeactivateSubscriber(ctx context.Context, topic, subId string) error
}
