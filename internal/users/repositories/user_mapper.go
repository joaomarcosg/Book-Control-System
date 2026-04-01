package repositories

import (
	"github.com/joaomarcosg/Book-Control-System/internal/database/sqlc"
	"github.com/joaomarcosg/Book-Control-System/internal/users/models"
)

func toDomainFromCreate(u sqlc.CreateUserRow) *models.User {
	return &models.User{
		ID:        int64(u.ID),
		Name:      u.Name,
		Email:     u.Email,
		CreatedAt: u.CreatedAt.Time,
	}
}

func toCreateParams(u *models.User) sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Name:  u.Name,
		Email: u.Email,
	}
}
