-- name: CreatePost :many
INSERT INTO posts(id, created_at, updated_at, title, url, description, published_at, feed_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
RETURNING *;

-- name: GetPostByUser :many
SELECT posts.* FROM posts
INNER JOIN feed_follows on posts.feed_id = feed_follows.feed_id
WHERE feed_follows.user_id = $1
LIMIT $2;




