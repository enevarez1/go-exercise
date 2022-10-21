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

-- name: GetUser :one
SELECT * FROM Users WHERE id = $1;

-- name: UpdateUser :exec
UPDATE Users 
SET user_name = $1,
    full_name = $2,
    password = $3
WHERE id = $4;

-- name: DeleteUser :exec
DELETE FROM Users WHERE id = $1;