package data_api

type ResponseParam struct {
	Page       int             `json:"id"`
	PerPage    int             `json:"per_page"`
	Total      int             `json:"total"`
	TotalPages int             `json:"total_pages"`
	Data       []CustomerParam `json:"data"`
}

type CustomerParam struct {
	Id        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Avatar    string `json:"avatar"`
}

type CustomersParam struct {
	Data []CustomerParam `json:"data"`
}
