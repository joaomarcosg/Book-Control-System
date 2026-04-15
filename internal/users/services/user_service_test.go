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

func TestCreateUser_ShouldCreateUser_WhenDataIsValid(t *testing.T) {
	var expectedID int64 = 1

	mockUserRepository := &MockUserRepository{
		CreateUserFn: func(ctx context.Context, user *models.User) (int64, error) {
			if user.Name != "John Doe" {
				t.Errorf("unexpected user name: %v", user.Name)
			}
			if user.Email != "johndoe@email.com" {
				t.Errorf("unexpected user email: %v", user.Email)
			}
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
		t.Errorf("expected %v, got %v", expectedID, id)
	}
}

func TestCreateUser_ShouldReturnDuplicateUserError_WhenUserAlreadyExists(t *testing.T) {

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

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !mockUserRepository.CreateUserCalled {
		t.Errorf("expected CreateUser to be called")
	}

	if !errors.Is(err, repositories.ErrDuplicateUserNameOrEmail) {
		t.Fatalf("expected duplicate error, got %v", err)
	}

}

func TestGetUser_ShouldReturnUser_WhenUserIdIsValid(t *testing.T) {
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
			if id != 1 {
				t.Errorf("expected id 1, got %v", id)
			}
			return expectedUser, nil
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

func TestGetUser_ShouldRetornUserNotFoundError_WhenUserIsNotFound(t *testing.T) {
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
		t.Fatalf("expected nil user when error occurs, got %v", user)
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

func TestGetAllUsers_ShouldReturnEmptyList_WhenUsersNotFound(t *testing.T) {
	mockUserRepository := &MockUserRepository{
		GetAllUsersFn: func(ctx context.Context) ([]*models.User, error) {
			if ctx == nil {
				t.Error("expected non-nil context")
			}
			return []*models.User{}, nil
		},
	}

	service := NewUserService(mockUserRepository)

	users, err := service.GetAllUsers(context.Background())

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if users == nil {
		t.Fatalf("expected empty slice, got nil")
	}

	if len(users) != 0 {
		t.Fatalf("expected empty list, got %v", users)
	}

	if !mockUserRepository.GetAllUsersCalled {
		t.Errorf("expected GetAllUsers to be called")
	}

}

func TestUpdateUser_ShouldUpdateUser_WhenDataIsValid(t *testing.T) {
	var id int64 = 1

	updatedUser := &models.User{
		Name:  "John Doe",
		Email: "john.doe@email.com",
	}

	mockUserRepository := &MockUserRepository{
		UpdateUserFn: func(ctx context.Context, receivedID int64, user *models.User) error {
			if ctx == nil {
				t.Error("expected non-nil context")
			}
			if receivedID != id {
				t.Errorf("expected id %v, got %v", id, receivedID)
			}
			if user.Name != updatedUser.Name {
				t.Errorf("expected name %v, got %v", updatedUser.Name, user.Name)
			}
			if user.Email != updatedUser.Email {
				t.Errorf("expected email %v, got %v", updatedUser.Email, user.Email)
			}
			return nil
		},
	}

	service := NewUserService(mockUserRepository)

	err := service.UpdateUser(context.Background(), id, updatedUser)

	if err != nil {
		t.Fatalf("unexpected error %v", err)
	}

	if !mockUserRepository.UpdateUserCalled {
		t.Error("expected UpdateUser to be called")
	}
}

func TestUpdateUser_ShouldReturnUserNotFoundError_WhenUserNotFound(t *testing.T) {

	var id int64 = 1

	updatedUser := &models.User{
		Name:  "John Doe",
		Email: "john.doe@email.com",
	}

	mockUserRepository := &MockUserRepository{
		UpdateUserFn: func(ctx context.Context, id int64, user *models.User) error {
			return repositories.ErrUserNotFound
		},
	}

	service := NewUserService(mockUserRepository)

	err := service.UpdateUser(context.Background(), id, updatedUser)

	if err == nil {
		t.Fatalf("expected error, got nil")
	}

	if !errors.Is(err, repositories.ErrUserNotFound) {
		t.Fatalf("expected ErrUserNotFound, got %v", err)
	}

	if !mockUserRepository.UpdateUserCalled {
		t.Error("expected UpdateUser to be called")
	}

}
