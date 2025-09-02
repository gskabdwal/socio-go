package users

import (
	"context"
	"socio/internals/dto"
	"socio/models/users"
)

type User struct {
	User  *dto.User
	Users *dto.Users
}

func New() *User {
	return &User{}
}

func (u *User) Create(ctx context.Context) {
	m := users.New()
	m.Name = u.User.Name
	m.Email = u.User.Email
	m.Password = u.User.Password

	m.Create(ctx)

	u.User.ID = m.ID
	u.User.CreatedAt = &m.CreatedAt
	u.User.UpdatedAt = nil

}

func (u *User) Get(ctx context.Context) error {
	m := users.New()
	m.ID = u.User.ID

	m.User = u.User

	if err := m.Get(ctx); err != nil {
		return err
	}

	return nil
}

func (u *User) Delete(ctx context.Context) error {
	m := users.New()
	m.ID = u.User.ID

	m.User = u.User

	if err := m.Delete(ctx); err != nil {
		return err
	}

	return nil
}

func (u *User) GetAll(ctx context.Context) error {
	m := users.New()

	m.Users = u.Users

	if err := m.GetAll(ctx); err != nil {
		return err
	}

	return nil
}
