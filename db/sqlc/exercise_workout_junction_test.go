package db

import (
	"context"
	"testing"

	"github.com/stretchr/testify/require"
)

func testSetup(t *testing.T) (User, []Exercise, Workout){
	// Make all the data we need

	// Need One User
	user := createTestingUser(t)

	// Need 3 Exercises
	exercises := make([]Exercise, 3)
	exerArgs := []CreateExerciseParams{
		{
			ExerciseName: "Bench",
			ExerciseTypeID: "Chest",
			UserID: user.UserName,
		},
		{
			ExerciseName: "Deadlift",
			ExerciseTypeID: "Back",
			UserID: user.UserName,
		},
		{
			ExerciseName: "Squat",
			ExerciseTypeID: "Legs",
			UserID: user.UserName,
		},
	}

	for i, arg := range exerArgs {
		exercise, err := testQueries.CreateExercise(context.Background(), arg)
		require.NoError(t, err)
		require.NotEmpty(t, exercise)
		exercises[i] = exercise
	}

	// Need 1 Workout
	workArg := CreateWorkoutParams{
		WorkoutName: "TestWorkout",
		WorkoutType: []string{"Chest", "Back", "Legs"},
		UserID: user.UserName,
	}

	work, err := testQueries.CreateWorkout(context.Background(), workArg)
	require.NoError(t, err)
	require.NotEmpty(t, work)

	//Return data to tests
	return user, exercises, work
}

func testTearDown(t *testing.T, u User) {
	deleteTestingUser(t, u)
}

func TestInsertNewWorkJunc(t *testing.T) {
	user, exercises, workout := testSetup(t)

	defer testTearDown(t, user)

	args := InsertNewWorkJuncParams{
		JunctionID: 1,
		ExerciseID: exercises[0].ID,
		WorkoutID: workout.ID,
		UserID: user.UserName,
	}

	err := testQueries.InsertNewWorkJunc(context.Background(), args)
	require.NoError(t, err)
}

func TestDeleteWorkJunc(t *testing.T) {
	user, exercises, workout := testSetup(t)

	defer testTearDown(t, user)

	args := InsertNewWorkJuncParams{
		JunctionID: 1,
		ExerciseID: exercises[0].ID,
		WorkoutID: workout.ID,
		UserID: user.UserName,
	}

	err := testQueries.InsertNewWorkJunc(context.Background(), args)
	require.NoError(t, err)

	arg := DeleteWorkoutJunctionParams{
		JunctionID: 1,
		UserID: user.UserName,
	}

	err = testQueries.DeleteWorkoutJunction(context.Background(), arg)
	require.NoError(t, err)
}

func TestGetJunc(t *testing.T) {
	user, exercises, workout := testSetup(t)

	defer testTearDown(t, user)

	args := []InsertNewWorkJuncParams{
		{
			JunctionID: 1,
			ExerciseID: exercises[0].ID,
			WorkoutID: workout.ID,
			UserID: user.UserName,
		},
		{
			JunctionID: 2,
			ExerciseID: exercises[1].ID,
			WorkoutID: workout.ID,
			UserID: user.UserName,
		},
		{
			JunctionID: 3,
			ExerciseID: exercises[2].ID,
			WorkoutID: workout.ID,
			UserID: user.UserName,
		},
	}

	for _, arg := range args {
		err := testQueries.InsertNewWorkJunc(context.Background(), arg)
		require.NoError(t, err)
	}

	get := GetExerWorkJuncParams{
		WorkoutID: workout.ID,
		UserID: user.UserName,
	}

	results, err := testQueries.GetExerWorkJunc(context.Background(), get)
	require.NoError(t, err)
	require.NotEmpty(t, results)

	for i, junc := range results {

		require.Equal(t, junc.ExerciseName, exercises[i].ExerciseName)
		require.Equal(t, junc.ExerciseTypeID, exercises[i].ExerciseTypeID)
		require.Equal(t, junc.WorkoutName, workout.WorkoutName)
		require.Equal(t, junc.WorkoutType, workout.WorkoutType)
	}
}

func TestRemoveOldJunc(t *testing.T) {
	user, exercises, workout := testSetup(t)

	defer testTearDown(t, user)

	args := []InsertNewWorkJuncParams{
		{
			JunctionID: 1,
			ExerciseID: exercises[0].ID,
			WorkoutID: workout.ID,
			UserID: user.UserName,
		},
		{
			JunctionID: 2,
			ExerciseID: exercises[1].ID,
			WorkoutID: workout.ID,
			UserID: user.UserName,
		},
		{
			JunctionID: 3,
			ExerciseID: exercises[2].ID,
			WorkoutID: workout.ID,
			UserID: user.UserName,
		},
	}

	for _, arg := range args {
		err := testQueries.InsertNewWorkJunc(context.Background(), arg)
		require.NoError(t, err)
	}

	del := RemoveOldWorkJuncParams{
		WorkoutID: workout.ID,
		Column2: []int32{exercises[1].ID},
	}

	err := testQueries.RemoveOldWorkJunc(context.Background(), del)
	require.NoError(t, err)

	get := GetExerWorkJuncParams{
		WorkoutID: workout.ID,
		UserID: user.UserName,
	}

	results, err := testQueries.GetExerWorkJunc(context.Background(), get)
	require.NoError(t, err)
	require.NotEmpty(t, results)

	require.Equal(t, results[0].ExerciseName, exercises[1].ExerciseName)
	require.Equal(t, results[0].ExerciseTypeID, exercises[1].ExerciseTypeID)
	require.Equal(t, results[0].WorkoutName, workout.WorkoutName)
	require.Equal(t, results[0].WorkoutType, workout.WorkoutType)
}