package repositories

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/entities"
	"gorm.io/gorm"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(dbCrud *gorm.DB) AuthRepo {
	return AuthRepo{
		db: dbCrud,
	}
}

func (ar AuthRepo) GetActorByUsername(username *string) (entities.Actor, error) {
	var admin entities.Actor
	var err = ar.db.Where("username = ?", username).Find(&admin).Error
	if err != nil {
		return admin, err
	}
	return admin, nil
}
