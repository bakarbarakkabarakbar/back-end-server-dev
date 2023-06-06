package auth

import (
	"back-end-server-dev/entities"
	"back-end-server-dev/repositories"
	"time"
)

type UseCase struct {
	authRepo repositories.AuthRepo
}

type UseCaseInterface interface {
	GetCredentialByUsername(account *CredentialParam) (CredentialParam, error)
	CreateActorSession(customer *ActorSession) error
	GetLastActorSessionByToken(account *ActorSession) (ActorSession, error)
}

func (uc UseCase) GetCredentialByUsername(account *CredentialParam) (CredentialParam, error) {
	var result, err = uc.authRepo.GetActorByUsername(&account.username)
	return CredentialParam{
		id:       result.Id,
		username: result.Username,
		password: result.Password,
		roleId:   result.RoleId,
	}, err
}

func (uc UseCase) GetLastActorSessionByToken(account *ActorSession) (ActorSession, error) {
	var result, err = uc.authRepo.GetLastActorSessionByToken(&account.Token)
	return ActorSession{
		Id:      result.Id,
		ActorId: result.UserId,
		Token:   result.Token,
	}, err
}

func (uc UseCase) CreateActorSession(customer *ActorSession) error {
	var newSession *entities.ActorSession

	newSession = &entities.ActorSession{
		Id:        customer.Id,
		UserId:    customer.ActorId,
		Token:     customer.Token,
		CreatedAt: time.Now(),
		ExpiresAt: time.Now().Add(time.Hour * 1),
	}

	var err = uc.authRepo.CreateActorSession(newSession)
	return err
}
