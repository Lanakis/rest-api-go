package auth

import (
	"gin-gorm/modules/auth/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterHandlers(authService *Service, router *gin.Engine) {
	authController := NewAuthController(authService)

	router.POST("/auth", authController.Login)
}

type Controller struct {
	AuthService *Service
}

func NewAuthController(authService *Service) *Controller {
	return &Controller{
		AuthService: authService,
	}
}

func (c *Controller) Login(ctx *gin.Context) {
	var data dto.SignAuthDto
	if err := ctx.ShouldBindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	token, err := c.AuthService.SignIn(ctx.Request.Context(), data)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"token": token})
}

func (c *Controller) Profile(ctx *gin.Context) {
	user := ctx.MustGet("user").(*User)
	ctx.JSON(http.StatusOK, user)
}
