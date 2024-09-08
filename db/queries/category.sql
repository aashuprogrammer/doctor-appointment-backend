-- name: CreateCategory :one
INSERT INTO category(category_name)
VALUES ($1)
RETURNING *;

-- name: GetAllCategory :many
SELECT * FROM category;