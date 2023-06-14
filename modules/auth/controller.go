package auth

import (
	"back-end-server-dev/dto"
	"crypto/sha1"
	"crypto/subtle"
	"errors"
	"fmt"
)

type Controller struct {
	uc UseCaseInterface
}

func NewController(uc UseCase) Controller {
	return Controller{uc: uc}
}

type ControllerInterface interface {
	CheckAccountCredential(req *CredentialParam) (ResponseParam, error)
	GetLastActorSessionByToken(req *ActorSessionParam) (ResponseParam, error)
	CreateActorSession(req *ActorSessionParam) (ResponseParam, error)
}

func (ctrl Controller) CheckAccountCredential(req *CredentialParam) (ResponseParam, error) {
	var account, err = ctrl.uc.GetCredentialByUsername(req)
	var hash = sha1.New()
	hash.Write([]byte(req.password))
	if err != nil {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed CheckAdminAuthorization",
				Message:      err.Error(),
				ResponseTime: "",
			},
			Data: CredentialParam{}}, err
	}
	if (subtle.ConstantTimeCompare([]byte(fmt.Sprintf("%x", hash.Sum(nil))), []byte(account.password))) != 1 {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed CheckAdminAuthorization",
				Message:      "credential not match",
				ResponseTime: "",
			},
			Data: CredentialParam{}}, errors.New("failed CheckAdminAuthorization")
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success CheckAdminAuthorization",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: account,
	}
	return res, nil
}

func (ctrl Controller) GetLastActorSessionByToken(req *ActorSessionParam) (ResponseParam, error) {
	var account, err = ctrl.uc.GetLastActorSessionByToken(req)

	if err != nil {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed GetLastActorSessionByToken",
				Message:      err.Error(),
				ResponseTime: "",
			},
			Data: CredentialParam{}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success GetLastActorSessionByToken",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: account,
	}
	return res, nil
}

func (ctrl Controller) CreateActorSession(req *ActorSessionParam) (ResponseParam, error) {
	var err = ctrl.uc.CreateActorSession(req)
	if err != nil {
		return ResponseParam{ResponseMeta: dto.ResponseMeta{
			Success:      false,
			MessageTitle: "Failed CreateActorSession",
			Message:      err.Error(),
			ResponseTime: "",
		}}, err
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success CreateActorSession",
			Message:      "Success",
			ResponseTime: "",
		},
	}
	return res, nil
}
