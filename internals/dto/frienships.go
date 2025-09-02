package dto

import (
	"time"

	"github.com/google/uuid"
)

type FriendsCrate struct {
	UserID   uuid.UUID `json:"user_id" validate:"required"`
	FriendID uuid.UUID `json:"friend_id" validate:"required"`
}

type Friends struct {
	ID        uint       `json:"id"`
	UserID    uuid.UUID  `json:"user_id"`
	FriendID  uuid.UUID  `json:"friend_id"`
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

type AllFriends struct {
	FriendID uuid.UUID `json:"friend_id"`
}
