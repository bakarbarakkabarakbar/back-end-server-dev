package auth

import "github.com/dibimbing-satkom-indo/onion-architecture-go/dto"

type CredentialParam struct {
	username string
	password string
}

type ResponseParam struct {
	dto.ResponseMeta
	Data CredentialParam
}

type Header struct {
	Bearer string `header:"Authorization"`
}
