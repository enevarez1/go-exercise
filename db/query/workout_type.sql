-- name: CreateWorkoutType :one
INSERT INTO Workout_Type (
    workout_type
) VALUES (
    $1
)
RETURNING *;

-- name: UpdateWorkoutType :exec
UPDATE Workout_Type
SET workout_type = $1
WHERE workout_type = $2 AND user_id = $3;

-- name: DeleteWorkoutType :exec
DELETE from Workout_Type where workout_type = $1;

-- name: GetWorkoutTypes :many
SELECT * FROM Workout_Type WHERE user_id = $1;
