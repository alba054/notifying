package notification

import (
	"alba054/kartjis-notify/internal/model/request"
	"context"
)

type NotificationService interface {
	AddMessageToTopic(ctx context.Context, payload request.PostNotificationMessagePayload) error
}
