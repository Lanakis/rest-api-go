package user

import (
	"time"
)

type User struct {
	Id         int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	Username   string
	Password   string `json:"-"`
	Role       string
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Head       bool
}
