package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestWorkout(t *testing.T) Workout {
	user := createTestingUser(t)
	
	arg := CreateWorkoutParams {
		WorkoutName: "Upper Push",
		WorkoutTypeID: "Chest, Triceps, Shoulders", // make this object slice
	}

	workout, err := testQueries.CreateWorkout(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, workout)

	require.Equal(t, arg.WorkoutName, workout.WorkoutName)
	require.Equal(t, arg.WorkoutTypeID, workout.WorkoutTypeID)

	return workout

}

func deleteTestWorkout(t *testing.T, e Exercise) {
	err := testQueries.DeleteExercise(context.Background(), e.ID)

	require.NoError(t, err)
	testQueries.DeleteUser(context.Background(), e.ID)
}

func TestCreateWorkout(t *testing.T) {
	e := createTestExercise(t)
	deleteTestExercise(t, e)
}

func TestGatherWorkout(t *testing.T) {
	exer1 := createTestExercise(t)
	exer2, err := testQueries.GatherExercises(context.Background(), exer1.UserID)

	require.NoError(t, err)
	require.NotEmpty(t, exer2)

	require.Equal(t, exer1.ExerciseName, exer2[0].ExerciseName)
	require.Equal(t, exer1.ExerciseTypeID, exer2[0].ExerciseTypeID)

	deleteTestExercise(t, exer1)
}

func TestUpdateWorkout(t *testing.T) {
	exer1 := createTestExercise(t)

	arg := UpdateExerciseParams {
		ExerciseName: "Squat",
		ExerciseTypeID: "Legs",
		ID:	exer1.ID,
	}

	err := testQueries.UpdateExercise(context.Background(), arg)
	require.NoError(t, err)

	exer2, err := testQueries.GatherExercises(context.Background(), exer1.UserID)

	require.NoError(t, err)
	require.NotEmpty(t, exer2)

	require.Equal(t, arg.ExerciseName, exer2[0].ExerciseName)
	require.Equal(t, arg.ExerciseTypeID, exer2[0].ExerciseTypeID)

	deleteTestExercise(t, exer1)

}

func TestDeleteWorkout(t *testing.T) {
	e := createTestExercise(t)
	deleteTestExercise(t, e)
}