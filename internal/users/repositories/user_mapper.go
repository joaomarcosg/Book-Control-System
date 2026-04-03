package repositories

import (
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/joaomarcosg/Book-Control-System/internal/database/sqlc"
	"github.com/joaomarcosg/Book-Control-System/internal/users/models"
)

func mapToDomain(
	id int32,
	name,
	email string,
	createdAt,
	updatedAt pgtype.Timestamptz,
) *models.User {
	return &models.User{
		ID:        int64(id),
		Name:      name,
		Email:     email,
		CreatedAt: createdAt.Time,
		UpdatedAt: updatedAt.Time,
	}
}

func toCreateParams(u *models.User) sqlc.CreateUserParams {
	return sqlc.CreateUserParams{
		Name:  u.Name,
		Email: u.Email,
	}
}

func toDomain(u sqlc.User) *models.User {
	return mapToDomain(
		u.ID,
		u.Name,
		u.Email,
		u.CreatedAt,
		u.UpdatedAt,
	)
}

func toDomainList(users []sqlc.User) []*models.User {
	var result []*models.User

	for _, u := range users {
		result = append(result, toDomain(u))
	}

	return result
}
