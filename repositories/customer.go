package repositories

import (
	"fmt"
	"gorm.io/gorm"
	"user-management-backend/entities"
)

type CustomerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(gorm *gorm.DB) CustomerRepo {
	return CustomerRepo{
		db: gorm,
	}
}

//go:generate mockery --name CustomerRepoInterface
type CustomerRepoInterface interface {
	GetCustomerById(id *uint) (entities.Customer, error)
	GetCustomerByEmail(email *string) (entities.Customer, error)
}

func (cr CustomerRepo) GetCustomerById(id *uint) (entities.Customer, error) {
	fmt.Println(*id)
	var customer entities.Customer
	var err = cr.db.First(&customer, id).Error
	if err != nil {
		fmt.Println("error GetCustomersById", err)
		return customer, err
	}
	return customer, nil
}

func (cr CustomerRepo) GetCustomerByEmail(email *string) (entities.Customer, error) {
	var user entities.Customer
	var err = cr.db.First(&user, "Email", email).Error
	if err != nil {
		fmt.Println("error GetCustomerByEmail", err)
		return user, err
	}
	return user, nil
}
