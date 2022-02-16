-- name: CreateUser :one
INSERT INTO users (
  username,
  password,
  first_name,
  last_name,
  email
) VALUES (
  $1, $2, $3, $4, $5
)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;