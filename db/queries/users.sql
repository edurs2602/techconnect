-- name: CreateUser :one
INSERT INTO users (username, email, password)
VALUES ($1, $2, $3)
RETURNING *;

-- name: GetUserByEmail :one
SELECT * FROM users
WHERE email = $1;

-- name: GetUserByUsername :one
SELECT * FROM users 
WHERE username = $1;

-- name: ExistsByEmail :one
SELECT EXISTS(
    SELECT 1 FROM users WHERE email = $1
) AS exists;

-- name: ExistsByUsername :one
SELECT EXISTS(
    SELECT 1 FROM users WHERE username = $1
) AS exists;


-- name: UpdateUser :one
UPDATE users
SET
    username = COALESCE($2, username),
    bio = COALESCE($3, bio)
WHERE id = $1
RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;