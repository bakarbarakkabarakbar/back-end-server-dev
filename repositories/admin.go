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
	CreateCustomer(user *entities.User) (*entities.User, error)
	RemoveCustomerById(id uint) (entities.User, error)
	GetCustomers(user *entities.User) ([]entities.User, error)
	GetCustomerById(id uint) (entities.User, error)
}

func (ar AdminRepo) CreateCustomer(customer *entities.User) (*entities.User, error) {
	err := ar.db.Model(&entities.User{}).Create(customer).Error
	return customer, err
}

func (ar AdminRepo) RemoveCustomerById(id uint) (entities.User, error) {
	var customer = entities.User{Id: id}
	var err = ar.db.Delete(&customer).Error
	if err != nil {
		fmt.Println("Error RemoveCustomerById", err)
		return customer, err
	}
	return customer, nil
}

func (ar AdminRepo) GetCustomers(user *entities.User) ([]entities.User, error) {
	var users = make([]entities.User, 0)
	var err = ar.db.Find(&users).Error
	if err != nil {
		fmt.Println("error GetCustomers")
		return []entities.User{}, err
	}
	return users, nil
}

func (ar AdminRepo) GetCustomerById(id uint) (entities.User, error) {
	var user entities.User
	var err = ar.db.First(&user, "id", id).Error
	if err != nil {
		fmt.Println("error GetCustomers", err)
		return user, err
	}
	return user, nil
}
