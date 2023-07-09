-- name: CreateFeedFollow :one
INSERT INTO feed_followes (id, created_at, updated_at, user_id, feed_id)
VALUES ($1, $2, $3, $4, $5) 
RETURNING *;

-- name: GetFeedFollowes :many
SELECT * FROM feed_followes WHERE user_id = $1;

-- name: DeleteFeedFollow :exec
DELETE from feed_followes WHERE id = $1 AND user_id = $2;