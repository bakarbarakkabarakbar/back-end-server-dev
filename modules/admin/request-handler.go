package admin

import (
	"github.com/dibimbing-satkom-indo/onion-architecture-go/dto"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/repositories"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type RequestHandler struct {
	ctrl Controller
}

func NewRequestHandler(dbCrud *gorm.DB) RequestHandler {
	return RequestHandler{
		ctrl: Controller{
			uc: UseCase{
				customerRepo: repositories.NewCustomerRepo(dbCrud)},
		},
	}
}

func (rh RequestHandler) GetCustomerById(c *gin.Context) {
	var request = CustomerParam{}
	var err = c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	var customerId uint64
	customerId, err = strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	var res ResponseParam
	res, err = rh.ctrl.GetCustomerById(&CustomerParam{Id: uint(customerId)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) GetCustomerByName(c *gin.Context) {
	var request = CustomerParam{}
	var err = c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	var customerFirstName string
	customerFirstName = c.Param("first_name")
	if customerFirstName == "" {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	var customerLastName string
	customerLastName = c.Param("last_name")
	if customerLastName == "" {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	var res ResponseParam
	res, err = rh.ctrl.GetCustomerByName(&CustomerParam{
		FirstName: customerFirstName,
		LastName:  customerLastName})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) GetCustomerByEmail(c *gin.Context) {
	var request = CustomerParam{}
	var err = c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	var customerEmail string
	customerEmail = c.Param("email")
	if customerEmail == "" {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	var res ResponseParam
	res, err = rh.ctrl.GetCustomersByEmail(&CustomerParam{Email: customerEmail})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) CreateCustomer(c *gin.Context) {
	var request = CustomerParam{}
	var err = c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	var res ResponseParam
	res, err = rh.ctrl.CreateCustomer(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) RemoveCustomerById(c *gin.Context) {
	var request = CustomerParam{}
	var err = c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	var customerId uint64
	customerId, err = strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	var res ResponseParam
	res, err = rh.ctrl.RemoveCustomerById(&CustomerParam{Id: uint(customerId)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}
