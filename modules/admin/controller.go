package admin

import "github.com/dibimbing-satkom-indo/onion-architecture-go/dto"

type Controller struct {
	uc UseCase
}

type ControllerInterface interface {
	GetCustomerById(req *CustomerParam) (ResponseParam, error)
	GetCustomerByName(req *CustomerParam) (ResponseParam, error)
	GetCustomersByEmail(req *CustomerParam) (ResponseParam, error)
	CreateCustomer(req *CustomerParam) (ResponseParam, error)
	RemoveCustomerById(req *CustomerParam) (ResponseParam, error)
}

func (ctrl Controller) GetCustomerById(req *CustomerParam) (ResponseParam, error) {
	var customer, err = ctrl.uc.GetCustomerById(req)
	if err != nil {
		return ResponseParam{}, err
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
		return ResponseParam{}, err
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
		return ResponseParam{}, err
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

func (ctrl Controller) RemoveCustomerById(req *CustomerParam) (ResponseParam, error) {
	var customer, err = ctrl.uc.RemoveCustomerById(req)
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed RemoveCustomerById",
			Message:      "Failed",
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success RemoveCustomerById",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: customer,
	}
	return res, nil
}
