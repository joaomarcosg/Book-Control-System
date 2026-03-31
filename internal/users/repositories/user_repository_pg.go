package repositories

import (
	"context"

	db "github.com/joaomarcosg/Book-Control-System/internal/database/sqlc"
	"github.com/joaomarcosg/Book-Control-System/internal/users/models"
)

type userRepository struct {
	queries *db.Queries
}

// CreateUser implements [UserRepository].
func (u *userRepository) CreateUser(ctx context.Context, user models.User) (models.User, error) {
	panic("unimplemented")
}

// DeleteUser implements [UserRepository].
func (u *userRepository) DeleteUser(ctx context.Context, id int64) error {
	panic("unimplemented")
}

// GetAllUsers implements [UserRepository].
func (u *userRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	panic("unimplemented")
}

// GetUser implements [UserRepository].
func (u *userRepository) GetUser(ctx context.Context, id int64) (models.User, error) {
	panic("unimplemented")
}

// UpdateUser implements [UserRepository].
func (u *userRepository) UpdateUser(ctx context.Context, id int64, user models.User) error {
	panic("unimplemented")
}

func NewUserRepository(q *db.Queries) UserRepository {
	return &userRepository{
		queries: q,
	}
}
