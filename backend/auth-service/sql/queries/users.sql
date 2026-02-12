-- name: GetUserByID :one
SELECT * from users WHERE id=$1;

-- name: GetUsers :many
SELECT * FROM users;

-- name: CreateUser :one
INSERT INTO users (email, password_hash) VALUES ($1, $2) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id=$1;

-- name: UpdateUser :exec
UPDATE users
SET email = @email,
    password_hash = @password_hash 
WHERE id = $1
RETURNING *;
