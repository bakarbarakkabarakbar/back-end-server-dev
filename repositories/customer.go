package repositories

import (
	"fmt"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/entities"
	"gorm.io/gorm"
)

type CustomerRepo struct {
	db *gorm.DB
}

func NewCustomerRepo(dbCrud *gorm.DB) CustomerRepo {
	return CustomerRepo{
		db: dbCrud,
	}
}

type CustomerRepoInterface interface {
	GetByID(id int) []entities.User
	GetCustomers(user *entities.User) ([]entities.User, error)
	GetCustomerById(id uint) (entities.User, error)
	CreateCustomer(user *entities.User) (*entities.User, error)
}

func (repo CustomerRepo) GetByID(id int) []entities.User {
	// implementasi query get customers by id
	return []entities.User{}
}

func (repo CustomerRepo) CreateCustomer(user *entities.User) (*entities.User, error) {
	err := repo.db.Model(&entities.User{}).Create(user).Error
	return user, err
}

func (repo CustomerRepo) GetCustomers(user *entities.User) ([]entities.User, error) {
	var users = make([]entities.User, 0)
	var err = repo.db.Find(&users).Error
	if err != nil {
		fmt.Println("error GetCustomers")
		return []entities.User{}, err
	}
	return users, nil
}

func (repo CustomerRepo) GetCustomerById(id uint) (entities.User, error) {
	var user entities.User
	var err = repo.db.First(&user, "id", id).Error
	if err != nil {
		fmt.Println("error GetCustomers", err)
		return user, err
	}
	return user, nil
}
