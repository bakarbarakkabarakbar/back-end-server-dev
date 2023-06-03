package main

import (
	"fmt"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/modules/router"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/utils/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	var engine = gin.New()

	// open connection db
	var dbCrud = db.GormMysql()

	if dbCrud == nil {
		fmt.Println("connection failed to init..!")
		return
	}
	//check connection
	checkDB, err := dbCrud.DB()
	if err != nil {
		log.Fatal(err)
		return
	}

	//ping to database
	var errConn = checkDB.Ping()
	if err != nil {
		log.Fatal(errConn)
		return
	}

	fmt.Println("database connected..!")

	var route = router.NewRouter(dbCrud)
	route.Router(engine)

	errRouter := engine.Run(":8081")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}

}
