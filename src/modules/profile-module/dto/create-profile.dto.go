package dto

type CreateProfileDTO struct {
	Name       string `json:"name"`
	SecondName string `json:"secondName"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Head       bool   `json:"head"`
}
