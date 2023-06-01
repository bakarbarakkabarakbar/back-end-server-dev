package user

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	rq RequestHandlerInterface
}

func NewRouter() Router {
	return Router{}
}

func (r Router) Route(request dto.Request) {

}

func (ru Router) Router(router *gin.Engine) {
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
	userPathGroup.POST("/register", ru.rh.CreateUser)
	userPathGroup.GET("/:id", ru.rh.GetUsedById)

}
