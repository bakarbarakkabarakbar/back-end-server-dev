package repositories

import (
	"fmt"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/entities"
	"gorm.io/gorm"
)

type AdminRepo struct {
	db *gorm.DB
}

func NewAccountRepo(dbCrud *gorm.DB) AdminRepo {
	return AdminRepo{
		db: dbCrud,
	}
}

type AdminRepoInterface interface {
	GetByID(id int) []entities.Customer
	GetCustomers(user *entities.Customer) ([]entities.Customer, error)
	GetCustomerById(id uint) (entities.Customer, error)
	CreateCustomer(user *entities.Customer) (*entities.Customer, error)
}

func (ar AdminRepo) CreateCustomer(customer *entities.Customer) (*entities.Customer, error) {
	err := ar.db.Model(&entities.Customer{}).Create(customer).Error
	return customer, err
}

func (ar AdminRepo) RemoveCustomerById(id uint) (entities.Customer, error) {
	var customer = entities.Customer{Id: id}
	var err = ar.db.Delete(&customer).Error
	if err != nil {
		fmt.Println("Error RemoveCustomerById", err)
		return customer, err
	}
	return customer, nil
}

func (ar AdminRepo) GetCustomers(user *entities.Customer) ([]entities.Customer, error) {
	var users = make([]entities.Customer, 0)
	var err = ar.db.Find(&users).Error
	if err != nil {
		fmt.Println("error GetCustomers")
		return []entities.Customer{}, err
	}
	return users, nil
}

func (ar AdminRepo) GetCustomerById(id uint) (entities.Customer, error) {
	var user entities.Customer
	var err = ar.db.First(&user, "id", id).Error
	if err != nil {
		fmt.Println("error GetCustomers", err)
		return user, err
	}
	return user, nil
}
