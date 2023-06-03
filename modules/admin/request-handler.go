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
				adminRepo:    repositories.NewAdminRepo(dbCrud),
				customerRepo: repositories.NewCustomerRepo(dbCrud),
			},
		},
	}
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
				c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
				return
			}
			c.JSON(http.StatusOK, res)
			return
		case "name":
			if value[0] == "" {
				c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
				return
			}
			var customerName = value[0]

			res, err = rh.ctrl.GetCustomerByName(&CustomerParam{FirstName: customerName})
			if err != nil {
				c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
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
			res, err = rh.ctrl.GetCustomersByEmail(&CustomerParam{Email: customerEmail})
			if err != nil {
				c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
				return
			}
			c.JSON(http.StatusOK, res)
			return
		}
	}
	c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
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
