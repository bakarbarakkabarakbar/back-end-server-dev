package customers

import (
	"back-end-server-dev/repositories"
)

type UseCase struct {
	cr repositories.CustomerRepoInterface
}

func NewUseCase(cr repositories.CustomerRepo) UseCase {
	return UseCase{cr: cr}
}

type UseCaseInterface interface {
	GetCustomerById(customer *CustomerParam) (CustomerParam, error)
	GetCustomerByEmail(customer *CustomerParam) (CustomerParam, error)
}

func (uc UseCase) GetCustomerById(customer *CustomerParam) (CustomerParam, error) {
	var result, err = uc.cr.GetCustomerById(&customer.Id)
	return CustomerParam{
		Id:        result.Id,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
		Avatar:    result.Avatar,
	}, err
}

func (uc UseCase) GetCustomerByEmail(customer *CustomerParam) (CustomerParam, error) {
	var result, err = uc.cr.GetCustomerByEmail(&customer.Email)
	return CustomerParam{
		Id:        result.Id,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
		Avatar:    result.Avatar,
	}, err
}
