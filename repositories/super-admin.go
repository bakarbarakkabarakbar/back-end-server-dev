package repositories

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/entities"
	"gorm.io/gorm"
)

type SuperAdminRepo struct {
	db *gorm.DB
}

func NewSuperAdminRepo(dbCrud *gorm.DB) SuperAdminRepo {
	return SuperAdminRepo{
		db: dbCrud,
	}
}

type SuperAdminRepoInterface interface {
	GetVerifiedAdmins() ([]entities.Actor, error)
	GetActiveAdmins() ([]entities.Actor, error)
	RemoveAdminById(id *uint) error
}

func (sar SuperAdminRepo) GetVerifiedAdmins() ([]entities.Actor, error) {
	var result = make([]entities.Actor, 0)
	var err = sar.db.Where("is_verified = ?", "true").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (sar SuperAdminRepo) GetActiveAdmins() ([]entities.Actor, error) {
	var result = make([]entities.Actor, 0)
	var err = sar.db.Where("is_active = ?", "true").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (sar SuperAdminRepo) RemoveAdminById(id *uint) error {
	var admin *entities.Actor
	var err error
	err = sar.db.Delete(&admin, id).Error

	if err != nil {
		return err
	}
	return nil
}
