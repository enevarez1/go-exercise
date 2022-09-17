-- name CreateWorkoutType :one
INSERT INTO Workout_Type (
    workout_type
) VALUES (
    $1
);

-- name UpdateWorkoutType :one
UPDATE Workout_Type
SET workout_type = $1
WHERE id = $3;

-- name DeleteWorkoutType :exec
DELETE from Workout_Type where workout_type = $1;

SELECT * FROM Workout_Type WHERE user_id = $1
