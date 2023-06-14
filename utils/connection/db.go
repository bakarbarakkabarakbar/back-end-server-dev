package connection

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	DsnSQL string
}

func NewDatabase(dsn string) Database {
	return Database{DsnSQL: dsn}
}

//go:generate mockery --name DatabaseInterface
type DatabaseInterface interface {
	MySql() gorm.Dialector
}

func (dc Database) MySql() gorm.Dialector {
	var conn gorm.Dialector
	conn = mysql.Open(dc.DsnSQL)
	return conn
}

//"golang-service-account:STRONG.password79@tcp(34.224.99.112:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=UTC"
