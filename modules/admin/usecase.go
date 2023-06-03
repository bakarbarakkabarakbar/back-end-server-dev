package admin

import (
	"fmt"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/entities"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/repositories"
	"time"
)

type UseCase struct {
	adminRepo    repositories.AdminRepoInterface
	customerRepo repositories.CustomerRepoInterface
}

type UseCaseInterface interface {
	GetCustomerById(customer *CustomerParam) (CustomerParam, error)
	GetCustomersByName(customer *CustomerParam) ([]CustomerParam, error)
	GetCustomersByEmail(customer *CustomerParam) ([]CustomerParam, error)
	CreateCustomer(user *CustomerParam) error
	RemoveCustomerById(customer *CustomerParam) (CustomerParam, error)
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

func (uc UseCase) GetCustomersByName(customer *CustomerParam) ([]CustomerParam, error) {
	var customers = make([]CustomerParam, 0)
	var name = fmt.Sprintf("%s %s", customer.FirstName, customer.LastName)
	var results, err = uc.adminRepo.GetCustomersByName(&name)
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		customers = append(customers, CustomerParam{
			Id:        result.Id,
			FirstName: result.FirstName,
			LastName:  result.LastName,
			Email:     result.Email,
			Avatar:    result.Avatar,
		})
	}
	return customers, nil
}

func (uc UseCase) GetCustomersByEmail(customer *CustomerParam) ([]CustomerParam, error) {
	var customers = make([]CustomerParam, 0)
	var results, err = uc.adminRepo.GetCustomersByEmail(&customer.Email)
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		customers = append(customers, CustomerParam{
			Id:        result.Id,
			FirstName: result.FirstName,
			LastName:  result.LastName,
			Email:     result.Email,
			Avatar:    result.Avatar,
		})
	}
	return customers, nil
}

func (uc UseCase) CreateCustomer(user *CustomerParam) error {
	var newUser *entities.Customer

	newUser = &entities.Customer{
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		Email:      user.Email,
		Avatar:     user.Avatar,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	var err = uc.adminRepo.CreateCustomer(newUser)
	return err
}

func (uc UseCase) RemoveCustomerById(customer *CustomerParam) (CustomerParam, error) {
	var result, err = uc.adminRepo.RemoveCustomerById(&customer.Id)
	var deletedCustomer = CustomerParam{
		Id:        result.Id,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
		Avatar:    result.Avatar,
	}

	return deletedCustomer, err
}
