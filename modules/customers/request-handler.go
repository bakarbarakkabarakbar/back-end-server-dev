package customers

import (
	"back-end-server-dev/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"back-end-server-dev/dto"
)

type RequestHandler struct {
	ctrl ControllerInterface
}

func NewRequestHandler(gorm *gorm.DB) RequestHandler {
	return RequestHandler{
		ctrl: Controller{
			uc: UseCase{
				customerRepo: repositories.NewCustomerRepo(gorm)},
		},
	}
}

type RequestHandlerInterface interface {
	GetCustomer(c *gin.Context)
}

func (rh RequestHandler) GetCustomer(c *gin.Context) {
	var res ResponseParam
	var err error
	var queryParam map[string][]string

	queryParam = c.Request.URL.Query()

	for key, value := range queryParam {
		switch key {
		case "id":
			var customerId uint64
			customerId, err = strconv.ParseUint(value[0], 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
				return
			}
			res, err = rh.ctrl.GetCustomerById(&CustomerParam{Id: uint(customerId)})
			if err != nil {
				c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
				return
			}
			c.JSON(http.StatusOK, res)
			return
		case "email":
			if value[0] == "" {
				c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
				return
			}
			var customerEmail = value[0]
			res, err = rh.ctrl.GetCustomerByEmail(&CustomerParam{Email: customerEmail})
			if err != nil {
				c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
				return
			}
			c.JSON(http.StatusOK, res)
			return
		}
	}
}
