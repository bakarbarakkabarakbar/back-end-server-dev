package super_admin

import "back-end-server-dev/dto"

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
	Id           uint   `gorm:"primary_key"`
	AdminId      uint   `gorm:"column:admin_id"`
	SuperAdminId uint   `gorm:"column:super_admin_id"`
	Status       string `gorm:"column:status"`
}
