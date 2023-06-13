package auth

import (
	"back-end-server-dev/entities"
	"back-end-server-dev/repositories"
)

type UseCase struct {
	authRepo repositories.AuthRepoInterface
}

type UseCaseInterface interface {
	GetCredentialByUsername(account *CredentialParam) (CredentialParam, error)
	CreateActorSession(customer *ActorSessionParam) error
	GetLastActorSessionByToken(account *ActorSessionParam) (ActorSessionParam, error)
}

func (uc UseCase) GetCredentialByUsername(account *CredentialParam) (CredentialParam, error) {
	var result, err = uc.authRepo.GetActorByUsername(&account.username)
	if err != nil {
		return CredentialParam{}, err
	}
	return CredentialParam{
		id:       result.Id,
		username: result.Username,
		password: result.Password,
		roleId:   result.RoleId,
	}, err
}

func (uc UseCase) GetLastActorSessionByToken(account *ActorSessionParam) (ActorSessionParam, error) {
	var result, err = uc.authRepo.GetLastActorSessionByToken(&account.Token)
	return ActorSessionParam{
		Id:      result.Id,
		ActorId: result.UserId,
		Token:   result.Token,
	}, err
}

func (uc UseCase) CreateActorSession(customer *ActorSessionParam) error {
	var newSession *entities.ActorSession

	newSession = &entities.ActorSession{
		Id:     customer.Id,
		UserId: customer.ActorId,
		Token:  customer.Token,
	}

	var err = uc.authRepo.CreateActorSession(newSession)
	return err
}
