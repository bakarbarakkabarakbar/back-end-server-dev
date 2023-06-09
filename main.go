package main

import (
	"back-end-server-dev/modules/router"
	"back-end-server-dev/utils/orm"
	"fmt"
	"github.com/gin-gonic/gin"
)

type Connection struct {
	orm orm.ObjectRelationalMappingInterface
}

func main() {
	var engine = gin.New()

	// open connection to db
	//var dbConnection = connection.NewDatabaseConnection()
	//dbConnection.MySql()
	//check connection

	var dsn = "golang-service-account:STRONG.password79@tcp(host.docker.internal:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=UTC"
	var conn = Connection{orm: orm.NewObjectRelationalMapping(&dsn)}
	var gormInstances, err = conn.orm.Gorm()
	if err != nil {
		fmt.Println("Error init gorm", err)
		return
	}

	var route = router.NewRouter(gormInstances)
	route.Router(engine)

	errRouter := engine.Run(":8081")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}

}
