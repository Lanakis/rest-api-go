package user

import (
	"authorization/modules/auth"
	"authorization/modules/profile"
	"authorization/modules/user/entity"
	"authorization/modules/user/repository"
	"database/sql"
)

func InitModule(db *sql.DB) entity.IUserService {
	profileService := profile.InitModule(db)
	userRepository := repository.NewUserRepository(db)

	userService := NewUserService(userRepository, profileService)
	authService := auth.NewAuthService(userService)
	registerRoutes(userService, authService)
	return userService
}
