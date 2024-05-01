package dto

type Response struct {
	Profiles interface{} `json:"profiles,omitempty"`
	Count    int         `json:"count"`
}
