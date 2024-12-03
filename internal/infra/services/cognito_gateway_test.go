package services

import (
	"tech-challenge-fase-1/internal/tests/mocks"
	"testing"

	cognito "github.com/aws/aws-sdk-go/service/cognitoidentityprovider"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetUserBySub(t *testing.T) {

	client := mocks.NewMockCognitoClientInterface(t)

	cognitoClient := NewCognitoClient(client, "")

	sub := "sub"
	userName := "Um nome qualquer"
	email := "email"
	name := "name"

	output := &cognito.ListUsersOutput{
		Users: []*cognito.UserType{
			{
				Username: &userName,
				Attributes: []*cognito.AttributeType{
					{Name: &email, Value: &email},
					{Name: &sub, Value: &sub},
					{Name: &name, Value: &name},
				},
			},
		},
	}

	client.On("ListUsers", mock.Anything).Return(output, nil).Once()

	response, err := cognitoClient.GetUserBySub(sub)

	assert.Nil(t, err)
	assert.Equal(t, response.Id, sub)

}

func TestGetUserByName(t *testing.T) {

	client := mocks.NewMockCognitoClientInterface(t)

	cognitoClient := NewCognitoClient(client, "")

	sub := "sub"
	userName := "Um nome qualquer"
	email := "email"
	name := "name"

	output := &cognito.AdminGetUserOutput{
		Username: &userName,
		UserAttributes: []*cognito.AttributeType{
			{Name: &email, Value: &email},
			{Name: &sub, Value: &sub},
			{Name: &name, Value: &name},
		},
	}

	client.On("AdminGetUser", mock.Anything).Return(output, nil).Once()

	response, err := cognitoClient.GetUser(sub)

	assert.Nil(t, err)
	assert.Equal(t, response.Id, sub)

}

func TestCreateUser(t *testing.T) {

	client := mocks.NewMockCognitoClientInterface(t)

	cognitoClient := NewCognitoClient(client, "")

	sub := "sub"
	userName := "Um nome qualquer"
	email := "email"
	name := "name"

	customer := &CognitoCreateUser{Username: userName, Name: name, Email: email}

	output := &cognito.AdminCreateUserOutput{
		User: &cognito.UserType{Username: &userName,
			Attributes: []*cognito.AttributeType{
				{Name: &email, Value: &email},
				{Name: &sub, Value: &sub},
				{Name: &name, Value: &name},
			},
		},
	}

	client.On("AdminCreateUser", mock.Anything).Return(output, nil).Once()

	response, err := cognitoClient.CreateUser(customer)

	assert.Nil(t, err)
	assert.Equal(t, response.Id, sub)

}
