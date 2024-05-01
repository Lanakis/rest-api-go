package user_module

import (
	"authorization/src/api"
	auth_module "authorization/src/modules/auth-module"
	"authorization/src/modules/user-module/entity"
	"authorization/src/utils/filter"
	"net/http"
)

func registerRoutes(userService entity.IUserService, authService *auth_module.AuthService) {
	// Получаем экземпляр миддлвар для проверки токена
	userMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
		return authService.Middleware(next)
	}

	// Регистрируем маршруты с использованием миддлвар для проверки токена
	api.HandleFunc("/user/", filter.Middleware(userMiddleware(Handler(userService)), 10))

}
