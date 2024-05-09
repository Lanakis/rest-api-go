package models

import "time"

type User struct {
	Id         int       `json:"id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Username   string    `json:"username"`
	Password   string    `json:"-"`
	Role       string    `json:"role"`
	FirstName  string    `json:"first_name"`
	MiddleName string    `json:"middle_name"`
	LastName   string    `json:"last_name"`
	Age        int       `json:"age"`
	Head       bool      `json:"head"`
}
