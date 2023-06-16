package customers

import (
	"user-management-backend/dto"
)

type CustomerParam struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type ResponseParam struct {
	dto.ResponseMeta
	Data any `json:"data"`
}
