package repositories

import (
	"fmt"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/entities"
	"gorm.io/gorm"
)

type AdminRepo struct {
	db *gorm.DB
}

func NewAdminRepo(dbCrud *gorm.DB) AdminRepo {
	return AdminRepo{
		db: dbCrud,
	}
}

type AdminRepoInterface interface {
	GetCustomersByName(name *string) ([]entities.Customer, error)
	GetCustomersByEmail(email *string) ([]entities.Customer, error)
	CreateCustomer(customer *entities.Customer) error
	RemoveCustomerById(id *uint) (*entities.Customer, error)
}

func (ar AdminRepo) GetCustomersByName(name *string) ([]entities.Customer, error) {
	var customers = make([]entities.Customer, 0)
	var err = ar.db.Raw(
		"SELECT * FROM customers WHERE CONCAT(first_name, ' ', last_name) LIKE %", name, "%").Scan(&customers).Error
	//var err = ar.db.Where("Name LIKE ?", "%"+*name+"%").Find(&customers).Error
	if err != nil {
		fmt.Println("error GetCustomersByName", err)
		return nil, err
	}
	return customers, nil
}

func (ar AdminRepo) GetCustomersByEmail(email *string) ([]entities.Customer, error) {
	var customers = make([]entities.Customer, 0)
	var err = ar.db.Where(&entities.Customer{Email: *email}).Find(&customers).Error
	if err != nil {
		fmt.Println("error GetCustomersByEmail", err)
		return nil, err
	}
	return customers, nil
}

func (ar AdminRepo) CreateCustomer(customer *entities.Customer) error {
	fmt.Println(customer)
	err := ar.db.Model(&entities.Customer{}).Create(customer).Error
	return err
}

func (ar AdminRepo) RemoveCustomerById(id *uint) (*entities.Customer, error) {
	var customer *entities.Customer

	if err := ar.db.First(customer, "id", *id).Error; err != nil {
		fmt.Println("Error RemoveCustomerById: Customer Not Found", err)
		return customer, err
	}

	if err := ar.db.Delete(customer).Error; err != nil {
		fmt.Println("Error RemoveCustomerById", err)
		return customer, err
	}

	return customer, nil
}
