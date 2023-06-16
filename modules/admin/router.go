package admin

import (
	"github.com/gin-gonic/gin"
	"user-management-backend/modules/auth"
)

type Router struct {
	engine *gin.Engine
	rh     RequestHandlerInterface
	auth   auth.RequestHandlerInterface
}

func NewRouter(engine *gin.Engine, rh RequestHandlerInterface, auth auth.RequestHandlerInterface) Router {
	return Router{
		engine: engine,
		rh:     rh,
		auth:   auth,
	}
}

func (r Router) Init(adminPath string) {
	var adminPathGroup = r.engine.Group(adminPath, r.auth.CheckAdminAuthorization)
	adminPathGroup.POST("/", r.rh.CreateAdmin)
	adminPathGroup.GET("/", r.rh.GetAdmin)
	adminPathGroup.PUT("/", r.rh.ModifyAdmin)
	adminPathGroup.GET("/customers", r.rh.GetAllCustomers)
	adminPathGroup.POST("/register", r.rh.CreateRegisterAdmin)

	var adminCustomerPath = "/customer"
	var adminCustomerPathGroup = adminPathGroup.Group(adminCustomerPath, r.auth.CheckAdminAuthorization)
	adminCustomerPathGroup.POST("/", r.rh.CreateCustomer)
	adminCustomerPathGroup.GET("/", r.rh.GetCustomers)
	adminCustomerPathGroup.PUT("/", r.rh.ModifyCustomer)
	adminCustomerPathGroup.DELETE("/", r.rh.RemoveCustomer)
}
