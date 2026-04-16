package controllers

import (
	"context"

	"github.com/joaomarcosg/Book-Control-System/internal/users/models"
)

type MockUserService struct {
	CreateUserFn  func(ctx context.Context, user *models.User) (int64, error)
	GetUserFn     func(ctx context.Context, id int64) (*models.User, error)
	GetAllUsersFn func(ctx context.Context) ([]*models.User, error)
	UpdateUserFn  func(ctx context.Context, id int64, user *models.User) error
	DeleteUserFn  func(ctx context.Context, id int64) error

	CreateUserCalled  bool
	GetUserCalled     bool
	GetAllUsersCalled bool
	UpdateUserCalled  bool
	DeleteUserCalled  bool
}

func (m *MockUserService) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	m.CreateUserCalled = true
	return m.CreateUserFn(ctx, user)
}

func (m *MockUserService) GetUser(ctx context.Context, id int64) (*models.User, error) {
	m.GetUserCalled = true
	return m.GetUserFn(ctx, id)
}

func (m *MockUserService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	m.GetAllUsersCalled = true
	return m.GetAllUsersFn(ctx)
}

func (m *MockUserService) UpdateUser(ctx context.Context, id int64, user *models.User) error {
	m.UpdateUserCalled = true
	return m.UpdateUserFn(ctx, id, user)
}

func (m *MockUserService) DeleteUser(ctx context.Context, id int64) error {
	m.DeleteUserCalled = true
	return m.DeleteUserFn(ctx, id)
}
