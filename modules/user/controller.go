package user

import (
	"gin-gorm/modules/user/dto"
	"gin-gorm/modules/user/entity"
	"gin-gorm/utils/filter"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller struct {
	service entity.IUserService
}

func NewUserController(service entity.IUserService) *Controller {
	return &Controller{service: service}
}

func (controller *Controller) createUser(ctx *gin.Context) {

	var userCreateRequest dto.Create
	if err := ctx.ShouldBindJSON(&userCreateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode JSON"})
		return
	}

	if _, err := controller.service.Create(ctx.Request.Context(), userCreateRequest); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "User created successfully", "data": userCreateRequest})
}

func (controller *Controller) findAllUsers(ctx *gin.Context) {

	var filterOptions filter.Option
	usersResponse, err := controller.service.FindAll(ctx.Request.Context(), filterOptions)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}
	response := dto.Response{
		Users: usersResponse,
		Count: len(usersResponse),
	}
	ctx.JSON(http.StatusOK, response)

}

func (controller *Controller) findOneUser(ctx *gin.Context) {

	id := ctx.Param("userId")
	userID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	user, err := controller.service.FindOne(ctx.Request.Context(), userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to find user"})
		return
	}

	ctx.JSON(http.StatusOK, user)

}

func (controller *Controller) updateUser(ctx *gin.Context) {

	id := ctx.Param("userId")
	userID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	var userUpdateRequest dto.Update
	if err := ctx.ShouldBindJSON(&userUpdateRequest); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to decode JSON"})
		return
	}

	controller.service.Update(ctx.Request.Context(), userUpdateRequest, userID)

	ctx.JSON(http.StatusOK, gin.H{"message": "User updated successfully", "data": userUpdateRequest})

}

func (controller *Controller) deleteUser(ctx *gin.Context) {

	id := ctx.Param("userId")
	userID, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID"})
		return
	}

	controller.service.Delete(ctx.Request.Context(), userID)

	ctx.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})

}
