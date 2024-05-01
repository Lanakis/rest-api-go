package user_module

import (
	auth_module "authorization/src/modules/auth-module"
	"authorization/src/modules/profile-module" // Импортируйте пакет сервиса профиля
	"authorization/src/modules/user-module/entity"
	"authorization/src/modules/user-module/repository"
	"database/sql"
)

func InitModule(db *sql.DB) entity.IUserService {
	profileService := profile_module.InitModule(db)
	userRepository := repository.NewUserRepository(db)

	userService := NewUserService(userRepository, profileService)
	authService := auth_module.NewAuthService(userService)
	registerRoutes(userService, authService)
	return userService
}
