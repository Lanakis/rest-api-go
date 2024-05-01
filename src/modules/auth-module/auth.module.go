package auth_module

import (
	"authorization/src/api"
	"authorization/src/modules/user-module/entity"
)

func InitModule(service entity.IUserService) {

	authService := NewAuthService(service)
	authController := NewAuthController(authService)
	api.HandleFunc("/login", Handler(authController))

}
