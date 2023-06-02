package repositories

import (
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
