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
SELECT * FROM Users WHERE user_name = $1;

-- name: UpdateUser :exec
UPDATE Users 
SET full_name = $1,
    password = $2
WHERE user_name = $3;

-- name: DeleteUser :exec
DELETE FROM Users WHERE user_name = $1;