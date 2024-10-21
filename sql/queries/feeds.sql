-- name: AddFeed :one
INSERT INTO feeds (id, name, url, created_at, updated_at, user_id)
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
)
RETURNING *;

-- name: ListFeeds :many
SELECT f.name, f.url, u.name as user_name
FROM feeds f
JOIN users u ON f.user_id = u.id;
