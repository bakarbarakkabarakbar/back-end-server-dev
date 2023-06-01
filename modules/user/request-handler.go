package user

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

	"github.com/dibimbing-satkom-indo/onion-architecture-go/dto"
)

type RequestHandler struct {
	ctrl ControllerInterface
}

type RequestHandlerInterface interface {
	GetUserByID(request dto.Request) dto.Response
}

func (rh RequestHandler) GetUsedById(c *gin.Context) {
	var request = UserParam{}
	var err = c.BindQuery(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	var userId uint64
	userId, err = strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	var res FindUser
	res, err = rh.ctrl.GetUserById(uint(userId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) CreateUser(c *gin.Context) {
	var request = UserParam{}
	var err = c.Bind(&request)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	res, err := rh.ctrl.CreateUser(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorResponse())
		return
	}
	c.JSON(http.StatusOK, res)
