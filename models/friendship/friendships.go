package friendships

import (
	"context"
	"fmt"
	"socio/internals/database"
	"socio/internals/dto"
	"socio/models/users"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Friendships struct {
	gorm.Model
	UserID   uuid.UUID `gorm:"uniqueIndex:idx_user_friend" json:"user_id"`
	FriendID uuid.UUID `gorm:"uniqueIndex:idx_user_friend" json:"friend_id"`

	User   users.Users `gorm:"foreignKey:UserID;references:ID" json:"-"`
	Friend users.Users `gorm:"foreignKey:FriendID;references:ID" json:"-"`

	Friends    *dto.Friends     `gorm:"-"`
	AllFriends []dto.AllFriends `gorm:"-"`
}

func New() *Friendships {
	return &Friendships{}
}

func (f *Friendships) Create(ctx context.Context) error {
	if err := database.Client().Table("friendships").Create(&f.Friends).Error; err != nil {
		fmt.Printf("Unable to create user: %v", err)
		return err
	}

	return nil
}

func (f *Friendships) Get(ctx context.Context) error {
	if err := database.Client().Table("friendships").
		Where("user_id = ?", f.UserID).
		Find(&f.AllFriends).Error; err != nil {
		fmt.Printf("Unable to create user: %v", err)
		return err
	}

	return nil
}

func (f *Friendships) Delete(ctx context.Context) error {
	if err := database.Client().
		Where("user_id = ?", f.UserID).
		Where("friend_id = ?", f.FriendID).
		Unscoped(). // This forces a hard delete (bypasses the soft delete)
		Delete(f).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("Error getting user: %v\n", err)
			return err
		}
	}
	return nil
}
