package super_admin

import (
	"back-end-server-dev/entities"
	"back-end-server-dev/repositories"
	"errors"
	"time"
)

type UseCase struct {
	superAdminRepo repositories.SuperAdminRepoInterface
	adminRepo      repositories.AdminRepoInterface
}

type UseCaseInterface interface {
	GetVerifiedAdmins() ([]ActorParam, error)
	GetActiveAdmins() ([]ActorParam, error)
	GetRegisterAdminById(register *RegisterApprovalParam) (RegisterApprovalParam, error)
	GetApprovedAdmins() ([]RegisterApprovalParam, error)
	GetRejectedAdmin() ([]RegisterApprovalParam, error)
	GetPendingAdmins() ([]RegisterApprovalParam, error)
	ModifyStatusAdminById(actor *ActorParam) error
	ModifyRegisterAdminById(register *RegisterApprovalParam) error
	RemoveAdminById(admin *ActorParam) (ActorParam, error)
	RemoveRegisterAdminById(register *RegisterApprovalParam) (RegisterApprovalParam, error)
}

func (uc UseCase) GetVerifiedAdmins() ([]ActorParam, error) {
	var actors = make([]ActorParam, 0)
	var results, err = uc.superAdminRepo.GetVerifiedAdmins()
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("no match found")
	}
	for _, result := range results {
		actors = append(actors, ActorParam{
			Id:         result.Id,
			Username:   result.Username,
			RoleId:     result.RoleId,
			IsVerified: result.IsVerified,
			IsActive:   result.IsActive,
		})
	}
	return actors, nil
}

func (uc UseCase) GetActiveAdmins() ([]ActorParam, error) {
	var actors = make([]ActorParam, 0)
	var results, err = uc.superAdminRepo.GetActiveAdmins()
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("no match found")
	}
	for _, result := range results {
		actors = append(actors, ActorParam{
			Id:         result.Id,
			Username:   result.Username,
			RoleId:     result.RoleId,
			IsVerified: result.IsVerified,
			IsActive:   result.IsActive,
		})
	}
	return actors, nil
}

func (uc UseCase) GetRegisterAdminById(register *RegisterApprovalParam) (RegisterApprovalParam, error) {
	var newRegister RegisterApprovalParam
	var result, err = uc.superAdminRepo.GetRegisterAdminById(&register.Id)
	if err != nil {
		return RegisterApprovalParam{}, err
	}

	newRegister = RegisterApprovalParam{
		Id:           result.Id,
		AdminId:      result.AdminId,
		SuperAdminId: result.SuperAdminId,
		Status:       result.Status,
	}
	return newRegister, err
}

func (uc UseCase) GetApprovedAdmins() ([]RegisterApprovalParam, error) {
	var registers = make([]RegisterApprovalParam, 0)
	var results, err = uc.superAdminRepo.GetApprovedAdmins()
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("no match found")
	}
	for _, result := range results {
		registers = append(registers, RegisterApprovalParam{
			Id:           result.Id,
			AdminId:      result.AdminId,
			SuperAdminId: result.SuperAdminId,
			Status:       result.Status,
		})
	}
	return registers, nil
}

func (uc UseCase) GetRejectedAdmin() ([]RegisterApprovalParam, error) {
	var registers = make([]RegisterApprovalParam, 0)
	var results, err = uc.superAdminRepo.GetRejectedAdmins()
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("no match found")
	}
	for _, result := range results {
		registers = append(registers, RegisterApprovalParam{
			Id:           result.Id,
			AdminId:      result.AdminId,
			SuperAdminId: result.SuperAdminId,
			Status:       result.Status,
		})
	}
	return registers, nil
}

func (uc UseCase) GetPendingAdmins() ([]RegisterApprovalParam, error) {
	var registers = make([]RegisterApprovalParam, 0)
	var results, err = uc.superAdminRepo.GetPendingAdmins()
	if err != nil {
		return nil, err
	}
	if len(results) == 0 {
		return nil, errors.New("no match found")
	}
	for _, result := range results {
		registers = append(registers, RegisterApprovalParam{
			Id:           result.Id,
			AdminId:      result.AdminId,
			SuperAdminId: result.SuperAdminId,
			Status:       result.Status,
		})
	}
	return registers, nil
}

func (uc UseCase) ModifyStatusAdminById(actor *ActorParam) error {
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

func (uc UseCase) ModifyRegisterAdminById(register *RegisterApprovalParam) error {
	var newAdmin *entities.RegisterApproval
	var result, err = uc.superAdminRepo.GetRegisterAdminById(&register.Id)
	if err != nil {
		return err
	}

	newAdmin = &entities.RegisterApproval{
		Id:           result.Id,
		AdminId:      result.AdminId,
		SuperAdminId: result.SuperAdminId,
		Status:       register.Status,
	}
	err = uc.superAdminRepo.ModifyAdminApproval(newAdmin)
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

func (uc UseCase) RemoveRegisterAdminById(register *RegisterApprovalParam) (RegisterApprovalParam, error) {
	var result, err = uc.superAdminRepo.GetRegisterAdminById(&register.Id)
	if err != nil {
		return RegisterApprovalParam{}, err
	}

	err = uc.superAdminRepo.RemoveRegisterAdminById(&register.Id)
	if err != nil {
		return RegisterApprovalParam{}, err
	}
	var deletedRegister = RegisterApprovalParam{
		Id:           result.Id,
		AdminId:      result.AdminId,
		SuperAdminId: result.SuperAdminId,
		Status:       result.Status,
	}

	return deletedRegister, err
}
