package customers

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/entities"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/repositories"
	"time"
)

type UseCase struct {
	userRepo repositories.CustomerRepo
}

type UseCaseInterface interface {
	CreateUser(user UserParam) (entities.Customer, error)
	GetUserById(id uint) (entities.Customer, error)
}

func (uc UseCase) GetUserById(id uint) (entities.Customer, error) {
	var user, err = uc.userRepo.GetCustomerById(id)
	return user, err
}

func (uc UseCase) CreateUser(user UserParam) (entities.Customer, error) {
	var newUser *entities.Customer

	newUser = &entities.Customer{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	_, err := uc.userRepo.CreateCustomer(newUser)
	if err != nil {
		return *newUser, err
	}
	return *newUser, nil
}
