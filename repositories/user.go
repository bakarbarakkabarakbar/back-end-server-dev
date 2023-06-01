package repositories

import (
	"fmt"
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

func (repo UserRepo) GetUsers(user *entity.User) ([]entity.User, error) {
	var users = make([]entity.User, 0)
	var err = repo.db.Find(&users).Error
	if err != nil {
		fmt.Println("error GetUsers")
		return []entity.User{}, err
	}
	return users, nil
}

func (repo UserRepo) GetUserById(id uint) (entity.User, error) {
	var user entity.User
	var err = repo.db.First(&user, "id", id).Error
	if err != nil {
		fmt.Println("error GetUsers", err)
		return user, err
	}
	return user, nil
}
