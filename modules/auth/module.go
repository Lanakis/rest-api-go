package auth

import (
	"gin-gorm/modules/user/entity"
	"github.com/gin-gonic/gin"
)

func InitModule(service entity.IUserService, router *gin.Engine) {

	authService := NewAuthService(service)
	RegisterHandlers(authService, router)

}
