package users

import (
	"context"
	"fmt"
	"socio/internals/database"
	"socio/internals/dto"

	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	ID        uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	User  *dto.User  `gorm:"-"`
	Users *dto.Users `gorm:"-"`
}

func New() *Users {
	return &Users{}
}

func (u *Users) Create(ctx context.Context) error {
	if err := database.Client().Create(&u).Error; err != nil {
		fmt.Printf("Unable to create user: %v", err)
		return err
	}

	return nil
}

func (u *Users) Get(ctx context.Context) error {
	if err := database.Client().First(&u.User, u.ID).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("Error getting user: %v\n", err)
			return err
		}
	}
	return nil
}

func (u *Users) Delete(ctx context.Context) error {
	if err := database.Client().Delete(u).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("Error getting user: %v\n", err)
			return err
		}
	}
	return nil
}

func (u *Users) GetAll(ctx context.Context) error {
	if err := database.Client().Find(&u.Users.Users).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			fmt.Printf("Error getting users: %v\n", err)
			return err
		}
	}
	return nil
}
