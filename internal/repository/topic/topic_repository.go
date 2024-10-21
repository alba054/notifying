package topic

import (
	"alba054/kartjis-notify/internal/model/entity"
	"context"
	"database/sql"
)

type TopicRepository interface {
	FindTopicByName(ctx context.Context, db *sql.DB, name string) (*entity.TopicEntity, error)
	CreateTopic(ctx context.Context, db *sql.DB, name string) error
}
