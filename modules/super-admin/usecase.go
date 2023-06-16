package super_admin

import (
	"errors"
	"user-management-backend/entities"
	"user-management-backend/repositories"
)

type UseCase struct {
	sar repositories.SuperAdminRepoInterface
	ar  repositories.AdminRepoInterface
}

func NewUseCase(sar repositories.SuperAdminRepo, ar repositories.AdminRepo) UseCase {
	return UseCase{
		sar: sar,
		ar:  ar,
	}
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
	var results, err = uc.sar.GetVerifiedAdmins()
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
	var results, err = uc.sar.GetActiveAdmins()
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
	var result, err = uc.sar.GetRegisterAdminById(&register.Id)
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
	var results, err = uc.sar.GetApprovedAdmins()
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
	var results, err = uc.sar.GetRejectedAdmins()
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
	var results, err = uc.sar.GetPendingAdmins()
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
	var result, err = uc.ar.GetAdminById(&actor.Id)
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
	}
	err = uc.ar.ModifyAdmin(newAdmin)
	return err
}

func (uc UseCase) ModifyRegisterAdminById(register *RegisterApprovalParam) error {
	var newAdmin *entities.RegisterApproval
	var result, err = uc.sar.GetRegisterAdminById(&register.Id)
	if err != nil {
		return err
	}

	newAdmin = &entities.RegisterApproval{
		Id:           result.Id,
		AdminId:      result.AdminId,
		SuperAdminId: result.SuperAdminId,
		Status:       register.Status,
	}
	err = uc.sar.ModifyRegisterAdminById(newAdmin)
	return err
}

func (uc UseCase) RemoveAdminById(admin *ActorParam) (ActorParam, error) {
	var result, err = uc.ar.GetAdminById(&admin.Id)
	if err != nil {
		return ActorParam{}, err
	}

	err = uc.sar.RemoveAdminById(&admin.Id)
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
	var result, err = uc.sar.GetRegisterAdminById(&register.Id)
	if err != nil {
		return RegisterApprovalParam{}, err
	}

	err = uc.sar.RemoveRegisterAdminById(&register.Id)
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
