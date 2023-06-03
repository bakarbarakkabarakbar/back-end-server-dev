package entities

import "time"

type Actor struct {
	Id         uint   `gorm:"primary_key"`
	Username   string `gorm:"column:user_id"`
	Password   string `gorm:"column:token"`
	RoleId     uint   `gorm:"column:role_id"`
	IsVerified bool   `gorm:"column:is_verified"`
	IsActive   bool   `gorm:"column:is_active"`
	CreatedAt  time.Time
	ModifiedAt time.Time
}
