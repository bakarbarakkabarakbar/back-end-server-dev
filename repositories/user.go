package repositories

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/entities"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepo(dbCrud *gorm.DB) UserRepo {
	return UserRepo{
		db: dbCrud,
	}
}

type UserRepositoryInterface interface {
	GetByID(id int) []entities.User
}

func (repo UserRepository) GetByID(id int) []entities.User {
	// implementasi query get user by id
	return []entities.User{}
}

func (repo UserRepo) CreateUser(user *entity.User) (*entity.User, error) {
	err := repo.db.Model(&entity.User{}).Create(user).Error
	return user, err
}
