package user

import (
	"authorization/api"
	"authorization/modules/auth"
	"authorization/modules/user/entity"
	"authorization/utils/filter"
	"net/http"
)

func registerRoutes(userService entity.IUserService, authService *auth.AuthService) {
	// Получаем экземпляр миддлвар для проверки токена
	userMiddleware := func(next http.HandlerFunc) http.HandlerFunc {
		return authService.Middleware(next)
	}
	// Регистрируем маршруты с использованием миддлвар для проверки токена
	api.HandleFunc("/user/", filter.Middleware(userMiddleware(Handler(userService)), 10))

}
