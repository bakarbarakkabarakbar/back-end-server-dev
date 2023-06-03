package admin

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Router struct {
	rh RequestHandler
}

func NewRouter(dbCrud *gorm.DB) Router {
	return Router{
		rh: NewRequestHandler(dbCrud)}
}

func (r Router) Router(router *gin.Engine) {
	var adminPath = "/admin"
	var adminPathGroup = router.Group(adminPath)
	adminPathGroup.POST("/customer", r.rh.CreateCustomer)
	adminPathGroup.GET("/customer", r.rh.GetCustomerById)
}
