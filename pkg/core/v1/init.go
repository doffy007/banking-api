package v1

import (
	"github.com/doffy007/banking-api/wire"
	"github.com/labstack/echo"
)

func InitRouter(e *echo.Echo) {
	v1 := e.Group("/v1")
	{
		v1User := v1.Group("/users")
		v1UserController := wire.UserControllerV1()
		v1UserController.Mount(v1User)
	}
}
