package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"os"
	"strconv"
	"user-management-backend/modules/admin"
	"user-management-backend/modules/auth"
	"user-management-backend/modules/customers"
	superAdmin "user-management-backend/modules/super-admin"
	"user-management-backend/repositories"
	"user-management-backend/utils/connection"
	"user-management-backend/utils/orm"
)

func GetEnv(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("%s NOT SET\n", key))
	} else {
		return val
	}
}

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host petstore.swagger.io
// @BasePath /v2
func main() {
	var err error
	var MYSQL_SERVER_HOST = GetEnv("MYSQL_SERVER_HOST")
	var MYSQL_SERVER_PORT = GetEnv("MYSQL_SERVER_PORT")
	var MYSQL_SERVER_SCHEMA = GetEnv("MYSQL_SERVER_SCHEMA")
	var MYSQL_SERVER_USER = GetEnv("MYSQL_SERVER_USER")
	var MYSQL_SERVER_PASSWORD = GetEnv("MYSQL_SERVER_PASSWORD")

	var API_PORT = GetEnv("API_PORT")
	var KAFKA_SERVER_HOST = GetEnv("KAFKA_SERVER_HOST")
	var KAFKA_SERVER_PORT = GetEnv("KAFKA_SERVER_PORT")
	var KAFKA_TOPIC = GetEnv("KAFKA_TOPIC")
	var KAFKA_PARTITION int
	KAFKA_PARTITION, err = strconv.Atoi(GetEnv("KAFKA_PARTITION"))
	if err != nil {
		panic(err)
	}
	//var SWAGGER_PORT = GetEnv("SWAGGER_PORT")
	//var APP_PORT = GetEnv("APP_PORT")

	var kafkaBroker = connection.NewBrokerSegmentIo(
		KAFKA_SERVER_HOST,
		KAFKA_SERVER_PORT,
		KAFKA_TOPIC,
		KAFKA_PARTITION)

	//err = kafkaBroker.WriteMessage("coba", "testing kafka")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	err = kafkaBroker.GetMessages()
	if err != nil {
		fmt.Println(err)
		return
	}

	gin.SetMode(gin.ReleaseMode)
	var engine = gin.New()
	// golang-service-account:STRONG.password79@tcp(34.224.99.112:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=UTC
	var dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=UTC",
		MYSQL_SERVER_USER,
		MYSQL_SERVER_PASSWORD,
		MYSQL_SERVER_HOST,
		MYSQL_SERVER_PORT,
		MYSQL_SERVER_SCHEMA)
	//fmt.Println(dsn)
	//var dsn = "root:1234QWERasdf.@tcp(localhost:3306)/miniproject?charset=utf8mb4&parseTime=True&loc=UTC"
	var dbConn = connection.NewDatabase(dsn)
	var NewOrm = orm.NewObjectRelationalMapping(dbConn)
	var gormInstances *gorm.DB
	gormInstances, err = NewOrm.Gorm()
	if err != nil {
		panic(err)
	}

	var adminRepo = repositories.NewAdminRepo(gormInstances)
	var authRepo = repositories.NewAuthRepo(gormInstances)
	var customerRepo = repositories.NewCustomerRepo(gormInstances)
	var superAdminRepo = repositories.NewSuperAdminRepo(gormInstances)

	var adminUseCase = admin.NewUseCase(adminRepo, customerRepo)
	var authUseCase = auth.NewUseCase(authRepo)
	var customerUseCase = customers.NewUseCase(customerRepo)
	var superAdminUseCase = superAdmin.NewUseCase(superAdminRepo, adminRepo)

	var adminController = admin.NewController(adminUseCase)
	var authController = auth.NewController(authUseCase)
	var customerController = customers.NewController(customerUseCase)
	var superAdminController = superAdmin.NewController(superAdminUseCase)

	var adminReqHandler = admin.NewRequestHandler(adminController)
	var authReqHandler = auth.NewRequestHandler(authController)
	var customerReqHandler = customers.NewRequestHandler(customerController)
	var superAdminReqHandler = superAdmin.NewRequestHandler(superAdminController)

	var adminRouter = admin.NewRouter(engine, adminReqHandler, authReqHandler)
	var authRouter = auth.NewRouter(engine, authReqHandler)
	var customerRouter = customers.NewRouter(engine, customerReqHandler)
	var superAdminRouter = superAdmin.NewRouter(engine, superAdminReqHandler, adminReqHandler, authReqHandler)

	adminRouter.Init("/admin")
	authRouter.Init("/login")
	customerRouter.Init("/customer")
	superAdminRouter.Init("/super-admin")

	//var route = router.NewRouter(gormInstances)
	//route.Router(engine)

	err = engine.Run(fmt.Sprintf(":%s", API_PORT))
	//err = engine.Run(":8081")
	if err != nil {
		panic(err)
	}

}
