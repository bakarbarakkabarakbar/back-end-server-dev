package repositories

import (
	"back-end-server-dev/entities"
	"fmt"
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

//go:generate mockery --name AdminRepoInterface
type AdminRepoInterface interface {
	GetCustomersByName(name *string) ([]entities.Customer, error)
	GetCustomersByEmail(email *string) ([]entities.Customer, error)
	GetAllCustomers(page *uint) ([]entities.Customer, error)
	CreateCustomer(customer *entities.Customer) error
	ModifyCustomer(customer *entities.Customer) error
	RemoveCustomerById(id *uint) error

	GetAdminById(id *uint) (entities.Actor, error)
	GetAdminsByUsername(username *string) ([]entities.Actor, error)
	GetAllAdmins(page *uint) ([]entities.Actor, error)
	CreateAdmin(admin *entities.Actor) error
	CreateRegisterAdmin(register *entities.RegisterApproval) error
	ModifyAdmin(admin *entities.Actor) error
}

func (ar AdminRepo) GetCustomersByName(name *string) ([]entities.Customer, error) {
	var customers = make([]entities.Customer, 0)
	var err = ar.db.Raw(
		fmt.Sprint("SELECT * FROM customers WHERE CONCAT(first_name, ' ', last_name) LIKE \"%", *name, "%\"")).Scan(&customers).Error
	//var err = ar.connection.Where("Name LIKE ?", "%"+*name+"%").Find(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (ar AdminRepo) GetCustomersByEmail(email *string) ([]entities.Customer, error) {
	var customers = make([]entities.Customer, 0)
	var err = ar.db.Raw(
		fmt.Sprint("SELECT * FROM customers WHERE email LIKE \"%", *email, "%\"")).Scan(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (ar AdminRepo) GetAllCustomers(page *uint) ([]entities.Customer, error) {
	var customers = make([]entities.Customer, 0)
	var err = ar.db.Raw(
		fmt.Sprint("SELECT * FROM customers LIMIT 6 OFFSET ", (*page-1)*6)).Scan(&customers).Error
	if err != nil {
		return nil, err
	}
	return customers, nil
}

func (ar AdminRepo) CreateCustomer(customer *entities.Customer) error {
	err := ar.db.Model(&entities.Customer{}).Create(customer).Error
	return err
}

func (ar AdminRepo) ModifyCustomer(customer *entities.Customer) error {
	err := ar.db.Save(&customer).Error
	return err
}

func (ar AdminRepo) RemoveCustomerById(id *uint) error {
	var customer *entities.Customer
	var err error
	err = ar.db.Delete(&customer, id).Error

	if err != nil {
		return err
	}
	return nil
}

func (ar AdminRepo) GetAdminById(id *uint) (entities.Actor, error) {
	var admin entities.Actor
	var err = ar.db.First(&admin, id).Error
	if err != nil {
		return admin, err
	}
	return admin, nil
}

func (ar AdminRepo) GetAdminsByUsername(username *string) ([]entities.Actor, error) {
	var actors = make([]entities.Actor, 0)
	var err = ar.db.Raw(
		fmt.Sprint("SELECT * FROM actors WHERE username LIKE \"%", *username, "%\"")).Scan(&actors).Error
	if err != nil {
		return nil, err
	}
	return actors, nil
}

func (ar AdminRepo) GetAllAdmins(page *uint) ([]entities.Actor, error) {
	var actors = make([]entities.Actor, 0)
	var err = ar.db.Raw(
		fmt.Sprint("SELECT * FROM actors LIMIT 6 OFFSET ", (*page-1)*6)).Scan(&actors).Error
	if err != nil {
		return nil, err
	}
	return actors, nil
}

func (ar AdminRepo) CreateAdmin(admin *entities.Actor) error {
	err := ar.db.Model(&entities.Actor{}).Create(admin).Error
	return err
}

func (ar AdminRepo) CreateRegisterAdmin(register *entities.RegisterApproval) error {
	err := ar.db.Model(&entities.RegisterApproval{}).Create(register).Error
	return err
}

func (ar AdminRepo) ModifyAdmin(admin *entities.Actor) error {
	err := ar.db.Save(&admin).Error
	return err
}
