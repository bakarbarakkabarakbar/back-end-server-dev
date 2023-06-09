package connection

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DatabaseConnection struct {
	DsnSQL string
}

func NewDatabaseConnection(dsn string) DatabaseConnection {
	return DatabaseConnection{DsnSQL: dsn}
}

//go:generate mockery --name DatabaseInterface
type DatabaseInterface interface {
	MySql() gorm.Dialector
}

func (dc DatabaseConnection) MySql() gorm.Dialector {
	var conn gorm.Dialector
	conn = mysql.Open(dc.DsnSQL)
	return conn
}

//"golang-service-account:STRONG.password79@tcp(34.224.99.112:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=UTC"
