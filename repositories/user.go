package repositories

import (
	"fmt"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/entities"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(dbCrud *gorm.DB) UserRepo {
	return UserRepo{
		db: dbCrud,
	}
}

type UserRepoInterface interface {
	GetByID(id int) []entities.User
	GetUsers(user *entities.User) ([]entities.User, error)
	GetUserById(id uint) (entities.User, error)
	CreateUser(user *entities.User) (*entities.User, error)
}

func (repo UserRepo) GetByID(id int) []entities.User {
	// implementasi query get user by id
	return []entities.User{}
}

func (repo UserRepo) CreateUser(user *entities.User) (*entities.User, error) {
	err := repo.db.Model(&entities.User{}).Create(user).Error
	return user, err
}

func (repo UserRepo) GetUsers(user *entities.User) ([]entities.User, error) {
	var users = make([]entities.User, 0)
	var err = repo.db.Find(&users).Error
	if err != nil {
		fmt.Println("error GetUsers")
		return []entities.User{}, err
	}
	return users, nil
}

func (repo UserRepo) GetUserById(id uint) (entities.User, error) {
	var user entities.User
	var err = repo.db.First(&user, "id", id).Error
	if err != nil {
		fmt.Println("error GetUsers", err)
		return user, err
	}
	return user, nil
}
