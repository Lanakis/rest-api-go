package utils

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data,omitempty"`
}

type DeleteResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}
