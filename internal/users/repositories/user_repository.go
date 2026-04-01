package repositories

import (
	"context"
	"errors"

	"github.com/joaomarcosg/Book-Control-System/internal/users/models"
)

var (
	ErrDuplicateUserEmail = errors.New("user email already exists")
	ErrUserNotFound       = errors.New("user not found")
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *models.User) (*models.User, error)
	GetUser(ctx context.Context, id int64) (*models.User, error)
	GetAllUsers(ctx context.Context) ([]*models.User, error)
	UpdateUser(ctx context.Context, id int64, user *models.User) error
	DeleteUser(ctx context.Context, id int64) error
}
