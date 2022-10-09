-- name: InsertNewWorkJunc :exec
INSERT INTO Exercise_Workout_Junction (
    junction_id,
    exercise_id,
    workout_id,
    user_id
) VALUES ($1, $2, $3, $4)
ON CONFLICT (exercise_id, workout_id, user_id) DO NOTHING;

-- name: RemoveOldWorkJunc :exec
DELETE FROM Exercise_Workout_Junction
WHERE workout_id = $1 AND NOT (exercise_id = ANY($2::int[]));

-- name: GetExerWorkJunc :many
SELECT exercise_name, exercise_type_id, workout_name, workout_type
FROM Exercise_Workout_Junction 
JOIN Exercise ON Exercise.id = Exercise_Workout_Junction.exercise_id
JOIN Workout ON Workout.id = Exercise_Workout_Junction.workout_id
WHERE workout_id = $1 AND Exercise_Workout_Junction.user_id = $2;

-- name: DeleteWorkoutJunction :exec
DELETE FROM Exercise_Workout_Junction WHERE junction_id = $1 AND Exercise_Workout_Junction.user_id = $2;