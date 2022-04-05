package dto

import (
	"forum/internal/object"
)

type DTO interface {
	Create() *object.QuerySettings
	Delete() *object.QuerySettings
}

type Buf struct {
	Category     *Category
	CategoryPost *CategoryPost
	Comment      *Comment
	Like         *Ratio
	Post         *Post
	Session      *Session
}
