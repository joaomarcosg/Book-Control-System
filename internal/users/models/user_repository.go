package models

import "context"

type UserRepository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUser(ctx context.Context, id int64) (*User, error)
	GetAllUsers(ctx context.Context) ([]*User, error)
	UpdateUser(ctx context.Context, id int64, user *User) error
	DeleteUser(ctx context.Context, id int64) error
}
