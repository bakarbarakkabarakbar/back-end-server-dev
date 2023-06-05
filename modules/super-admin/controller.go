package super_admin

import "github.com/dibimbing-satkom-indo/onion-architecture-go/dto"

type Controller struct {
	uc UseCaseInterface
}

type ControllerInterface interface {
	GetVerifiedAdmin() (ResponseParam, error)
	GetActiveAdmin() (ResponseParam, error)
	ModifyAdminStatusById(req *ActorStatusParam) (ResponseParam, error)
	RemoveAdminById(req *ActorParam) (ResponseParam, error)
}

func (ctrl Controller) GetVerifiedAdmin() (ResponseParam, error) {
	var results []ActorStatusParam
	var err error
	results, err = ctrl.uc.GetVerifiedAdmins()
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed GetVerifiedAdmin",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetVerifiedAdmin",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: results,
	}
	return res, nil
}

func (ctrl Controller) GetActiveAdmin() (ResponseParam, error) {
	var results []ActorStatusParam
	var err error
	results, err = ctrl.uc.GetActiveAdmins()
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed GetActiveAdmin",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetActiveAdmin",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: results,
	}
	return res, nil
}

func (ctrl Controller) ModifyAdminStatusById(req *ActorStatusParam) (ResponseParam, error) {
	var err error
	err = ctrl.uc.ModifyAdminStatusById(req)
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed ModifyAdminStatusById",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success ModifyAdminStatusById",
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
