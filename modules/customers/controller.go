package customers

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/dto"
)

type Controller struct {
	uc UseCaseInterface
}

type ControllerInterface interface {
	GetCustomerById(req *CustomerParam) (ResponseParam, error)
	GetCustomerByEmail(req *CustomerParam) (ResponseParam, error)
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

func (ctrl Controller) GetCustomerByEmail(req *CustomerParam) (ResponseParam, error) {
	var customers, err = ctrl.uc.GetCustomerByEmail(req)
	if err != nil {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed GetCustomerByEmail",
				Message:      err.Error(),
				ResponseTime: "",
			},
			Data: nil,
		}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetCustomerByEmail",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: customers,
	}
	return res, nil
}
