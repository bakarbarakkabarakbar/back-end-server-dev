package admin

import (
	"back-end-server-dev/entities"
	"back-end-server-dev/repositories"
	"crypto/sha1"
	"errors"
	"fmt"
)

type UseCase struct {
	adminRepo    repositories.AdminRepoInterface
	customerRepo repositories.CustomerRepoInterface
}

func NewUseCase(ar repositories.AdminRepo,
	cr repositories.CustomerRepo) UseCase {
	return UseCase{
		adminRepo:    ar,
		customerRepo: cr,
	}
}

type UseCaseInterface interface {
	GetCustomerById(customer *CustomerParam) (CustomerParam, error)
	GetCustomersByName(customer *CustomerParam) ([]CustomerParam, error)
	GetCustomersByEmail(customer *CustomerParam) ([]CustomerParam, error)
	GetAllCustomers(page *uint) ([]CustomerParam, error)
	CreateCustomer(customer *CustomerParam) error
	ModifyCustomer(customer *CustomerParam) error
	RemoveCustomerById(customer *CustomerParam) (CustomerParam, error)

	GetAdminById(admin *ActorParam) (ActorParam, error)
	GetAdminsByUsername(admin *ActorParam) ([]ActorParam, error)
	GetAllAdmins(page *uint) ([]ActorParam, error)
	CreateAdmin(admin *ActorParamWithPassword) error
	CreateRegisterAdmin(admin *RegisterApprovalParam) error
	ModifyAdmin(admin *ActorParamWithPassword) error
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
	if len(results) == 0 {
		return nil, errors.New("no match found")
	}
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
	if len(results) == 0 {
		return nil, errors.New("no match found")
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

func (uc UseCase) GetAllCustomers(page *uint) ([]CustomerParam, error) {
	var customers = make([]CustomerParam, 0)
	var results, err = uc.adminRepo.GetAllCustomers(page)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("no entry found")
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
		Id:        customer.Id,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
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
		Id:        result.Id,
		FirstName: customer.FirstName,
		LastName:  customer.LastName,
		Email:     customer.Email,
		Avatar:    customer.Avatar,
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

func (uc UseCase) GetAdminsByUsername(admin *ActorParam) ([]ActorParam, error) {
	var actors = make([]ActorParam, 0)
	var results, err = uc.adminRepo.GetAdminsByUsername(&admin.Username)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("no entry found")
	}

	for _, result := range results {
		actors = append(actors, ActorParam{
			Id:         result.Id,
			Username:   result.Username,
			RoleId:     result.RoleId,
			IsVerified: result.IsVerified,
			IsActive:   result.IsActive,
		})
	}
	return actors, nil
}

func (uc UseCase) GetAllAdmins(page *uint) ([]ActorParam, error) {
	var actors = make([]ActorParam, 0)
	var results, err = uc.adminRepo.GetAllAdmins(page)
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("no entry found")
	}

	for _, result := range results {
		actors = append(actors, ActorParam{
			Id:         result.Id,
			Username:   result.Username,
			RoleId:     result.RoleId,
			IsVerified: result.IsVerified,
			IsActive:   result.IsActive,
		})
	}
	return actors, nil
}

func (uc UseCase) CreateAdmin(admin *ActorParamWithPassword) error {
	var newAdmin *entities.Actor
	var hash = sha1.New()
	hash.Write([]byte(admin.Password))

	newAdmin = &entities.Actor{
		Id:         admin.Id,
		Username:   admin.Username,
		Password:   fmt.Sprintf("%x", hash.Sum(nil)),
		RoleId:     admin.RoleId,
		IsVerified: "false",
		IsActive:   "false",
	}

	var err = uc.adminRepo.CreateAdmin(newAdmin)
	return err
}

func (uc UseCase) CreateRegisterAdmin(admin *RegisterApprovalParam) error {
	var newRegistry *entities.RegisterApproval

	newRegistry = &entities.RegisterApproval{
		Id:           admin.Id,
		AdminId:      admin.AdminId,
		SuperAdminId: admin.SuperAdminId,
		Status:       "pending",
	}

	var err = uc.adminRepo.CreateRegisterAdmin(newRegistry)
	return err
}

func (uc UseCase) ModifyAdmin(admin *ActorParamWithPassword) error {
	var newAdmin *entities.Actor
	var result, err = uc.adminRepo.GetAdminById(&admin.Id)
	if err != nil {
		return err
	}
	var hash = sha1.New()
	hash.Write([]byte(admin.Password))

	newAdmin = &entities.Actor{
		Id:         admin.Id,
		Username:   admin.Username,
		Password:   fmt.Sprintf("%x", hash.Sum(nil)),
		RoleId:     admin.RoleId,
		IsVerified: result.IsVerified,
		IsActive:   result.IsActive,
	}

	err = uc.adminRepo.ModifyAdmin(newAdmin)
	return err
}
