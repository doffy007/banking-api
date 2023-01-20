//go:build wireinject
// +build wireinject

package wire

import (
	"github.com/doffy007/banking-api/pkg/core/v1/user/controller"
	"github.com/doffy007/banking-api/pkg/core/v1/user/resource"
	"github.com/doffy007/banking-api/pkg/core/v1/user/service"
	"github.com/google/wire"
)

func UserControllerV1() *controller.UserController {
	wire.Build(
		controller.InitUserController,
		service.InitUserService,
		resource.InitUserResource,
	)
	return &controller.UserController{}
}
