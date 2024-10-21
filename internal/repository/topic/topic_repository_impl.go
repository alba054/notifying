package topic

import (
	"alba054/kartjis-notify/internal/model/entity"
	"context"
	"database/sql"
	"fmt"
)

type TopicRepositoryImpl struct {
	tableName string
}

func New(tableName string) *TopicRepositoryImpl {
	return &TopicRepositoryImpl{
		tableName: tableName,
	}
}

func (r *TopicRepositoryImpl) CreateTopic(ctx context.Context, db *sql.DB, name string) error {
	queryStmt := fmt.Sprintf("INSERT INTO %s (name) VALUES (?)", r.tableName)

	stmt, err := db.PrepareContext(ctx, queryStmt)

	if err != nil {
		return err
	}

	_, err = stmt.ExecContext(ctx, name)

	if err != nil {
		return err
	}

	return nil
}

func (r *TopicRepositoryImpl) FindTopicByName(ctx context.Context, db *sql.DB, name string) (*entity.TopicEntity, error) {
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

	if err != nil {
		return nil, err
	}
	err = result.Close()

	if err != nil {
		return nil, err
	}

	return &topic, nil
}
