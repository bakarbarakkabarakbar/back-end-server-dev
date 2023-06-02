package repositories

import (
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

func (ar AdminRepo) CreateCustomer(customer *entities.User) (*entities.User, error) {
	err := ar.db.Model(&entities.User{}).Create(customer).Error
	return customer, err
}
