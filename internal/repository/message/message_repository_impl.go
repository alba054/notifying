package message

import (
	"alba054/kartjis-notify/internal/model/entity"
	"context"
	"database/sql"
	"fmt"
)

type MessageRepositoryImpl struct {
	tableName string
}

func New(tableName string) *MessageRepositoryImpl {
	return &MessageRepositoryImpl{
		tableName: tableName,
	}
}

func (r *MessageRepositoryImpl) CreateMessage(ctx context.Context, db *sql.DB, payload entity.MessageEntity) error {
	queryStmt := fmt.Sprintf("INSERT INTO %s (id, message, topicId) VALUES (?, ?, ?)", r.tableName)

	stmt, err := db.PrepareContext(ctx, queryStmt)

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, payload.Id, payload.Message.String, payload.TopicId)

	if err != nil {
		return err
	}

	return nil
}
