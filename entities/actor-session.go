package entities

type ActorSession struct {
	Id     uint   `gorm:"primary_key"`
	UserId uint   `gorm:"column:user_id"`
	Token  string `gorm:"column:token"`
}
