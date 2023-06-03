package customers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

type Router struct {
	rh RequestHandlerInterface
}

func NewRouter(dbCrud *gorm.DB) Router {
	return Router{
		rh: NewRequestHandler(dbCrud)}
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

	var userPath = "/customers"
	var userPathGroup = router.Group(userPath)
	userPathGroup.GET("/", r.rh.GetCustomer)

}
