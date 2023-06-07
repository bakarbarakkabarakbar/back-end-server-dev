package repositories

import (
	"back-end-server-dev/entities"
	"gorm.io/gorm"
)

type SuperAdminRepo struct {
	db *gorm.DB
}

func NewSuperAdminRepo(dbCrud *gorm.DB) SuperAdminRepo {
	return SuperAdminRepo{
		db: dbCrud,
	}
}

type SuperAdminRepoInterface interface {
	GetVerifiedAdmins() ([]entities.Actor, error)
	GetActiveAdmins() ([]entities.Actor, error)
	GetRegisterAdminById(id *uint) (entities.RegisterApproval, error)
	GetApprovedAdmins() ([]entities.RegisterApproval, error)
	GetRejectedAdmins() ([]entities.RegisterApproval, error)
	GetPendingAdmins() ([]entities.RegisterApproval, error)
	GetRegisterRequestAdmins() ([]entities.RegisterApproval, error)
	ModifyRegisterAdminById(adminRegister *entities.RegisterApproval) error
	RemoveAdminById(id *uint) error
	RemoveRegisterAdminById(id *uint) error
}

func (sar SuperAdminRepo) GetVerifiedAdmins() ([]entities.Actor, error) {
	var result = make([]entities.Actor, 0)
	var err = sar.db.Where("is_verified = ?", "true").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (sar SuperAdminRepo) GetActiveAdmins() ([]entities.Actor, error) {
	var result = make([]entities.Actor, 0)
	var err = sar.db.Where("is_active = ?", "true").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (sar SuperAdminRepo) GetRegisterAdminById(id *uint) (entities.RegisterApproval, error) {
	var register entities.RegisterApproval
	var err = sar.db.First(&register, id).Error
	if err != nil {
		return register, err
	}
	return register, nil
}

func (sar SuperAdminRepo) GetApprovedAdmins() ([]entities.RegisterApproval, error) {
	var result = make([]entities.RegisterApproval, 0)
	var err = sar.db.Where("status = ?", "approved").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (sar SuperAdminRepo) GetRejectedAdmins() ([]entities.RegisterApproval, error) {
	var result = make([]entities.RegisterApproval, 0)
	var err = sar.db.Where("status = ?", "rejected").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (sar SuperAdminRepo) GetPendingAdmins() ([]entities.RegisterApproval, error) {
	var result = make([]entities.RegisterApproval, 0)
	var err = sar.db.Where("status = ?", "pending").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (sar SuperAdminRepo) GetRegisterRequestAdmins() ([]entities.RegisterApproval, error) {
	var result = make([]entities.RegisterApproval, 0)
	var err = sar.db.Where("status = ?", "pending").Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (sar SuperAdminRepo) ModifyRegisterAdminById(adminRegister *entities.RegisterApproval) error {
	err := sar.db.Save(&adminRegister).Error
	return err
}

func (sar SuperAdminRepo) RemoveAdminById(id *uint) error {
	var admin *entities.Actor
	var err error
	err = sar.db.Delete(&admin, id).Error

	if err != nil {
		return err
	}
	return nil
}

func (sar SuperAdminRepo) RemoveRegisterAdminById(id *uint) error {
	var register *entities.RegisterApproval
	var err error
	err = sar.db.Delete(&register, id).Error

	if err != nil {
		return err
	}
	return nil
}
