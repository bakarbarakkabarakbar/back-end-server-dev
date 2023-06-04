package entities

type ActorRole struct {
	Id       uint   `gorm:"primary_key"`
	RoleName string `gorm:"column:role_name"`
}
