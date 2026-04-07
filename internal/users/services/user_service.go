package services

import (
	"context"

	"github.com/joaomarcosg/Book-Control-System/internal/users/models"
	"github.com/joaomarcosg/Book-Control-System/internal/users/repositories"
)

type UserService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) repositories.UserRepository {
	return &UserService{
		userRepo: userRepo,
	}
}

func (u *UserService) CreateUser(ctx context.Context, user *models.User) (int64, error) {

	id, err := u.userRepo.CreateUser(ctx, user)

	if err != nil {
		return 0, err
	}

	return id, nil

}

func (u *UserService) DeleteUser(ctx context.Context, id int64) error {
	panic("unimplemented")
}

func (u *UserService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	panic("unimplemented")
}

func (u *UserService) GetUser(ctx context.Context, id int64) (*models.User, error) {
	panic("unimplemented")
}

func (u *UserService) UpdateUser(ctx context.Context, id int64, user *models.User) error {
	panic("unimplemented")
}
