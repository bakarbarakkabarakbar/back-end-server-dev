package router

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/modules/admin"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/modules/auth"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/modules/customers"
	super_admin "github.com/dibimbing-satkom-indo/onion-architecture-go/modules/super-admin"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Router struct {
	customerReqHandler   customers.RequestHandlerInterface
	adminReqHandler      admin.RequestHandlerInterface
	authReqHandler       auth.RequestHandlerInterface
	superAdminReqHandler super_admin.RequestHandlerInterface
}

func NewRouter(dbCrud *gorm.DB) Router {
	return Router{
		customerReqHandler:   customers.NewRequestHandler(dbCrud),
		adminReqHandler:      admin.NewRequestHandler(dbCrud),
		authReqHandler:       auth.NewRequestHandler(dbCrud),
		superAdminReqHandler: super_admin.NewRequestHandler(dbCrud),
	}
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
	basePathGroup.GET("/login", r.authReqHandler.CreateAuthorization)

	var userPath = "/customer"
	var userPathGroup = router.Group(userPath)
	userPathGroup.GET("/", r.customerReqHandler.GetCustomer)

	var adminPath = "/admin"
	var adminPathGroup = router.Group(adminPath, r.authReqHandler.CheckAuthorization)

	adminPathGroup.POST("/", r.adminReqHandler.CreateAdmin)
	adminPathGroup.GET("/", r.adminReqHandler.GetAdmin)
	adminPathGroup.PUT("/", r.adminReqHandler.ModifyAdmin)

	adminPathGroup.POST("/customer", r.adminReqHandler.CreateCustomer)
	adminPathGroup.GET("/customer", r.adminReqHandler.GetCustomers)
	adminPathGroup.PUT("/customer", r.adminReqHandler.ModifyCustomer)
	adminPathGroup.DELETE("/customer", r.adminReqHandler.RemoveCustomer)

	var superAdminPath = "/super-admin"
	var superAdminPathGroup = router.Group(superAdminPath, r.authReqHandler.CheckAuthorization)
	superAdminPathGroup.POST("/", r.adminReqHandler.CreateAdmin)
	superAdminPathGroup.GET("/", r.adminReqHandler.GetAdmin)
	superAdminPathGroup.PUT("/", r.adminReqHandler.ModifyAdmin)
	superAdminPathGroup.DELETE("/", r.superAdminReqHandler.RemoveAdmin)

	superAdminPathGroup.POST("/customer", r.adminReqHandler.CreateCustomer)
	superAdminPathGroup.GET("/customer", r.adminReqHandler.GetCustomers)
	superAdminPathGroup.PUT("/customer", r.adminReqHandler.ModifyCustomer)
	superAdminPathGroup.DELETE("/customer", r.adminReqHandler.RemoveCustomer)

	superAdminPathGroup.GET("/verified-admin", r.superAdminReqHandler.GetVerifiedAdmin)
	superAdminPathGroup.GET("/active-admin", r.superAdminReqHandler.GetActiveAdmin)
	superAdminPathGroup.PUT("/status-admin", r.superAdminReqHandler.ModifyAdminStatusById)

}
