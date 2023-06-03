package admin

import "github.com/dibimbing-satkom-indo/onion-architecture-go/dto"

type Controller struct {
	uc UseCaseInterface
}

type ControllerInterface interface {
	GetCustomerById(req *CustomerParam) (ResponseParam, error)
	GetCustomerByName(req *CustomerParam) (ResponseParam, error)
	GetCustomersByEmail(req *CustomerParam) (ResponseParam, error)
	CreateCustomer(req *CustomerParam) (ResponseParam, error)
	ModifyCustomer(req *CustomerParam) (ResponseParam, error)
	RemoveCustomerById(req *CustomerParam) (ResponseParam, error)

	GetAdminById(req *ActorParam) (ResponseParam, error)
	CreateAdmin(req *ActorParamWithPassword) (ResponseParam, error)
	ModifyAdmin(req *ActorParamWithPassword) (ResponseParam, error)
	RemoveAdminById(req *ActorParam) (ResponseParam, error)
}

func (ctrl Controller) GetCustomerById(req *CustomerParam) (ResponseParam, error) {
	var customer, err = ctrl.uc.GetCustomerById(req)
	if err != nil {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed GetCustomerById",
				Message:      err.Error(),
				ResponseTime: "",
			},
			Data: nil,
		}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetCustomerById",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: customer,
	}
	return res, nil
}

func (ctrl Controller) GetCustomerByName(req *CustomerParam) (ResponseParam, error) {
	var customers, err = ctrl.uc.GetCustomersByName(req)
	if err != nil {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed GetCustomerByName",
				Message:      err.Error(),
				ResponseTime: "",
			},
			Data: nil,
		}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetCustomerByName",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: customers,
	}
	return res, nil
}

func (ctrl Controller) GetCustomersByEmail(req *CustomerParam) (ResponseParam, error) {
	var customers, err = ctrl.uc.GetCustomersByEmail(req)
	if err != nil {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed GetCustomersByEmail",
				Message:      err.Error(),
				ResponseTime: "",
			},
			Data: nil,
		}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetCustomersByEmail",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: customers,
	}
	return res, nil
}

func (ctrl Controller) CreateCustomer(req *CustomerParam) (ResponseParam, error) {
	var err = ctrl.uc.CreateCustomer(req)
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed CreateCustomer",
			Message:      "Failed",
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success CreateCustomer",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: CustomerParam{
			Id:        req.Id,
			FirstName: req.FirstName,
			LastName:  req.LastName,
			Email:     req.Email,
			Avatar:    req.Avatar,
		},
	}
	return res, nil
}

func (ctrl Controller) ModifyCustomer(req *CustomerParam) (ResponseParam, error) {
	var err = ctrl.uc.ModifyCustomer(req)
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed ModifyCustomer",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success ModifyCustomer",
			Message:      "Success",
			ResponseTime: "",
		},
	}
	return res, nil
}

func (ctrl Controller) RemoveCustomerById(req *CustomerParam) (ResponseParam, error) {
	var customer, err = ctrl.uc.RemoveCustomerById(req)
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed RemoveCustomer",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success RemoveCustomer",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: customer,
	}
	return res, nil
}

func (ctrl Controller) GetAdminById(req *ActorParam) (ResponseParam, error) {
	var admin, err = ctrl.uc.GetAdminById(req)
	if err != nil {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed GetAdminById",
				Message:      err.Error(),
				ResponseTime: "",
			},
			Data: nil,
		}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetAdminById",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: admin,
	}
	return res, nil
}

func (ctrl Controller) CreateAdmin(req *ActorParamWithPassword) (ResponseParam, error) {
	var err = ctrl.uc.CreateAdmin(req)
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed CreateAdmin",
			Message:      "Failed",
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success CreateAdmin",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: ActorParam{
			Username:   req.Username,
			RoleId:     req.Id,
			IsVerified: "false",
			IsActive:   "false",
		},
	}
	return res, nil
}

func (ctrl Controller) ModifyAdmin(req *ActorParamWithPassword) (ResponseParam, error) {
	var err = ctrl.uc.ModifyAdmin(req)
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed ModifyAdmin",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success ModifyAdmin",
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
