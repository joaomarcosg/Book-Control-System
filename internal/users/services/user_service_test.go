package services

import (
	"context"
	"testing"

	"github.com/joaomarcosg/Book-Control-System/internal/users/models"
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
