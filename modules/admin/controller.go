package admin

import (
	"back-end-server-dev/dto"
	data_api "back-end-server-dev/function/data-api"
	"errors"
)

type Controller struct {
	uc UseCaseInterface
}

type ControllerInterface interface {
	GetCustomerById(req *CustomerParam) (ResponseParam, error)
	GetCustomersByName(req *CustomerParam) (ResponseParam, error)
	GetCustomersByEmail(req *CustomerParam) (ResponseParam, error)
	GetAllCustomers(req *uint) (ResponseParam, error)
	CreateCustomer(req *CustomerParam) (ResponseParam, error)
	ModifyCustomer(req *CustomerParam) (ResponseParam, error)
	RemoveCustomerById(req *CustomerParam) (ResponseParam, error)

	GetAdminById(req *ActorParam) (ResponseParam, error)
	GetAdminsByUsername(req *ActorParam) (ResponseParam, error)
	GetAllAdmins(req *uint) (ResponseParam, error)
	CreateAdmin(req *ActorParamWithPassword) (ResponseParam, error)
	CreateRegisterAdmin(req *RegisterApprovalParam) (ResponseParam, error)
	ModifyAdmin(req *ActorParamWithPassword) (ResponseParam, error)
}

func (ctrl Controller) GetCustomerById(req *CustomerParam) (ResponseParam, error) {
	var url = "https://reqres.in/api/users?page=2"
	var results []data_api.CustomerParam
	var err error
	var customer CustomerParam

	results, err = data_api.GetData(&url)
	for _, result := range results {
		_ = ctrl.uc.CreateCustomer(&CustomerParam{
			Id:        result.Id,
			FirstName: result.FirstName,
			LastName:  result.LastName,
			Email:     result.Email,
			Avatar:    result.Avatar,
		})
	}

	customer, err = ctrl.uc.GetCustomerById(req)
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

func (ctrl Controller) GetCustomersByName(req *CustomerParam) (ResponseParam, error) {
	var url = "https://reqres.in/api/users?page=2"
	var results []data_api.CustomerParam
	var err error
	var customers []CustomerParam

	results, err = data_api.GetData(&url)
	for _, result := range results {
		_ = ctrl.uc.CreateCustomer(&CustomerParam{
			Id:        result.Id,
			FirstName: result.FirstName,
			LastName:  result.LastName,
			Email:     result.Email,
			Avatar:    result.Avatar,
		})
	}

	customers, err = ctrl.uc.GetCustomersByName(req)
	if err != nil {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed GetCustomersByName",
				Message:      err.Error(),
				ResponseTime: "",
			},
			Data: nil,
		}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetCustomersByName",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: customers,
	}
	return res, nil
}

func (ctrl Controller) GetCustomersByEmail(req *CustomerParam) (ResponseParam, error) {
	var url = "https://reqres.in/api/users?page=2"
	var results []data_api.CustomerParam
	var err error
	var customers []CustomerParam

	results, err = data_api.GetData(&url)
	for _, result := range results {
		_ = ctrl.uc.CreateCustomer(&CustomerParam{
			Id:        result.Id,
			FirstName: result.FirstName,
			LastName:  result.LastName,
			Email:     result.Email,
			Avatar:    result.Avatar,
		})
	}

	customers, err = ctrl.uc.GetCustomersByEmail(req)
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

func (ctrl Controller) GetAllCustomers(req *uint) (ResponseParam, error) {
	var url = "https://reqres.in/api/users?page=2"
	var results []data_api.CustomerParam
	var err error
	var customers []CustomerParam

	results, err = data_api.GetData(&url)
	for _, result := range results {
		_ = ctrl.uc.CreateCustomer(&CustomerParam{
			Id:        result.Id,
			FirstName: result.FirstName,
			LastName:  result.LastName,
			Email:     result.Email,
			Avatar:    result.Avatar,
		})
	}

	if *req < 1 {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed GetAllCustomers",
				Message:      errors.New("page is not valid").Error(),
				ResponseTime: "",
			},
			Data: nil,
		}, errors.New("page is not valid")
	}
	customers, err = ctrl.uc.GetAllCustomers(req)
	if err != nil {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed GetAllCustomers",
				Message:      err.Error(),
				ResponseTime: "",
			},
			Data: nil,
		}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetAllCustomers",
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
			Message:      err.Error(),
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

func (ctrl Controller) GetAdminsByUsername(req *ActorParam) (ResponseParam, error) {
	var admin, err = ctrl.uc.GetAdminsByUsername(req)
	if err != nil {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed GetAdminsByUsername",
				Message:      err.Error(),
				ResponseTime: "",
			},
			Data: nil,
		}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetAdminsByUsername",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: admin,
	}
	return res, nil
}

func (ctrl Controller) GetAllAdmins(req *uint) (ResponseParam, error) {
	if *req < 1 {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed GetAllAdmins",
				Message:      errors.New("page is not valid").Error(),
				ResponseTime: "",
			},
			Data: nil,
		}, errors.New("page is not valid")
	}
	var customers, err = ctrl.uc.GetAllAdmins(req)
	if err != nil {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed GetAllAdmins",
				Message:      err.Error(),
				ResponseTime: "",
			},
			Data: nil,
		}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetAllAdmins",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: customers,
	}
	return res, nil
}

func (ctrl Controller) CreateAdmin(req *ActorParamWithPassword) (ResponseParam, error) {
	var err = ctrl.uc.CreateAdmin(req)
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed CreateAdmin",
			Message:      err.Error(),
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
			Id:         req.Id,
			Username:   req.Username,
			RoleId:     req.RoleId,
			IsVerified: "false",
			IsActive:   "false",
		},
	}
	return res, nil
}

func (ctrl Controller) CreateRegisterAdmin(req *RegisterApprovalParam) (ResponseParam, error) {
	var err = ctrl.uc.CreateRegisterAdmin(req)
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed CreateRegisterAdmin",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success CreateRegisterAdmin",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: RegisterApprovalParam{
			Id:           req.Id,
			AdminId:      req.AdminId,
			SuperAdminId: req.SuperAdminId,
			Status:       "pending",
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
