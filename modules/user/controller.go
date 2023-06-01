package user

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/dto"
)

type Controller struct {
	uc UsecaseInterface
}

type ControllerInterface interface {
	GetUserByID(payload Payload) dto.Response
}

func (ctrl Controller) CreateUser(req UserParam) (any, error) {

	user, err := ctrl.uc.CreateUser(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create user",
			Message:      "Success Register",
			ResponseTime: "",
		},
		Data: UserParam{
			Name:     user.Name,
			Email:    user.Email,
			Password: user.Password,
		},
	}
	return res, nil
}
