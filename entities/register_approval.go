package entities

type RegisterApproval struct {
	Id           uint   `gorm:"primary_key"`
	AdminId      uint   `gorm:"column:admin_id"`
	SuperAdminId uint   `gorm:"column:super_admin_id"`
	Status       string `gorm:"column:status"`
}
