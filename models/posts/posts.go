package posts

import (
	"context"
	"fmt"
	"socio/internals/database"
	"socio/internals/dto"
	"socio/models/users"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Posts struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Content   string    `json:"content"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User users.Users `gorm:"foreignKey:UserID;references:ID" json:"-"`

	Post  *dto.Post  `gorm:"-"`
	Posts *dto.Posts `gorm:"-"`
}

func New() *Posts {
	return &Posts{}
}

func (p *Posts) Create(ctx context.Context) error {
	if err := database.Client().Table("posts").Create(&p.Post).Error; err != nil {
		fmt.Printf("Unable to create post: %v", err)
		return err
	}
	return nil
}

func (p *Posts) Get(ctx context.Context) error {
	if err := database.Client().
		Table("posts").
		Where("user_id = ?", p.UserID).
		Find(&p.Posts.Posts).Error; err != nil {
		fmt.Printf("Unable to get posts: %v", err)
		return err
	}

	return nil
}

func (p *Posts) Delete(ctx context.Context) error {
	if err := database.Client().
		Where("user_id = ?", p.UserID).
		Where("id = ?", p.ID).
		Delete(p).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("Error getting user: %v\n", err)
			return err
		}
	}
	return nil
}
