-- name: CreateExercise :one
INSERT INTO Exercise (
    exercise_name,
    exercise_type_id,
    user_id
) VALUES (
    $1, $2, $3
)
RETURNING *;

-- name: UpdateExercise :exec
UPDATE Exercise
SET exercise_name = $1,
    exercise_type_id = $2
WHERE id = $3;

-- name: DeleteExercise :exec
DELETE from Exercise where id = $1;

-- name: GatherExercises :many
SELECT * from Exercise where user_id = $1;
