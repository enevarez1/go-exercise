package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateUser(t *testing.T) {
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
}

func TestGetUser(t *testing.T) {
	user, err := testQueries.GetUser(context.Background(), 1)

	require.NoError(t, err)
	require.NotEmpty(t, user)

	require.Equal(t, "testy", user.UserName)
	require.Equal(t, "Test Mcgee", user.FullName)
	require.Equal(t, "test@test.com", user.Email)
	require.Equal(t, "testSecret", user.Password)
}

func TestUpdateUser(t *testing.T) {
	arg := UpdateUserParams {
		FullName: "TestUpdate",
		Password: "updatePassword",
		ID: 1,
	}

	err := testQueries.UpdateUser(context.Background(), arg)
	require.NoError(t, err)

	// Grab so we can see if it updated
	user, getErr := testQueries.GetUser(context.Background(),1)
	require.NoError(t, getErr)
	require.NotEmpty(t, user)

	require.Equal(t, "TestUpdate", user.FullName)
	require.Equal(t, "updatePassword", user.Password)
}

func TestDeleteUser(t *testing.T) {
	err := testQueries.DeleteUser(context.Background(), 1)
	require.NoError(t, err)

	user, err := testQueries.GetUser(context.Background(),1)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, user)
}