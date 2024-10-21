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

func (r *MessageRepositoryImpl) FindTopicByName(ctx context.Context, db *sql.DB, name string) (*entity.TopicEntity, error) {
	queryStmt := fmt.Sprintf("SELECT id, name FROM %s WHERE name = ?", r.tableName)

	stmt, err := db.PrepareContext(ctx, queryStmt)

	if err != nil {
		return nil, err
	}

	result, err := stmt.QueryContext(ctx, name)

	if err != nil {
		return nil, err
	}

	if !result.Next() {
		return nil, nil
	}

	var topic entity.TopicEntity
	err = result.Scan(&topic.Id, &topic.Name)

	if !result.Next() {
		return nil, err
	}
	err = result.Close()

	if !result.Next() {
		return nil, err
	}

	return &topic, nil
}
