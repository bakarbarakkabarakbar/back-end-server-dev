package entities

import "time"

type Customer struct {
	Id        uint   `gorm:"primary_key"`
	Name      string `gorm:"column:name"`
	Email     string `gorm:"column:email"`
	Password  string `gorm:"column:password"`
	CreatedAt time.Time
	UpdatedAt time.Time
}
