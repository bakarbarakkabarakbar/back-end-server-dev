package auth

import (
	"back-end-server-dev/dto"
)

type CredentialParam struct {
	id       uint
	username string
	password string
	roleId   uint
}

type ResponseParam struct {
	dto.ResponseMeta
	Data any `json:"data"`
}

type Header struct {
	Bearer string `header:"Authorization"`
}

type ActorSessionParam struct {
	Id      uint   `json:"id"`
	ActorId uint   `json:"user_id"`
	Token   string `json:"token"`
}
