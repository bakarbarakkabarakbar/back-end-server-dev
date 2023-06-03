package router

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/modules/admin"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/modules/customers"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Router struct {
	customerReqHandler customers.RequestHandlerInterface
	adminReqHandler    admin.RequestHandler
}

func NewRouter(dbCrud *gorm.DB) Router {
	return Router{
		customerReqHandler: customers.NewRequestHandler(dbCrud),
		adminReqHandler:    admin.NewRequestHandler(dbCrud)}
}

func (r Router) Router(router *gin.Engine) {
	var basePath = "/"
	var basePathGroup = router.Group(basePath)
	basePathGroup.GET("/ping",
		func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "pong",
			})
		})

	var userPath = "/customer"
	var userPathGroup = router.Group(userPath)
	userPathGroup.GET("/profile", r.customerReqHandler.GetCustomer)

	var adminPath = "/admin"
	var adminPathGroup = router.Group(adminPath)
	adminPathGroup.POST("/customer", r.adminReqHandler.CreateCustomer)
	adminPathGroup.GET("/customer", r.adminReqHandler.GetCustomer)

}
