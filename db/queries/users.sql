-- name: CreateUser :one
INSERT INTO users( name, email, password )
VALUES($1, $2, $3)
RETURNING *;

-- name: LoginUser :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email=$1;


-- name: GetUserById :one
SELECT * FROM users
WHERE id=$1;

-- name: GetAllusers :many
SELECT * FROM users
ORDER BY Created_at desc;

-- name: DeleteUser :exec
DELETE from users
WHERE id = $1;

-- name: UpdateUser :exec
UPDATE users
SET
name = COALESCE (sqlc.narg(name), name),
email = COALESCE (sqlc.narg(email), email),
password = COALESCE (sqlc.narg(password), password)
WHERE id = $1;

