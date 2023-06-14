package router

import (
	"back-end-server-dev/modules/admin"
	"back-end-server-dev/modules/auth"
	"back-end-server-dev/modules/customers"
	superAdmin "back-end-server-dev/modules/super-admin"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

type Router struct {
	customerReqHandler   customers.RequestHandlerInterface
	adminReqHandler      admin.RequestHandlerInterface
	authReqHandler       auth.RequestHandlerInterface
	superAdminReqHandler superAdmin.RequestHandlerInterface
}

//func NewRouter(gorm *gorm.DB) Router {
//	return Router{
//		customerReqHandler:   customers.NewRequestHandler(gorm),
//		adminReqHandler:      admin.NewRequestHandler(gorm),
//		authReqHandler:       auth.NewRequestHandler(gorm),
//		superAdminReqHandler: superAdmin.NewRequestHandler(gorm),
//	}
//}

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
	basePathGroup.GET("/swagger", ginSwagger.WrapHandler(swaggerFiles.Handler))

	var userPath = "/customer"
	var userPathGroup = router.Group(userPath)
	userPathGroup.GET("/", r.customerReqHandler.GetCustomer)

	var adminPath = "/admin"
	var adminPathGroup = router.Group(adminPath, r.authReqHandler.CheckAdminAuthorization)
	adminPathGroup.POST("/", r.adminReqHandler.CreateAdmin)
	adminPathGroup.GET("/", r.adminReqHandler.GetAdmin)
	adminPathGroup.PUT("/", r.adminReqHandler.ModifyAdmin)
	adminPathGroup.GET("/customers", r.adminReqHandler.GetAllCustomers)
	adminPathGroup.POST("/register", r.adminReqHandler.CreateRegisterAdmin)

	var adminCustomerPath = "/admin/customer"
	var adminCustomerPathGroup = router.Group(adminCustomerPath, r.authReqHandler.CheckAdminAuthorization)
	adminCustomerPathGroup.POST("/", r.adminReqHandler.CreateCustomer)
	adminCustomerPathGroup.GET("/", r.adminReqHandler.GetCustomers)
	adminCustomerPathGroup.PUT("/", r.adminReqHandler.ModifyCustomer)
	adminCustomerPathGroup.DELETE("/", r.adminReqHandler.RemoveCustomer)

	var superAdminPath = "/super-admin"
	var superAdminPathGroup = router.Group(superAdminPath, r.authReqHandler.CheckSuperAdminAuthorization)
	superAdminPathGroup.POST("/", r.adminReqHandler.CreateAdmin)
	superAdminPathGroup.GET("/", r.adminReqHandler.GetAdmin)
	superAdminPathGroup.PUT("/", r.adminReqHandler.ModifyAdmin)
	superAdminPathGroup.DELETE("/", r.superAdminReqHandler.RemoveAdmin)
	superAdminPathGroup.GET("/verified-admin", r.superAdminReqHandler.GetVerifiedAdmins)
	superAdminPathGroup.GET("/active-admin", r.superAdminReqHandler.GetActiveAdmins)
	superAdminPathGroup.PUT("/status-admin", r.superAdminReqHandler.ModifyStatusAdmin)
	superAdminPathGroup.GET("/customers", r.adminReqHandler.GetAllCustomers)
	superAdminPathGroup.GET("/admins", r.adminReqHandler.GetAllAdmins)

	var superAdminCustomerPath = "/super-admin/customer"
	var superAdminCustomerPathGroup = router.Group(superAdminCustomerPath, r.authReqHandler.CheckSuperAdminAuthorization)
	superAdminCustomerPathGroup.POST("/", r.adminReqHandler.CreateCustomer)
	superAdminCustomerPathGroup.GET("/", r.adminReqHandler.GetCustomers)
	superAdminCustomerPathGroup.PUT("/", r.adminReqHandler.ModifyCustomer)
	superAdminCustomerPathGroup.DELETE("/", r.adminReqHandler.RemoveCustomer)

	var superAdminRegisterPath = "/super-admin/register"
	var superAdminRegisterPathGroup = router.Group(superAdminRegisterPath, r.authReqHandler.CheckSuperAdminAuthorization)
	superAdminRegisterPathGroup.POST("/", r.adminReqHandler.CreateRegisterAdmin)
	superAdminRegisterPathGroup.GET("/", r.superAdminReqHandler.GetRegisterAdmin)
	superAdminRegisterPathGroup.GET("/approved", r.superAdminReqHandler.GetApprovedAdmins)
	superAdminRegisterPathGroup.GET("/rejected", r.superAdminReqHandler.GetRejectedAdmins)
	superAdminRegisterPathGroup.GET("/pending", r.superAdminReqHandler.GetPendingAdmins)
	superAdminRegisterPathGroup.PUT("/", r.superAdminReqHandler.ModifyRegisterAdmin)
	superAdminRegisterPathGroup.DELETE("/", r.superAdminReqHandler.RemoveRegisterAdmin)
}
