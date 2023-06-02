package user

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/entities"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/repositories"
)

type UseCase struct {
	userRepo repositories.UserRepoInterface
}

type UsecaseInterface interface {
	GetUserByID(payload Payload) []entities.User
}

func (uc UseCase) GetUserById(id uint) (entities.User, error) {
	var user, err = uc.userRepo.GetUserById(id)
	return user, err
}
