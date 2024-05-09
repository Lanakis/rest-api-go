package user

import (
	"github.com/gin-gonic/gin"
)

func UserRouter(userController *Controller, router *gin.Engine) gin.IRouter {

	userRouter := router.Group("/user")
	userRouter.GET("", userController.findAllUsers)
	userRouter.GET("/:userId", userController.findOneUser)
	userRouter.POST("", userController.createUser)
	userRouter.PATCH("/:userId", userController.updateUser)
	userRouter.DELETE("/:userId", userController.deleteUser)

	return userRouter
}
