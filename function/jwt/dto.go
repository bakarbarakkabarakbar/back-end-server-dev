package jwt

type CredentialParam struct {
	Username string
	Password string
}

type AuthHeader struct {
	Bearer string `header:"Authorization"`
}
