package customers

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

func (r Router) Init(customerPath string) {
	var customerPathGroup = r.engine.Group(customerPath)
	customerPathGroup.GET("/", r.rh.GetCustomer)
}
