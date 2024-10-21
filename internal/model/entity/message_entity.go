package entity

import "database/sql"

type MessageEntity struct {
	Id        string
	Message   sql.NullString
	TopicId   int
	CreatedAt int64
	UpdatedAt int64
}
