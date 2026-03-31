package services

import (
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
