-- name: CreateUser :one
INSERT INTO users (
    name,
    email
)
VALUES ($1, $2)
RETURNING id, name, email, created_at, updated_at;

-- name: GetUser :one
SELECT id, name, email, created_at
FROM users
WHERE id = $1;

-- name: GetAllUsers :many
SELECT id, name, email, created_at, updated_at
FROM users
ORDER BY id;

-- name: UpdateUser :one
UPDATE users
SET
    name = COALESCE($2, name),
    email = COALESCE($3, email)
WHERE id = $1
RETURNING id, name, email, created_at, updated_at;

-- name: DeleteUser :exec
DELETE FROM users
WHERE id = $1;