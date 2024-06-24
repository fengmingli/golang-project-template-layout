package controller

import (
	"net/http"
	"strconv"

	"github.com/fengmingli/golang-project-template-layout/internal/domain/entity"
	"github.com/fengmingli/golang-project-template-layout/internal/usecase"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	createUserUseCase *usecase.CreateUserUseCase
	getUserUseCase    *usecase.GetUserUseCase
	listUsersUseCase  *usecase.ListUsersUseCase
	updateUserUseCase *usecase.UpdateUserUseCase
}

func NewUserController(
	createUserUseCase *usecase.CreateUserUseCase,
	getUserUseCase *usecase.GetUserUseCase,
	listUsersUseCase *usecase.ListUsersUseCase,
	updateUserUseCase *usecase.UpdateUserUseCase,
) *UserController {
	return &UserController{
		createUserUseCase: createUserUseCase,
		getUserUseCase:    getUserUseCase,
		listUsersUseCase:  listUsersUseCase,
		updateUserUseCase: updateUserUseCase,
	}
}

func (c *UserController) CreateUser(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.createUserUseCase.Execute(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, user)
}

func (c *UserController) GetUser(ctx *gin.Context) {
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)
	user, err := c.getUserUseCase.Execute(id)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (c *UserController) ListUsers(ctx *gin.Context) {
	users, err := c.listUsersUseCase.Execute()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (c *UserController) UpdateUser(ctx *gin.Context) {
	var user entity.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := c.updateUserUseCase.Execute(&user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
