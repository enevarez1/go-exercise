-- name CreateExerWorkJunc :one
INSERT INTO Exercise_Workout_Junction (
    exercise_id,
    workout_id,
    user_id
) VALUES (
    $1, $2, $3
);

-- name UpdateExerWorkJunc :exec
UPDATE Exercise_Workout_Junction
SET exercise_id = $1,
    workout_id = $2
WHERE id = $3;

-- name DeleteExerWorkJunc :exec
DELETE FROM Exercise_Workout_Junction WHERE id = $1