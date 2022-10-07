package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestWorkout(t *testing.T) Workout {
	user := createTestingUser(t)

	arg := CreateWorkoutParams{
		WorkoutName: "Upper Push",
		WorkoutType: []string{"Chest, Triceps, Shoulders"},
		UserID:      user.UserName, // make this object slice
	}

	workout, err := testQueries.CreateWorkout(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, workout)

	require.Equal(t, arg.WorkoutName, workout.WorkoutName)
	require.Equal(t, arg.WorkoutType, workout.WorkoutType)

	return workout

}

func deleteTestWorkout(t *testing.T, w Workout) {
	err := testQueries.DeleteExercise(context.Background(), w.ID)

	require.NoError(t, err)
	testQueries.DeleteUser(context.Background(), w.UserID)
}

func TestCreateWorkout(t *testing.T) {
	e := createTestWorkout(t)
	deleteTestWorkout(t, e)
}

func TestGetWorkoutByName(t *testing.T) {
	work1 := createTestWorkout(t)

	arg := GetWorkoutNameParams{
		UserID:      work1.UserID,
		WorkoutName: work1.WorkoutName,
	}

	work2, err := testQueries.GetWorkoutName(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, work2)

	require.Equal(t, work1.WorkoutName, work2.WorkoutName)
	require.Equal(t, work1.WorkoutType, work2.WorkoutType)

	deleteTestWorkout(t, work1)
}

func TestGetWorkouts(t *testing.T) {
	work1 := createTestWorkout(t)
	//Create 2nd Workout
	arg2 := CreateWorkoutParams{
		WorkoutName: "Legs",
		WorkoutType: []string{"Quads, Hamstrings, Calves"},
		UserID:      work1.UserID, // make this object slice
	}

	work2, err := testQueries.CreateWorkout(context.Background(), arg2)
	require.NoError(t, err)
	require.NotEmpty(t, work2)

	// Grab all Workouts
	workouts, err := testQueries.GetWorkouts(context.Background(), work1.UserID)
	require.NoError(t, err)
	require.NotEmpty(t, workouts)

	require.Equal(t, work1.WorkoutName, workouts[0].WorkoutName)
	require.Equal(t, work1.WorkoutType, workouts[0].WorkoutType)
	require.Equal(t, work2.WorkoutName, workouts[1].WorkoutName)
	require.Equal(t, work2.WorkoutType, workouts[1].WorkoutType)

	deleteTestWorkout(t, work1)
}

func TestUpdateWorkout(t *testing.T) {
	work1 := createTestWorkout(t)

	arg := UpdateWorkoutParams{
		WorkoutName:    "Legs",
		WorkoutType: []string{"Quads, Hamstrings, Calves"},
		ID:             work1.ID,
	}

	get := GetWorkoutNameParams{
		UserID:      work1.UserID,
		WorkoutName: arg.WorkoutName,
	}

	err := testQueries.UpdateWorkout(context.Background(), arg)
	require.NoError(t, err)

	work2, err := testQueries.GetWorkoutName(context.Background(), get)

	require.NoError(t, err)
	require.NotEmpty(t, work2)

	require.Equal(t, arg.WorkoutName, work2.WorkoutName)
	require.Equal(t, arg.WorkoutType, work2.WorkoutType)

	deleteTestWorkout(t, work1)

}

func TestDeleteWorkout(t *testing.T) {
	e := createTestExercise(t)
	deleteTestExercise(t, e)
}
