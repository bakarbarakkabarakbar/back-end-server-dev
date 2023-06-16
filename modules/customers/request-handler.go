package customers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"user-management-backend/dto"
)

type RequestHandler struct {
	ctrl ControllerInterface
}

func NewRequestHandler(ctrl Controller) RequestHandler {
	return RequestHandler{
		ctrl: ctrl}
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
