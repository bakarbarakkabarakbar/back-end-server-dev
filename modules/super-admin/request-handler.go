package super_admin

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

func NewRequestHandler(dbCrud *gorm.DB) RequestHandler {
	return RequestHandler{
		ctrl: Controller{
			uc: UseCase{
				superAdminRepo: repositories.NewSuperAdminRepo(dbCrud),
				adminRepo:      repositories.NewAdminRepo(dbCrud),
			},
		},
	}
}

type RequestHandlerInterface interface {
	GetVerifiedAdmin(c *gin.Context)
	GetActiveAdmin(c *gin.Context)
	ModifyAdminStatusById(c *gin.Context)
	RemoveAdmin(c *gin.Context)
}

func (rh RequestHandler) GetVerifiedAdmin(c *gin.Context) {
	var res ResponseParam
	var err error

	res, err = rh.ctrl.GetVerifiedAdmin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) GetActiveAdmin(c *gin.Context) {
	var res ResponseParam
	var err error

	res, err = rh.ctrl.GetActiveAdmin()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) ModifyAdminStatusById(c *gin.Context) {
	var actorStatus ActorStatusParam
	var err error
	err = c.Bind(&actorStatus)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	var res ResponseParam
	res, err = rh.ctrl.ModifyAdminStatusById(&actorStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) RemoveAdmin(c *gin.Context) {
	var err error
	var adminId uint64
	adminId, err = strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	var res ResponseParam
	res, err = rh.ctrl.RemoveAdminById(&ActorParam{Id: uint(adminId)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}
