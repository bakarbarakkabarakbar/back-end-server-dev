package super_admin

import "github.com/dibimbing-satkom-indo/onion-architecture-go/dto"

type ActorStatusParam struct {
	Id         uint   `json:"id"`
	Username   string `json:"username"`
	RoleId     uint   `json:"role_id"`
	IsVerified string `json:"is_verified"`
	IsActive   string `json:"is_active"`
}

type ResponseParam struct {
	dto.ResponseMeta
	Data any `json:"data"`
}

type ActorParam struct {
	Id         uint   `json:"id"`
	Username   string `json:"username"`
	RoleId     uint   `json:"role_id"`
	IsVerified string `json:"is_verified"`
	IsActive   string `json:"is_active"`
}
