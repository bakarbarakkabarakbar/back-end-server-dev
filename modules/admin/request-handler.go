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
	ctrl ControllerInterface
}

type RequestHandlerInterface interface {
	GetCustomers(c *gin.Context)
	GetAllCustomers(c *gin.Context)
	CreateCustomer(c *gin.Context)
	ModifyCustomer(c *gin.Context)
	RemoveCustomer(c *gin.Context)

	GetAdmin(c *gin.Context)
	GetAllAdmins(c *gin.Context)
	CreateAdmin(c *gin.Context)
	ModifyAdmin(c *gin.Context)
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

func (rh RequestHandler) GetCustomers(c *gin.Context) {
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
		case "name":
			if value[0] == "" {
				c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
				return
			}
			var customerName = value[0]

			res, err = rh.ctrl.GetCustomersByName(&CustomerParam{FirstName: customerName})
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
			res, err = rh.ctrl.GetCustomersByEmail(&CustomerParam{Email: customerEmail})
			if err != nil {
				c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
				return
			}
			c.JSON(http.StatusOK, res)
			return
		}
	}
	c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
}

func (rh RequestHandler) GetAllCustomers(c *gin.Context) {
	var res ResponseParam
	var err error
	var queryParam map[string][]string

	queryParam = c.Request.URL.Query()

	for key, value := range queryParam {
		switch key {
		case "page":
			var page uint64
			var pageConverted uint
			page, err = strconv.ParseUint(value[0], 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
				return
			}
			pageConverted = uint(page)
			res, err = rh.ctrl.GetAllCustomers(&pageConverted)
			if err != nil {
				c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
				return
			}
			c.JSON(http.StatusOK, res)
			return
		}
	}
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
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) ModifyCustomer(c *gin.Context) {
	var request = CustomerParam{}
	var err = c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	var res ResponseParam
	res, err = rh.ctrl.ModifyCustomer(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) RemoveCustomer(c *gin.Context) {
	var err error
	var customerId uint64
	customerId, err = strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	var res ResponseParam
	res, err = rh.ctrl.RemoveCustomerById(&CustomerParam{Id: uint(customerId)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) GetAdmin(c *gin.Context) {
	var res ResponseParam
	var err error
	var queryParam map[string][]string

	queryParam = c.Request.URL.Query()

	for key, value := range queryParam {
		switch key {
		case "id":
			var adminId uint64
			adminId, err = strconv.ParseUint(value[0], 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
				return
			}
			res, err = rh.ctrl.GetAdminById(&ActorParam{Id: uint(adminId)})
			if err != nil {
				c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
				return
			}
			c.JSON(http.StatusOK, res)
			return
		case "username":
			if value[0] == "" {
				c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
				return
			}
			var username = value[0]

			res, err = rh.ctrl.GetAdminsByUsername(&ActorParam{Username: username})
			if err != nil {
				c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
				return
			}
			c.JSON(http.StatusOK, res)
			return
		}
	}
}

func (rh RequestHandler) GetAllAdmins(c *gin.Context) {
	var res ResponseParam
	var err error
	var queryParam map[string][]string

	queryParam = c.Request.URL.Query()

	for key, value := range queryParam {
		switch key {
		case "page":
			var page uint64
			var pageConverted uint
			page, err = strconv.ParseUint(value[0], 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
				return
			}
			pageConverted = uint(page)
			res, err = rh.ctrl.GetAllAdmins(&pageConverted)
			if err != nil {
				c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
				return
			}
			c.JSON(http.StatusOK, res)
			return
		}
	}
}

func (rh RequestHandler) CreateAdmin(c *gin.Context) {

	var request = ActorParamWithPassword{}
	var err = c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	var res ResponseParam
	res, err = rh.ctrl.CreateAdmin(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) ModifyAdmin(c *gin.Context) {
	var request = ActorParamWithPassword{}
	var err = c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	var res ResponseParam
	res, err = rh.ctrl.ModifyAdmin(&request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}
