package modules

import (
	"gin-gorm/modules/auth"
	"gin-gorm/modules/user"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitModule(db *gorm.DB, router *gin.Engine) {
	userService := user.InitModule(db, router)
	auth.InitModule(userService, router)
}
