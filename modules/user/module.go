package user

import (
	"gin-gorm/modules/auth"
	"gin-gorm/modules/profile"
	"gin-gorm/modules/user/entity"
	"gin-gorm/modules/user/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitModule(db *gorm.DB, router *gin.Engine) entity.IUserService {
	profileService := profile.InitModule(db)
	userRepository := repository.NewUserRepository(db)
	userService := NewUserService(userRepository, profileService)
	userController := NewUserController(userService)
	auth.NewAuthService(userService)
	UserRouter(userController, router)
	return userService
}
