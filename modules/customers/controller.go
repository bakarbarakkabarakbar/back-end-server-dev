package customers

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/dto"
)

type Controller struct {
	uc UseCase
}

type ControllerInterface interface {
	CreateUser(req UserParam) (any, error)
	GetUserById(id uint) (FindUser, error)
}

func (ctrl Controller) CreateUser(req UserParam) (any, error) {

	user, err := ctrl.uc.CreateUser(req)
	if err != nil {
		return SuccessCreate{}, err
	}
	res := SuccessCreate{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success create customers",
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

func (ctrl Controller) GetUserById(id uint) (FindUser, error) {
	var res FindUser
	var user, err = ctrl.uc.GetUserById(id)
	if err != nil {
		return FindUser{}, err
	}

	res.Data = user
	res.ResponseMeta = dto.ResponseMeta{
		Success:      true,
		MessageTitle: "Success get customers",
		Message:      "Success",
		ResponseTime: "",
	}
	return res, nil
}
