package topic

import (
	"alba054/kartjis-notify/internal/model/entity"
	"context"
	"database/sql"
)

type TopicRepositoryNilImpl struct {
}

func NewNil() *TopicRepositoryNilImpl {
	return &TopicRepositoryNilImpl{}
}

func (r *TopicRepositoryNilImpl) CreateTopic(ctx context.Context, db *sql.DB, name string) (int64, error) {
	return -1, nil
}

func (r *TopicRepositoryNilImpl) FindTopicByName(ctx context.Context, db *sql.DB, name string) (*entity.TopicEntity, error) {
	return nil, nil
}
