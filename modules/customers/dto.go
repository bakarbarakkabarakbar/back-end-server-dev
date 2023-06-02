package customers

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/dto"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/entities"
)

type Payload struct {
	ID int
}

type UserParam struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SuccessCreate struct {
	dto.ResponseMeta
	Data UserParam `json:"data"`
}

type FindUser struct {
	dto.ResponseMeta
	Data entities.User `json:"data"`
}
