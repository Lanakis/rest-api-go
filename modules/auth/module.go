package auth

import (
	"authorization/api"
	"authorization/modules/user/entity"
)

func InitModule(service entity.IUserService) {

	authService := NewAuthService(service)
	authController := NewAuthController(authService)
	api.HandleFunc("/login", Handler(authController))

}
