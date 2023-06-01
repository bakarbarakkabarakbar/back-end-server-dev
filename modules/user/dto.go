package user

type Payload struct {
	ID int
}

type UserParam struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
