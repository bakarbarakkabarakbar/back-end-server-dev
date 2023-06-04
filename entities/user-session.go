package entities

import "time"

type UserSession struct {
	Id        uint   `gorm:"primary_key"`
	UserId    uint   `gorm:"column:user_id"`
	Token     string `gorm:"column:token"`
	CreatedAt time.Time
	ExpiresAt time.Time
}
