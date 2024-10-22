-- name: CreateFeedFollow :one
WITH new_feed_follow AS (
    INSERT INTO feed_follows (id, feed_id, user_id, created_at, updated_at)
    VALUES (
        $1,
        $2,
        $3,
        $4,
        $5
    )
    RETURNING *
)
SELECT ff.*, f.name as feed_name, u.name as user_name
FROM new_feed_follow ff
JOIN feeds f ON ff.feed_id = f.id
JOIN users u ON ff.user_id = u.id;

-- name: GetFeedFollowsForUser :many
SELECT ff.*, f.name as feed_name, u.name as user_name
FROM feed_follows ff
JOIN feeds f ON ff.feed_id = f.id
JOIN users u ON ff.user_id = u.id
WHERE u.name = $1;

-- name: DeleteFeedFollow :exec
DELETE FROM feed_follows 
WHERE user_id IN (
    SELECT id 
    FROM users u 
    WHERE u.name = $1
) 
AND feed_id IN (
    SELECT id 
    FROM feeds f 
    WHERE f.url = $2
);