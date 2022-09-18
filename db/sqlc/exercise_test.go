package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateExercise(t *testing.T) {
	arg := CreateExerciseParams{
		ExerciseName: 	"Bench",
		ExerciseTypeID: "Chest",
		UserID:	100,
	}

	exercise, err := testQueries.CreateExercise(context.Background(), arg)
	fmt.Print(exercise)
	require.NoError(t, err)
	require.NotEmpty(t, exercise)

	// require.Equal(t, arg.ExerciseName, exercise.ExerciseName)
	// require.Equal(t, arg.ExerciseTypeID, exercise.ExerciseTypeID)
	// require.Equal(t, arg.UserID, exercise.UserID)

	// require.NotZero(t, exercise.ID)

}