package main

import (
	"fmt"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/modules/user"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/utils/db"
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	var router = gin.New()

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

	var userRouter = user.NewRouter(dbCrud)
	userRouter.Router(router)

	errRouter := router.Run(":8081")
	if errRouter != nil {
		fmt.Println("error running server", errRouter)
		return
	}

}
