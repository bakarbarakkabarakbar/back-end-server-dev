package admin

import "user-management-backend/dto"

type CustomerParam struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type ActorParamWithPassword struct {
	Id         uint   `json:"id"`
	Username   string `json:"username"`
	Password   string `json:"password"`
	RoleId     uint   `json:"role_id"`
	IsVerified string `json:"is_verified"`
	IsActive   string `json:"is_active"`
}

type ActorParam struct {
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

type RegisterApprovalParam struct {
	Id           uint   `json:"id"`
	AdminId      uint   `json:"admin_id"`
	SuperAdminId uint   `json:"super_admin_id"`
	Status       string `json:"status"`
}
