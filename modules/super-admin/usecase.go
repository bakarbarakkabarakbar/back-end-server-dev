package super_admin

import (
	"errors"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/entities"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/repositories"
	"time"
)

type UseCase struct {
	superAdminRepo repositories.SuperAdminRepoInterface
	adminRepo      repositories.AdminRepoInterface
}

type UseCaseInterface interface {
	GetVerifiedAdmins() ([]ActorStatusParam, error)
	GetActiveAdmins() ([]ActorStatusParam, error)
	ModifyAdminStatusById(actor *ActorStatusParam) error
	RemoveAdminById(admin *ActorParam) (ActorParam, error)
}

func (uc UseCase) GetVerifiedAdmins() ([]ActorStatusParam, error) {
	var actors = make([]ActorStatusParam, 0)
	var results, err = uc.superAdminRepo.GetVerifiedAdmins()
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("no match found")
	}
	for _, result := range results {
		actors = append(actors, ActorStatusParam{
			Id:         result.Id,
			Username:   result.Username,
			RoleId:     result.RoleId,
			IsVerified: result.IsVerified,
			IsActive:   result.IsActive,
		})
	}
	return actors, nil
}

func (uc UseCase) GetActiveAdmins() ([]ActorStatusParam, error) {
	var actors = make([]ActorStatusParam, 0)
	var results, err = uc.superAdminRepo.GetActiveAdmins()
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("no match found")
	}
	for _, result := range results {
		actors = append(actors, ActorStatusParam{
			Id:         result.Id,
			Username:   result.Username,
			RoleId:     result.RoleId,
			IsVerified: result.IsVerified,
			IsActive:   result.IsActive,
		})
	}
	return actors, nil
}

func (uc UseCase) ModifyAdminStatusById(actor *ActorStatusParam) error {
	var newAdmin *entities.Actor
	var result, err = uc.adminRepo.GetAdminById(&actor.Id)
	if err != nil {
		return err
	}

	newAdmin = &entities.Actor{
		Id:         actor.Id,
		Username:   result.Username,
		Password:   result.Password,
		RoleId:     result.RoleId,
		IsVerified: actor.IsVerified,
		IsActive:   actor.IsActive,
		CreatedAt:  result.CreatedAt,
		ModifiedAt: time.Now(),
	}
	err = uc.adminRepo.ModifyAdmin(newAdmin)
	return err
}

func (uc UseCase) RemoveAdminById(admin *ActorParam) (ActorParam, error) {
	var result, err = uc.adminRepo.GetAdminById(&admin.Id)
	if err != nil {
		return ActorParam{}, err
	}

	err = uc.superAdminRepo.RemoveAdminById(&admin.Id)
	if err != nil {
		return ActorParam{}, err
	}
	var deletedCustomer = ActorParam{
		Id:         admin.Id,
		Username:   result.Username,
		RoleId:     result.RoleId,
		IsVerified: result.IsVerified,
		IsActive:   result.IsActive,
	}

	return deletedCustomer, err
}
