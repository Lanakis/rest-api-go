package dto

type CreateUserDTO struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	Role       string `json:"role"`
	FirstName  string `json:"first_name"`
	MiddleName string `json:"middle_name"`
	LastName   string `json:"last_name"`
	Age        int    `json:"age"`
	Head       bool   `json:"head"`
}
