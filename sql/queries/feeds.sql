-- name: CreateFeed :one
INSERT INTO feeds(id, created_at, updated_at, name, url, user_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeed :many
SELECT * FROM feeds;

-- name: GetNextFeedsToFetch :many
SELECT * FROM feeds
ORDER BY last_fetch_at ASC NULLS FIRST;

-- name: MarkFeedFetch :exec
UPDATE feeds
SET last_fetch_at = $1, updated_at = $2
WHERE id = $3;