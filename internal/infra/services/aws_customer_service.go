package services

import (
	"tech-challenge-fase-1/internal/core/entities"
	valueobjects "tech-challenge-fase-1/internal/core/value_objects"
)

type AwsCustomerService struct {
	serviceClient *CognitoClient
}

func NewAwsCustomerService(
	client CognitoClientInterface,
	userPoolId string,
) *AwsCustomerService {
	cognitoClient := NewCognitoClient(client, userPoolId)

	return &AwsCustomerService{
		serviceClient: cognitoClient,
	}
}

func (a *AwsCustomerService) GetCustomerByCPF(cpf *valueobjects.CPF) (*entities.Customer, error) {
	user, err := a.serviceClient.GetUser(cpf.Value())
	if err != nil {
		return nil, err
	}
	return entities.RestoreCustomer(user.Id, user.Name, user.Email, user.Username)
}

func (a *AwsCustomerService) GetCustomerById(id string) (*entities.Customer, error) {
	user, err := a.serviceClient.GetUserBySub(id)
	if err != nil {
		return nil, err
	}
	return entities.RestoreCustomer(user.Id, user.Name, user.Email, user.Username)
}
