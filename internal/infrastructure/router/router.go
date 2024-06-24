package router

import (
	"github.com/fengmingli/golang-project-template-layout/internal/adapter/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(userController *controller.UserController) *gin.Engine {
	r := gin.Default()
	r.POST("/users", userController.CreateUser)
	r.GET("/users/:id", userController.GetUser)
	r.GET("/users", userController.ListUsers)
	r.PUT("/users/:id", userController.UpdateUser)
	return r
}
