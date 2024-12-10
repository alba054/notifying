package message

import (
	"alba054/kartjis-notify/internal/model/entity"
	"context"
	"database/sql"
)

type MessageRepositoryNilImpl struct {
}

func NewNil() *MessageRepositoryNilImpl {
	return &MessageRepositoryNilImpl{}
}

func (r *MessageRepositoryNilImpl) CreateMessage(ctx context.Context, db *sql.DB, payload entity.MessageEntity) error {
	return nil
}
