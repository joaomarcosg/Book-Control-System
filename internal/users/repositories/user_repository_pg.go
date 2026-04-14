package repositories

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	db "github.com/joaomarcosg/Book-Control-System/internal/database/sqlc"
	"github.com/joaomarcosg/Book-Control-System/internal/users/models"
)

type userRepository struct {
	queries *db.Queries
}

func NewUserRepository(q *db.Queries) UserRepository {
	return &userRepository{
		queries: q,
	}
}

func (u *userRepository) CreateUser(ctx context.Context, user *models.User) (int64, error) {

	params := toCreateParams(user)
	id, err := u.queries.CreateUser(ctx, params)

	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == "23505" {
			return 0, ErrDuplicateUserNameOrEmail
		}
		return 0, err
	}

	return int64(id), nil
}

func (u *userRepository) GetUser(ctx context.Context, id int64) (*models.User, error) {

	user, err := u.queries.GetUser(ctx, int32(id))
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, ErrUserNotFound
		}
		return nil, err
	}

	return toDomain(user), nil
}

func (u *userRepository) GetAllUsers(ctx context.Context) ([]*models.User, error) {

	users, err := u.queries.GetAllUsers(ctx)

	if err != nil {
		return []*models.User{}, err
	}

	return toDomainList(users), nil

}

func (u *userRepository) UpdateUser(ctx context.Context, id int64, user *models.User) error {

	params := toUpdateParams(id, user)

	_, err := u.queries.UpdateUser(ctx, params)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}

	return nil

}

func (u *userRepository) DeleteUser(ctx context.Context, id int64) error {

	err := u.queries.DeleteUser(ctx, int32(id))

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return ErrUserNotFound
		}
		return err
	}

	return nil

}
