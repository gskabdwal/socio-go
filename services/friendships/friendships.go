package friendships

import (
	"context"
	"encoding/json"
	"fmt"
	"socio/internals/cache"
	"socio/internals/dto"
	friendships "socio/models/friendship"
	"time"

	"github.com/google/uuid"
)

type Friends struct {
	UserID   uuid.UUID
	FriendID uuid.UUID

	Friends *dto.Friends

	AllFriends []dto.AllFriends
}

func New() *Friends {
	return &Friends{}
}

func (f *Friends) Create(ctx context.Context) {
	m := friendships.New()

	m.Friends = f.Friends

	m.Create(ctx)

	f.Friends.UpdatedAt = nil
}

func (f *Friends) GetAll(ctx context.Context) {
	val, err := cache.Client().Get(ctx, f.UserID.String()).Result()
	if val != "" && err == nil {
		json.Unmarshal([]byte(val), &f.AllFriends)
		return
	}

	m := friendships.New()
	m.UserID = f.UserID

	m.Get(ctx)

	f.AllFriends = m.AllFriends

	b, _ := json.Marshal(f.AllFriends)

	if err := cache.Client().Set(ctx, f.UserID.String(), b, 24*time.Hour).Err(); err != nil {
		fmt.Println("error setting up cache for friends: ", err)
	}
}

func (u *Friends) Delete(ctx context.Context) error {
	m := friendships.New()
	m.UserID = u.UserID
	m.FriendID = u.FriendID

	if err := m.Delete(ctx); err != nil {
		return err
	}

	return nil
}
