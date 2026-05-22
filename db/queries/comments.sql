-- name: CreateComment :one
INSERT INTO comments (post_id, user_id, content)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetCommentsByPostID :many
SELECT * FROM comments WHERE post_id = $1 ORDER BY created_at ASC;

-- name: DeleteComment :exec
DELETE FROM comments WHERE id = $1;