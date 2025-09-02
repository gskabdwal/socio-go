package dto

import (
	"time"

	"github.com/google/uuid"
)

type PostCreate struct {
	Content string `json:"content" validate:"required,max=2000"`
}

type Post struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Content   string     `json:"content"`
	UserID    uuid.UUID  `json:"user_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type Posts struct {
	Posts []Post `json:"posts"`
}
