package auth

import (
	"back-end-server-dev/dto"
	"back-end-server-dev/function/jwt"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type RequestHandler struct {
	ctrl ControllerInterface
}

func NewRequestHandler(ctrl Controller) RequestHandler {
	return RequestHandler{ctrl: ctrl}
}

type RequestHandlerInterface interface {
	CheckSuperAdminAuthorization(c *gin.Context)
	CheckAdminAuthorization(c *gin.Context)
	CreateAuthorization(c *gin.Context)
}

func (rh RequestHandler) CheckSuperAdminAuthorization(c *gin.Context) {
	var header jwt.AuthHeader
	var err error
	err = c.BindHeader(&header)
	if err != nil {
		fmt.Println("Error binding authorization")
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorInvalidDataWithMessage("Error binding authorization", err.Error()))
		c.Abort()
		return
	}

	if header.Bearer == "" {
		fmt.Println("Error no authorization token")
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorInvalidDataWithMessage("Error no authorization token", err.Error()))
		c.Abort()
		return
	}

	_, err = rh.ctrl.GetLastActorSessionByToken(&ActorSessionParam{Token: header.Bearer})
	if err != nil {
		fmt.Println("Error account credentials")
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorInvalidDataWithMessage("Error account recorded credentials", err.Error()))
		c.Abort()
		return
	}

	header, err = jwt.VerifySuperAdminToken(&header)
	if err != nil {
		fmt.Println("Error account credentials")
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorInvalidDataWithMessage("Error account credentials", err.Error()))
		c.Abort()
		return
	}
}

func (rh RequestHandler) CheckAdminAuthorization(c *gin.Context) {
	var header jwt.AuthHeader
	var err error
	err = c.BindHeader(&header)
	if err != nil {
		fmt.Println("Error binding authorization")
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorInvalidDataWithMessage("Error binding authorization", err.Error()))
		c.Abort()
		return
	}

	if header.Bearer == "" {
		fmt.Println("Error no authorization token")
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorInvalidDataWithMessage("Error no authorization token", err.Error()))
		c.Abort()
		return
	}

	_, err = rh.ctrl.GetLastActorSessionByToken(&ActorSessionParam{Token: header.Bearer})
	if err != nil {
		fmt.Println("Error account credentials")
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorInvalidDataWithMessage("Error account recorded credentials", err.Error()))
		c.Abort()
		return
	}

	header, err = jwt.VerifyAdminToken(&header)
	if err != nil {
		fmt.Println("Error account credentials")
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorInvalidDataWithMessage("Error account credentials", err.Error()))
		c.Abort()
		return
	}
}

func (rh RequestHandler) CreateAuthorization(c *gin.Context) {
	username, password, ok := c.Request.BasicAuth()
	if !ok {
		fmt.Println("Error parsing basic auth")
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorInvalidDataWithMessage("Error parsing basic auth", ""))
		c.Abort()
		return
	}

	var res, err = rh.ctrl.CheckAccountCredential(&CredentialParam{
		username: username,
		password: password,
	})

	if err != nil {
		fmt.Println("Error account credentials")
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorWithResponse(res.ResponseMeta))
		c.Abort()
		return
	}
	var header jwt.AuthHeader
	header, err = jwt.GenerateToken(&jwt.CredentialParam{
		Username: username,
		Password: password,
		RoleId:   res.Data.(CredentialParam).roleId,
	})
	if err != nil {
		fmt.Println("Error account credentials")
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorInvalidDataWithMessage("Error account credentials", err.Error()))
		c.Abort()
		return
	}
	res, err = rh.ctrl.CreateActorSession(&ActorSessionParam{
		ActorId: res.Data.(CredentialParam).id,
		Token:   header.Bearer,
	})
	if err != nil {
		fmt.Println("Error account credentials")
		c.JSON(http.StatusUnauthorized, dto.DefaultErrorInvalidDataWithMessage("Error account credentials", err.Error()))
		c.Abort()
		return
	}
	c.Header("Authorization", header.Bearer)
}
