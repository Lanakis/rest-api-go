package dto

type Create struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	Head       bool   `json:"head"`
}

type Response struct {
	Users interface{} `json:"users,omitempty"`
	Count int         `json:"count"`
}

type Update struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	Head       bool   `json:"head"`
}
