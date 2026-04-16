package controllers

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joaomarcosg/Book-Control-System/internal/users/models"
)

type MockUserService struct {
	CreateUserFn  func(ctx context.Context, user *models.User) (int64, error)
	GetUserFn     func(ctx context.Context, id int64) (*models.User, error)
	GetAllUsersFn func(ctx context.Context) ([]*models.User, error)
	UpdateUserFn  func(ctx context.Context, id int64, user *models.User) error
	DeleteUserFn  func(ctx context.Context, id int64) error
}

func (m *MockUserService) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	return m.CreateUserFn(ctx, user)
}

func (m *MockUserService) GetUser(ctx context.Context, id int64) (*models.User, error) {

	return m.GetUserFn(ctx, id)
}

func (m *MockUserService) GetAllUsers(ctx context.Context) ([]*models.User, error) {
	return m.GetAllUsersFn(ctx)
}

func (m *MockUserService) UpdateUser(ctx context.Context, id int64, user *models.User) error {
	return m.UpdateUserFn(ctx, id, user)
}

func (m *MockUserService) DeleteUser(ctx context.Context, id int64) error {
	return m.DeleteUserFn(ctx, id)
}

func setupRouter(controller *UserController) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	controller.RegisterRoutes(r)
	return r
}

func TestCreateUser_ShouldReturn201_WhenDataIsValid(t *testing.T) {

	mockService := &MockUserService{
		CreateUserFn: func(ctx context.Context, user *models.User) (int64, error) {
			return 1, nil
		},
	}

	controller := NewUserController(mockService)
	router := setupRouter(controller)

	body := `{"name":"John Doe","email":"john@email.com"}`
	req, _ := http.NewRequest(http.MethodPost, "/users", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Fatalf("expected status 201, got %d", w.Code)
	}
}
