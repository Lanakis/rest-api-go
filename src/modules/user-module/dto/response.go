package dto

type Response struct {
	Users interface{} `json:"users,omitempty"`
	Count int         `json:"count"`
}
