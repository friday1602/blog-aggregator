-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, name, api_key)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'))
RETURNING api_key;

-- name: GetUserByAPI :one
SELECT * FROM users
WHERE api_key = $1;
