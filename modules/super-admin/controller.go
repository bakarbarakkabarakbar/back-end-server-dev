package super_admin

import "back-end-server-dev/dto"

type Controller struct {
	uc UseCaseInterface
}

func NewController(uc UseCase) Controller {
	return Controller{uc: uc}
}

type ControllerInterface interface {
	GetVerifiedAdmins() (ResponseParam, error)
	GetActiveAdmins() (ResponseParam, error)
	GetRegisterAdminById(req *RegisterApprovalParam) (ResponseParam, error)
	GetApprovedAdmins() (ResponseParam, error)
	GetRejectedAdmins() (ResponseParam, error)
	GetPendingAdmins() (ResponseParam, error)
	ModifyStatusAdminById(req *ActorParam) (ResponseParam, error)
	ModifyRegisterAdminById(req *RegisterApprovalParam) (ResponseParam, error)
	RemoveAdminById(req *ActorParam) (ResponseParam, error)
	RemoveRegisterAdminById(req *RegisterApprovalParam) (ResponseParam, error)
}

func (ctrl Controller) GetVerifiedAdmins() (ResponseParam, error) {
	var results []ActorParam
	var err error
	results, err = ctrl.uc.GetVerifiedAdmins()
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed GetVerifiedAdmins",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetVerifiedAdmins",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: results,
	}
	return res, nil
}

func (ctrl Controller) GetActiveAdmins() (ResponseParam, error) {
	var results []ActorParam
	var err error
	results, err = ctrl.uc.GetActiveAdmins()
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed GetActiveAdmins",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetActiveAdmins",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: results,
	}
	return res, nil
}

func (ctrl Controller) GetRegisterAdminById(req *RegisterApprovalParam) (ResponseParam, error) {
	var result RegisterApprovalParam
	var err error
	result, err = ctrl.uc.GetRegisterAdminById(req)
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed GetRegisterAdminById",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetRegisterAdminById",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: result,
	}
	return res, nil
}

func (ctrl Controller) GetApprovedAdmins() (ResponseParam, error) {
	var results []RegisterApprovalParam
	var err error
	results, err = ctrl.uc.GetApprovedAdmins()
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed GetApprovedAdmins",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetApprovedAdmins",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: results,
	}
	return res, nil
}

func (ctrl Controller) GetRejectedAdmins() (ResponseParam, error) {
	var results []RegisterApprovalParam
	var err error
	results, err = ctrl.uc.GetRejectedAdmin()
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed GetRejectedAdmins",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetRejectedAdmins",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: results,
	}
	return res, nil
}

func (ctrl Controller) GetPendingAdmins() (ResponseParam, error) {
	var results []RegisterApprovalParam
	var err error
	results, err = ctrl.uc.GetPendingAdmins()
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed GetPendingAdmins",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetPendingAdmins",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: results,
	}
	return res, nil
}

func (ctrl Controller) ModifyStatusAdminById(req *ActorParam) (ResponseParam, error) {
	var err error
	err = ctrl.uc.ModifyStatusAdminById(req)
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed ModifyStatusAdmin",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success ModifyStatusAdmin",
			Message:      "Success",
			ResponseTime: "",
		},
	}
	return res, nil
}

func (ctrl Controller) ModifyRegisterAdminById(req *RegisterApprovalParam) (ResponseParam, error) {
	var err error
	err = ctrl.uc.ModifyRegisterAdminById(req)
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed ModifyRegisterAdminById",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success ModifyRegisterAdminById",
			Message:      "Success",
			ResponseTime: "",
		},
	}
	return res, nil
}

func (ctrl Controller) RemoveAdminById(req *ActorParam) (ResponseParam, error) {
	var admin, err = ctrl.uc.RemoveAdminById(req)
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed RemoveAdminById",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success RemoveAdminById",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: ActorParam{
			Id:         req.Id,
			Username:   admin.Username,
			RoleId:     admin.RoleId,
			IsVerified: admin.IsVerified,
			IsActive:   admin.IsActive,
		},
	}
	return res, nil
}

func (ctrl Controller) RemoveRegisterAdminById(req *RegisterApprovalParam) (ResponseParam, error) {
	var result, err = ctrl.uc.RemoveRegisterAdminById(req)
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed RemoveRegisterAdminById",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success RemoveRegisterAdminById",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: result,
	}
	return res, nil
}
