package services

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/joaomarcosg/Book-Control-System/internal/users/models"
	"github.com/joaomarcosg/Book-Control-System/internal/users/repositories"
)

type MockUserRepository struct {
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

func (m *MockUserRepository) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	m.CreateUserCalled = true
	return m.CreateUserFn(ctx, user)
}

func (m *MockUserRepository) GetUser(ctx context.Context, id int64) (*models.User, error) {
	m.GetUserCalled = true
	return m.GetUserFn(ctx, id)
}

func (m *MockUserRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	m.GetAllUsersCalled = true
	return m.GetAllUsersFn(ctx)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, id int64, user *models.User) error {
	m.UpdateUserCalled = true
	return m.UpdateUserFn(ctx, id, user)
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, id int64) error {
	m.DeleteUserCalled = true
	return m.DeleteUserFn(ctx, id)
}

func TestCreateUser_Success(t *testing.T) {
	var expectedID int64 = 1

	mockUserRepository := &MockUserRepository{
		CreateUserFn: func(ctx context.Context, user *models.User) (int64, error) {
			return int64(expectedID), nil
		},
	}

	service := NewUserService(mockUserRepository)

	newUser := &models.User{
		Name:  "John Doe",
		Email: "johndoe@email.com",
	}

	id, err := service.CreateUser(context.Background(), newUser)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if !mockUserRepository.CreateUserCalled {
		t.Errorf("expected CreateUser to be called")
	}

	if id != expectedID {
		t.Fatalf("expected %v, got %v", expectedID, id)
	}
}

func TestCreateUser_Duplicate(t *testing.T) {

	mockUserRepository := &MockUserRepository{
		CreateUserFn: func(ctx context.Context, user *models.User) (int64, error) {
			return 0, repositories.ErrDuplicateUserNameOrEmail
		},
	}

	service := NewUserService(mockUserRepository)

	newUser := &models.User{
		Name:  "John Doe",
		Email: "johndoe@email.com",
	}

	_, err := service.CreateUser(context.Background(), newUser)

	if !mockUserRepository.CreateUserCalled {
		t.Errorf("expected CreateUser to be called")
	}

	if !errors.Is(err, repositories.ErrDuplicateUserNameOrEmail) {
		t.Fatalf("expected duplicate error, got %v", err)
	}

}

func TestGetUser_Success(t *testing.T) {
	fixedTime := time.Date(2026, 4, 12, 14, 0, 0, 0, time.UTC)

	expectedUser := &models.User{
		ID:        1,
		Name:      "John Doe",
		Email:     "johndoe@email.com",
		CreatedAt: fixedTime,
		UpdatedAt: fixedTime,
	}

	mockUserRepository := &MockUserRepository{
		GetUserFn: func(ctx context.Context, id int64) (*models.User, error) {
			return &models.User{
				ID:        expectedUser.ID,
				Name:      expectedUser.Name,
				Email:     expectedUser.Email,
				CreatedAt: expectedUser.CreatedAt,
				UpdatedAt: expectedUser.UpdatedAt,
			}, nil
		},
	}

	service := NewUserService(mockUserRepository)

	user, err := service.GetUser(context.Background(), 1)

	if err != nil {
		t.Fatalf("unexpected erro %v", err)
	}

	if !mockUserRepository.GetUserCalled {
		t.Errorf("expected GetUser to be called")
	}

	if !reflect.DeepEqual(user, expectedUser) {
		t.Fatalf("expected %v, got %v", expectedUser, user)
	}
}

func TestGetUser_UserNotFound(t *testing.T) {
	var id int64 = 1
	mockUserRepository := &MockUserRepository{
		GetUserFn: func(ctx context.Context, id int64) (*models.User, error) {
			return nil, repositories.ErrUserNotFound
		},
	}

	service := NewUserService(mockUserRepository)

	user, err := service.GetUser(context.Background(), id)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !mockUserRepository.GetUserCalled {
		t.Errorf("expected GetUser to be called")
	}

	if !errors.Is(err, repositories.ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got %v", err)
	}

	if user != nil {
		t.Fatalf("expected nil, got %v", user)
	}

}

func TestGetAllUsers_ShouldReturnUsers_WhenRepositorySucceeds(t *testing.T) {
	fixedTime := time.Date(2026, 4, 12, 14, 0, 0, 0, time.UTC)

	expectedUsers := []*models.User{
		{
			ID:        1,
			Name:      "John Doe",
			Email:     "johndoe@email.com",
			CreatedAt: fixedTime,
			UpdatedAt: fixedTime,
		},
		{
			ID:        2,
			Name:      "Anne Frank",
			Email:     "annefranke@email.com",
			CreatedAt: fixedTime,
			UpdatedAt: fixedTime,
		},
	}

	mockUserRepository := &MockUserRepository{
		GetAllUsersFn: func(ctx context.Context) ([]*models.User, error) {
			return expectedUsers, nil
		},
	}

	service := NewUserService(mockUserRepository)

	users, err := service.GetAllUsers(context.Background())

	if err != nil {
		t.Fatalf("unexpected erro %v", err)
	}

	if !mockUserRepository.GetAllUsersCalled {
		t.Errorf("expected GetAllUsers to be called")
	}

	if !reflect.DeepEqual(users, expectedUsers) {
		t.Errorf("expected %v, got %v", expectedUsers, users)
	}
}
