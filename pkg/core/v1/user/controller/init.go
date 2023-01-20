package controller

import (
	"github.com/doffy007/banking-api/pkg/core/v1/user/service"
	"github.com/labstack/echo"
)

type UserController struct {
	UserService service.InterfaceUserService
}

func InitUserController(userService service.InterfaceUserService) *UserController {
	return &UserController{UserService: userService}
}

func (userController *UserController) Mount(group *echo.Group) {
	group.POST("/create", userController.CreateUser)
	group.GET("/login", userController.LoginUser)
}
