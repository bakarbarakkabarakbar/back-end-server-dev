package entities

type Auth struct {
	Id       uint   `gorm:"primary_key"`
	password string `gorm:"column:role_name"`
}
