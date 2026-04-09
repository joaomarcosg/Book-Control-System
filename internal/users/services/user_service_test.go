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
}

func (m *MockUserRepository) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	return m.CreateUserFn(ctx, user)
}

func (m *MockUserRepository) GetUser(ctx context.Context, id int64) (*models.User, error) {
	return m.GetUserFn(ctx, id)
}

func (m *MockUserRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return m.GetAllUsersFn(ctx)
}

func (m *MockUserRepository) UpdateUser(ctx context.Context, id int64, user *models.User) error {
	return m.UpdateUserFn(ctx, id, user)
}

func (m *MockUserRepository) DeleteUser(ctx context.Context, id int64) error {
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

	if !errors.Is(err, repositories.ErrDuplicateUserNameOrEmail) {
		t.Fatalf("expected duplicate error, got %v", err)
	}

}

func TestGetUser_Success(t *testing.T) {
	expectedUser := &models.User{
		ID:        1,
		Name:      "John Doe",
		Email:     "johndoe@email.com",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
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

	if !errors.Is(err, repositories.ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got %v", err)
	}

	if user != nil {
		t.Fatalf("expected nil, got %v", user)
	}

}

func TestGetAllUsers_Success(t *testing.T) {
	expectedUsers := []*models.User{
		{
			ID:        1,
			Name:      "John Doe",
			Email:     "johndoe@email.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		},
		{
			ID:        2,
			Name:      "Anne Frank",
			Email:     "annefranke@email.com",
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
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

	if len(users) < len(expectedUsers) {
		t.Fatalf("expected %v, got %v", expectedUsers, users)
	}
}
