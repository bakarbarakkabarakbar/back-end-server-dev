package super_admin

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
		ctrl: ctrl,
	}
}

type RequestHandlerInterface interface {
	GetVerifiedAdmins(c *gin.Context)
	GetActiveAdmins(c *gin.Context)
	GetRegisterAdmin(c *gin.Context)
	GetApprovedAdmins(c *gin.Context)
	GetRejectedAdmins(c *gin.Context)
	GetPendingAdmins(c *gin.Context)
	ModifyStatusAdmin(c *gin.Context)
	ModifyRegisterAdmin(c *gin.Context)
	RemoveAdmin(c *gin.Context)
	RemoveRegisterAdmin(c *gin.Context)
}

func (rh RequestHandler) GetVerifiedAdmins(c *gin.Context) {
	var res ResponseParam
	var err error

	res, err = rh.ctrl.GetVerifiedAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) GetActiveAdmins(c *gin.Context) {
	var res ResponseParam
	var err error

	res, err = rh.ctrl.GetActiveAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) GetRegisterAdmin(c *gin.Context) {
	var res ResponseParam
	var err error
	var queryParam map[string][]string

	queryParam = c.Request.URL.Query()

	for key, value := range queryParam {
		switch key {
		case "id":
			var id uint64
			id, err = strconv.ParseUint(value[0], 10, 64)
			if err != nil {
				c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
				return
			}
			res, err = rh.ctrl.GetRegisterAdminById(&RegisterApprovalParam{Id: uint(id)})
			if err != nil {
				c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
				return
			}
			c.JSON(http.StatusOK, res)
			return
		}
	}
}

func (rh RequestHandler) GetApprovedAdmins(c *gin.Context) {
	var res ResponseParam
	var err error

	res, err = rh.ctrl.GetApprovedAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) GetRejectedAdmins(c *gin.Context) {
	var res ResponseParam
	var err error

	res, err = rh.ctrl.GetRejectedAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) GetPendingAdmins(c *gin.Context) {
	var res ResponseParam
	var err error

	res, err = rh.ctrl.GetPendingAdmins()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) ModifyStatusAdmin(c *gin.Context) {
	var actorStatus ActorParam
	var err error
	err = c.Bind(&actorStatus)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	var res ResponseParam
	res, err = rh.ctrl.ModifyStatusAdminById(&actorStatus)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}

func (rh RequestHandler) ModifyRegisterAdmin(c *gin.Context) {
	var adminRegister RegisterApprovalParam
	var err error
	err = c.Bind(&adminRegister)

	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}
	var res ResponseParam
	res, err = rh.ctrl.ModifyRegisterAdminById(&adminRegister)
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

func (rh RequestHandler) RemoveRegisterAdmin(c *gin.Context) {
	var err error
	var id uint64
	id, err = strconv.ParseUint(c.Query("id"), 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, dto.DefaultBadRequestResponse())
		return
	}

	var res ResponseParam
	res, err = rh.ctrl.RemoveRegisterAdminById(&RegisterApprovalParam{Id: uint(id)})
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.DefaultErrorWithResponse(res.ResponseMeta))
		return
	}
	c.JSON(http.StatusOK, res)
}
