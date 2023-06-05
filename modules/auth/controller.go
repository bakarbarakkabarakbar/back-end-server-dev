package auth

import (
	"crypto/sha1"
	"crypto/subtle"
	"errors"
	"fmt"
	"github.com/dibimbing-satkom-indo/onion-architecture-go/dto"
)

type Controller struct {
	uc UseCaseInterface
}

type ControllerInterface interface {
	GetCredentialByUsername(account *CredentialParam) (CredentialParam, error)
}

func (ctrl Controller) CheckAccountCredential(req *CredentialParam) (ResponseParam, error) {
	var account, err = ctrl.uc.GetCredentialByUsername(req)
	var hash = sha1.New()
	hash.Write([]byte(req.password))

	if err != nil {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed CheckAuthorization",
				Message:      err.Error(),
				ResponseTime: "",
			},
			Data: CredentialParam{}}, err
	}
	if (subtle.ConstantTimeCompare([]byte(fmt.Sprintf("%x", hash.Sum(nil))), []byte(account.password))) != 1 {
		return ResponseParam{
			ResponseMeta: dto.ResponseMeta{
				Success:      false,
				MessageTitle: "Failed CheckAuthorization",
				Message:      "credential not match",
				ResponseTime: "",
			},
			Data: CredentialParam{}}, errors.New("failed CheckAuthorization")
	}

	var res = ResponseParam{
		ResponseMeta: dto.ResponseMeta{
			Success:      true,
			MessageTitle: "Success CheckAuthorization",
			Message:      "Success",
			ResponseTime: "",
		},
		Data: account,
	}
	return res, nil
}
