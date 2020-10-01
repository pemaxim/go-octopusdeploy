package integration

import (
	"testing"

	"github.com/OctopusDeploy/go-octopusdeploy/model"
	"github.com/stretchr/testify/assert"
)

func TestUsers(t *testing.T) {
	t.Run("GetAll", TestUsersGetAll)
	t.Run("GetMe", TestUsersGetMe)
}

func TestUsersGetAll(t *testing.T) {
	octopusClient := getOctopusClient()

	users, err := octopusClient.Users.GetAll()

	assert.NoError(t, err)
	assert.NotEmpty(t, users)

	if users == nil {
		return
	}
}

func TestUsersGetMe(t *testing.T) {
	octopusClient := getOctopusClient()

	user, err := octopusClient.Users.GetMe()

	assert.NoError(t, err)
	if !assert.NotEmpty(t, user) {
		return
	}

	assert.True(t, user.IsActive)
	assert.False(t, user.IsService)
	assert.NotEmpty(t, user.EmailAddress)
}

func TestGetUserByID(t *testing.T) {
	octopusClient := getOctopusClient()

	user, err := octopusClient.Users.GetMe()

	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	if user == nil {
		return
	}

	userToVerify, err := octopusClient.Users.GetByID(user.ID)

	assert.NoError(t, err)
	assert.NotEmpty(t, userToVerify)

	if userToVerify == nil {
		return
	}

	// TODO: add more asserts here
}

func TestGetSpaces(t *testing.T) {
	octopusClient := getOctopusClient()

	user, err := octopusClient.Users.GetMe()

	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	if user == nil {
		return
	}

	spaces, err := octopusClient.Users.GetSpaces(user)

	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	if spaces == nil {
		return
	}

	// TODO: add more asserts here
}

func TestGetAuthenticationForUser(t *testing.T) {
	octopusClient := getOctopusClient()

	user, err := octopusClient.Users.GetMe()

	assert.NoError(t, err)
	assert.NotEmpty(t, user)

	if user == nil {
		return
	}

	authentication, err := octopusClient.Users.GetAuthenticationForUser(user)

	assert.NoError(t, err)
	assert.NotEmpty(t, authentication)

	if authentication == nil {
		return
	}

	// TODO: add more asserts here
}

func TestGetAuthentication(t *testing.T) {
	octopusClient := getOctopusClient()

	authentication, err := octopusClient.Users.GetAuthentication()

	assert.NoError(t, err)
	assert.NotEmpty(t, authentication)

	if authentication == nil {
		return
	}

	// TODO: add more asserts here
}

func TestCreateUser(t *testing.T) {
	octopusClient := getOctopusClient()

	user := model.NewUser(getRandomName(), getRandomName())
	user.Password = getRandomName()
	user, err := octopusClient.Users.Add(user)

	assert.NoError(t, err)
	assert.NotNil(t, user)

	if user == nil {
		return
	}

	assert.True(t, user.IsActive)
	assert.False(t, user.IsService)
	assert.Empty(t, user.EmailAddress)
}

func TestCreateServiceUser(t *testing.T) {
	octopusClient := getOctopusClient()

	user := model.NewUser(getRandomName(), getRandomName())
	user.IsService = true
	user, err := octopusClient.Users.Add(user)

	if user == nil {
		return
	}

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.True(t, user.IsActive)
	assert.True(t, user.IsService)
	assert.Empty(t, user.EmailAddress)
}
