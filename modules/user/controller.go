package user

import (
	"onion-arch/dto"
)

type Controller struct {
	uc UsecaseInterface
}

type ControllerInterface interface {
	GetUserByID(payload Payload) dto.Response
}
