package entity

import "database/sql"

type MessageEntity struct {
	Id        string
	Message   sql.NullString
	TopicId   int64
	CreatedAt int64
	UpdatedAt int64
}
