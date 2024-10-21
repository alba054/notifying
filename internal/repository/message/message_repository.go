package message

import (
	"alba054/kartjis-notify/internal/model/entity"
	"context"
	"database/sql"
)

type MessageRepository interface {
	CreateMessage(ctx context.Context, db *sql.DB, payload entity.MessageEntity) error
}
