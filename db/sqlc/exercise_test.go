package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestExercise(t *testing.T) Exercise {
	user := createTestingUser(t)

	arg := CreateExerciseParams {
		ExerciseName: "Bench",
		ExerciseTypeID: "Chest",
		UserID: user.ID,
	}

	exercise, err := testQueries.CreateExercise(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, exercise)

	require.Equal(t, arg.ExerciseName, exercise.ExerciseName)
	require.Equal(t, arg.ExerciseTypeID, exercise.ExerciseTypeID)

	return exercise

}

func deleteTestExercise(t *testing.T, e Exercise) {
	err := testQueries.DeleteExercise(context.Background(), e.ID)

	require.NoError(t, err)
	testQueries.DeleteUser(context.Background(), e.UserID)
}

func TestCreateExercise(t *testing.T) {
	e := createTestExercise(t)
	deleteTestExercise(t, e)
}

func TestGatherExercises(t *testing.T) {
	exer1 := createTestExercise(t)
	exer2, err := testQueries.GatherExercises(context.Background(), exer1.UserID)

	require.NoError(t, err)
	require.NotEmpty(t, exer2)

	require.Equal(t, exer1.ExerciseName, exer2[0].ExerciseName)
	require.Equal(t, exer1.ExerciseTypeID, exer2[0].ExerciseTypeID)

	deleteTestExercise(t, exer1)
}

func TestUpdateExercise(t *testing.T) {
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

func TestDeleteExercise(t *testing.T) {
	e := createTestExercise(t)
	deleteTestExercise(t, e)
}
