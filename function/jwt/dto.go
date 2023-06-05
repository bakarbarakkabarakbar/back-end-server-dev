package jwt

type CredentialParam struct {
	Username string
	Password string
	RoleId   uint
}

type AuthHeader struct {
	Bearer string `header:"Authorization"`
}
