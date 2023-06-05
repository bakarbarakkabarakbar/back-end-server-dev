package auth

import "github.com/dibimbing-satkom-indo/onion-architecture-go/repositories"

type UseCase struct {
	authRepo repositories.AuthRepo
}

type UseCaseInterface interface {
	GetCredentialByUsername(account *CredentialParam) (CredentialParam, error)
}

func (uc UseCase) GetCredentialByUsername(account *CredentialParam) (CredentialParam, error) {
	var result, err = uc.authRepo.GetActorByUsername(&account.username)
	return CredentialParam{
		username: result.Username,
		password: result.Password,
	}, err
}
