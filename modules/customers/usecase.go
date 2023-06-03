package customers

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/repositories"
)

type UseCase struct {
	customerRepo repositories.CustomerRepoInterface
}

type UseCaseInterface interface {
	GetCustomerById(customer *CustomerParam) (CustomerParam, error)
	GetCustomerByEmail(customer *CustomerParam) (CustomerParam, error)
}

func (uc UseCase) GetCustomerById(customer *CustomerParam) (CustomerParam, error) {
	var result, err = uc.customerRepo.GetCustomerById(&customer.Id)
	return CustomerParam{
		Id:        result.Id,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
		Avatar:    result.Avatar,
	}, err
}

func (uc UseCase) GetCustomerByEmail(customer *CustomerParam) (CustomerParam, error) {
	var result, err = uc.customerRepo.GetCustomerByEmail(&customer.Email)
	return CustomerParam{
		Id:        result.Id,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
		Avatar:    result.Avatar,
	}, err
}
