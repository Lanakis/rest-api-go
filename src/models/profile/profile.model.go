package profile

import "time"

type Profile struct {
	Id         int
	CreatedAt  time.Time
	UpdatedAt  time.Time
	FirstName  string
	MiddleName string
	LastName   string
	Age        int
	Head       bool
}
