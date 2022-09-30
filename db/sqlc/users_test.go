package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func createTestingUser(t *testing.T) User{
	arg := CreateUserParams{
		UserName: "testy",
		FullName: "Test Mcgee",
		Email: "test@test.com",
		Password: "testSecret",
	}

	user, err := testQueries.CreateUser(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, arg.UserName, user.UserName)
	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.Password, user.Password)

	return user
}

func deleteTestingUser(t *testing.T, u User) {
	err := testQueries.DeleteUser(context.Background(), u.ID)
	require.NoError(t, err)

	user, err := testQueries.GetUser(context.Background(),1)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user)

}

func TestCreateUser(t *testing.T) {
	user := createTestingUser(t)
	deleteTestingUser(t, user)
}

func TestGetUser(t *testing.T) {
	rand := createTestingUser(t)
	user, err := testQueries.GetUser(context.Background(), rand.ID)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, rand.UserName, user.UserName)
	require.Equal(t, rand.FullName, user.FullName)
	require.Equal(t, rand.Email, user.Email)
	require.Equal(t, rand.Password, user.Password)

	deleteTestingUser(t, rand)
}

func TestUpdateUser(t *testing.T) {

	rand := createTestingUser(t)

	arg := UpdateUserParams {
		FullName: "TestUpdate",
		Password: "updatePassword",
		ID: rand.ID,
	}

	err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)

	// Grab so we can see if it updated
	user, getErr := testQueries.GetUser(context.Background(),rand.ID)
	require.NoError(t, getErr)
	require.NotEmpty(t, user)

	require.Equal(t, arg.FullName, user.FullName)
	require.Equal(t, arg.Password, user.Password)

	deleteTestingUser(t, rand)
}

func TestDeleteUser(t *testing.T) {
	user := createTestingUser(t)
	deleteTestingUser(t,user)
}