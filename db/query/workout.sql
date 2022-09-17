-- name: CreateWorkout :one
INSERT INTO Workout (
    workout_name,
    workout_type_id
) VALUES (
    $1, $2
)
RETURNING *;

-- name: UpdateWorkout :exec
UPDATE Workout
SET workout_name = $1,
    workout_type_id = $2
WHERE id = $3;

-- name: DeleteWorkout :exec
DELETE FROM Workout WHERE id = $1;

-- name: GetWorkouts :many
SELECT * FROM Workout WHERE user_id = $1;

-- name: GetWorkoutName :many
SELECT * FROM Workout 
    WHERE user_id = $1 AND workout_name = $2; 

-- name: GetWorkoutBasedOnType :many
SELECT * FROM Workout 
    WHERE user_id = $1 AND workout_type_id = $2;