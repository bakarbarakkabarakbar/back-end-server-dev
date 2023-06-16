package auth

import "github.com/gin-gonic/gin"

type Router struct {
	engine *gin.Engine
	rh     RequestHandlerInterface
}

func NewRouter(engine *gin.Engine, rh RequestHandlerInterface) Router {
	return Router{
		engine: engine,
		rh:     rh,
	}
}

func (r Router) Init(authPath string) {
	var authPathGroup = r.engine.Group(authPath)
	authPathGroup.GET("/", r.rh.CreateAuthorization)
}
