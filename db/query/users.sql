-- name: CreateUser :one
INSERT INTO Users (
    user_name,
    full_name,
    email, 
    password
) VALUES (
    $1, $2, $3, $4
)
RETURNING *;

-- name: UpdateUser :exec
UPDATE Users 
SET full_name = $1,
    password = $2
WHERE id = $3;