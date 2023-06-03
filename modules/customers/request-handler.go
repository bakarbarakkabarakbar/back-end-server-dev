package customers

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"

	"github.com/dibimbing-satkom-indo/onion-architecture-go/dto"
)

type RequestHandler struct {
	ctrl ControllerInterface
}

func NewRequestHandler(dbCrud *gorm.DB) RequestHandler {
	return RequestHandler{
		ctrl: Controller{
			uc: UseCase{
				customerRepo: repositories.NewCustomerRepo(dbCrud)},
		},
	}
}

type RequestHandlerInterface interface {
	GetCustomer(c *gin.Context)
}

func (rh RequestHandler) GetCustomer(c *gin.Context) {
	var customerId uint64
	var res ResponseParam
	var err error
	var customerEmail string
	customerId, err = strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		customerEmail = c.Param("email")
		if customerEmail == "" {
			c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
			return
		}

		res, err = rh.ctrl.GetCustomerByEmail(&CustomerParam{Email: customerEmail})
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
			return
		}
		c.JSON(http.StatusOK, res)
	} else {
		res, err = rh.ctrl.GetCustomerById(&CustomerParam{Id: uint(customerId)})
		if err != nil {
			c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
