package entities_test

import (
	"github.com/KPMGE/go-users-clean-api/src/domain/entities"
	"github.com/stretchr/testify/require"
	"testing"
)

const fakeName string = "any_name"
const fakeUserName string = "any_user_name"
const fakeEmail string = "any_valid_email@gmail.com"

func TestUser_New(t *testing.T) {
	newUser, err := entities.NewUser(fakeName, fakeUserName, fakeEmail)

	require.Nil(t, err)
	require.Nil(t, newUser.Books)
	require.NotNil(t, newUser)
	require.NotNil(t, newUser.ID)
	require.NotNil(t, newUser.UpdatedAt)
	require.NotNil(t, newUser.CreatedAt)
	require.Equal(t, fakeName, newUser.Name)
	require.Equal(t, fakeUserName, newUser.UserName)
	require.Equal(t, fakeEmail, newUser.Email)
}

func TestUser_New_WithInvalidEmail(t *testing.T) {
	newUser, err := entities.NewUser(fakeName, fakeUserName, "any_invalid_email")

	require.Error(t, err)
	require.Nil(t, newUser)
	require.Equal(t, err.Error(), "Invalid email!")
}

func TestUser_New_WithNullFields(t *testing.T) {
	newUser, err := entities.NewUser("", fakeUserName, fakeEmail)
	require.Error(t, err)
	require.Nil(t, newUser)

	newUser, err = entities.NewUser(fakeName, "", fakeEmail)
	require.Error(t, err)
	require.Nil(t, newUser)

	newUser, err = entities.NewUser(fakeName, fakeUserName, "")
	require.Error(t, err)
	require.Nil(t, newUser)
}
