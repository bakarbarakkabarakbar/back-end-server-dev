package super_admin

import (
	"github.com/gin-gonic/gin"
	"user-management-backend/modules/admin"
	"user-management-backend/modules/auth"
)

type Router struct {
	engine *gin.Engine
	rh     RequestHandlerInterface
	admin  admin.RequestHandlerInterface
	auth   auth.RequestHandlerInterface
}

func NewRouter(engine *gin.Engine,
	rh RequestHandlerInterface,
	admin admin.RequestHandlerInterface,
	auth auth.RequestHandlerInterface) Router {
	return Router{
		engine: engine,
		rh:     rh,
		admin:  admin,
		auth:   auth,
	}
}

func (r Router) Init(superAdminPath string) {
	var superAdminPathGroup = r.engine.Group(superAdminPath, r.auth.CheckAdminAuthorization)
	superAdminPathGroup.POST("/", r.admin.CreateAdmin)
	superAdminPathGroup.GET("/", r.admin.GetAdmin)
	superAdminPathGroup.PUT("/", r.admin.ModifyAdmin)
	superAdminPathGroup.DELETE("/", r.rh.RemoveAdmin)
	superAdminPathGroup.GET("/verified-admin", r.rh.GetVerifiedAdmins)
	superAdminPathGroup.GET("/active-admin", r.rh.GetActiveAdmins)
	superAdminPathGroup.PUT("/status-admin", r.rh.ModifyStatusAdmin)
	superAdminPathGroup.GET("/customers", r.admin.GetAllCustomers)
	superAdminPathGroup.GET("/admins", r.admin.GetAllAdmins)

	var superAdminCustomerPath = "/super-admin/customer"
	var superAdminCustomerPathGroup = r.engine.Group(superAdminCustomerPath, r.auth.CheckSuperAdminAuthorization)
	superAdminCustomerPathGroup.POST("/", r.admin.CreateCustomer)
	superAdminCustomerPathGroup.GET("/", r.admin.GetCustomers)
	superAdminCustomerPathGroup.PUT("/", r.admin.ModifyCustomer)
	superAdminCustomerPathGroup.DELETE("/", r.admin.RemoveCustomer)

	var superAdminRegisterPath = "/super-admin/register"
	var superAdminRegisterPathGroup = r.engine.Group(superAdminRegisterPath, r.auth.CheckSuperAdminAuthorization)
	superAdminRegisterPathGroup.POST("/", r.admin.CreateRegisterAdmin)
	superAdminRegisterPathGroup.GET("/", r.rh.GetRegisterAdmin)
	superAdminRegisterPathGroup.GET("/approved", r.rh.GetApprovedAdmins)
	superAdminRegisterPathGroup.GET("/rejected", r.rh.GetRejectedAdmins)
	superAdminRegisterPathGroup.GET("/pending", r.rh.GetPendingAdmins)
	superAdminRegisterPathGroup.PUT("/", r.rh.ModifyRegisterAdmin)
	superAdminRegisterPathGroup.DELETE("/", r.rh.RemoveRegisterAdmin)
}
