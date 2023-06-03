package admin

import (
	"crypto/sha1"
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
	CreateCustomer(customer *CustomerParam) error
	ModifyCustomer(customer *CustomerParam) error
	RemoveCustomerById(customer *CustomerParam) (CustomerParam, error)

	GetAdminById(admin *ActorParam) (ActorParam, error)
	CreateAdmin(admin *ActorParamWithPassword) error
	ModifyAdmin(admin *ActorParamWithPassword) error
	RemoveAdminById(admin *ActorParam) (ActorParam, error)
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
	var results, err = uc.adminRepo.GetCustomersByName(&customer.FirstName)
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

func (uc UseCase) CreateCustomer(customer *CustomerParam) error {
	var newCustomer *entities.Customer

	newCustomer = &entities.Customer{
		FirstName:  customer.FirstName,
		LastName:   customer.LastName,
		Email:      customer.Email,
		Avatar:     customer.Avatar,
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	var err = uc.adminRepo.CreateCustomer(newCustomer)
	return err
}

func (uc UseCase) ModifyCustomer(customer *CustomerParam) error {
	var newCustomer *entities.Customer
	var result, err = uc.customerRepo.GetCustomerById(&customer.Id)
	if err != nil {
		return err
	}
	newCustomer = &entities.Customer{
		Id:         result.Id,
		FirstName:  customer.FirstName,
		LastName:   customer.LastName,
		Email:      customer.Email,
		Avatar:     customer.Avatar,
		CreatedAt:  result.CreatedAt,
		ModifiedAt: time.Now(),
	}

	err = uc.adminRepo.ModifyCustomer(newCustomer)
	return err
}

func (uc UseCase) RemoveCustomerById(customer *CustomerParam) (CustomerParam, error) {
	var result, err = uc.customerRepo.GetCustomerById(&customer.Id)
	if err != nil {
		return CustomerParam{}, err
	}
	err = uc.adminRepo.RemoveCustomerById(&customer.Id)
	if err != nil {
		return CustomerParam{}, err
	}
	var deletedCustomer = CustomerParam{
		Id:        result.Id,
		FirstName: result.FirstName,
		LastName:  result.LastName,
		Email:     result.Email,
		Avatar:    result.Avatar,
	}

	return deletedCustomer, err
}

func (uc UseCase) GetAdminById(admin *ActorParam) (ActorParam, error) {
	var result, err = uc.adminRepo.GetAdminById(&admin.Id)
	if err != nil {
		return ActorParam{}, err
	}
	return ActorParam{
		Id:         admin.Id,
		Username:   result.Username,
		RoleId:     result.RoleId,
		IsVerified: result.IsVerified,
		IsActive:   result.IsActive,
	}, err
}

func (uc UseCase) CreateAdmin(admin *ActorParamWithPassword) error {
	var newAdmin *entities.Actor
	var hash = sha1.New()
	hash.Write([]byte(admin.Password))

	newAdmin = &entities.Actor{
		Username:   admin.Username,
		Password:   fmt.Sprintf("%x", hash.Sum(nil)),
		RoleId:     admin.RoleId,
		IsVerified: "false",
		IsActive:   "false",
		CreatedAt:  time.Now(),
		ModifiedAt: time.Now(),
	}

	var err = uc.adminRepo.CreateAdmin(newAdmin)
	return err
}

func (uc UseCase) ModifyAdmin(admin *ActorParamWithPassword) error {
	var newAdmin *entities.Actor
	var result, err = uc.adminRepo.GetAdminById(&admin.Id)
	if err != nil {
		return err
	}

	newAdmin = &entities.Actor{
		Id:         admin.Id,
		Username:   admin.Username,
		Password:   admin.Password,
		RoleId:     admin.RoleId,
		IsVerified: result.IsVerified,
		IsActive:   result.IsActive,
		CreatedAt:  result.CreatedAt,
		ModifiedAt: time.Now(),
	}

	err = uc.adminRepo.ModifyAdmin(newAdmin)
	return err
}

func (uc UseCase) RemoveAdminById(admin *ActorParam) (ActorParam, error) {
	var result, err = uc.adminRepo.GetAdminById(&admin.Id)
	if err != nil {
		return ActorParam{}, err
	}

	err = uc.adminRepo.RemoveAdminById(&admin.Id)
	if err != nil {
		return ActorParam{}, err
	}
	var deletedCustomer = ActorParam{
		Id:         admin.Id,
		Username:   result.Username,
		RoleId:     result.RoleId,
		IsVerified: result.IsVerified,
		IsActive:   result.IsActive,
	}

	return deletedCustomer, err
}
