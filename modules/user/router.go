package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	rh RequestHandler
}

func NewRouter() Router {
	return Router{}
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

	var userPath = "/user"
	var userPathGroup = router.Group(userPath)
	userPathGroup.POST("/register", r.rh.CreateUser)
	userPathGroup.GET("/:id", r.rh.GetUsedById)

}
