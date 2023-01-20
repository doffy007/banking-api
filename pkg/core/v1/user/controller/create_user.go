package controller

import (
	"net/http"

	"github.com/doffy007/banking-api/libs/dto"
	"github.com/doffy007/banking-api/libs/response"
	"github.com/doffy007/banking-api/libs/validator"
	request_dto "github.com/doffy007/banking-api/pkg/core/v1/user/dto"
	"github.com/doffy007/banking-api/pkg/core/v1/user/request"
	"github.com/labstack/echo"
)

func (c *UserController) CreateUser(ctx echo.Context) error {
	reqBody := new(request.CreateUserRequestBody)
	if err := validator.ValidateRequest(ctx, reqBody); err != nil {
		return err
	}

	if err := c.UserService.CreateUser(&request_dto.CreateUserDTO{
		DTO: dto.DTO[any, any, *request.CreateUserRequestBody]{
			Param: nil,
			Query: nil,
			Body:  reqBody,
		},
	}); err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.ResponseSuccess[any]("success create user", nil))
}
