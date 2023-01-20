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

func (c *UserController) LoginUser(ctx echo.Context) error {
	reqBody := new(request.LoginUserRequestBody)
	if err := validator.ValidateRequest(ctx, reqBody); err != nil {
		return err
	}

	userWithToken, err := c.UserService.GetUser(&request_dto.GetUserDTO{
		DTO: dto.DTO[any, any, *request.LoginUserRequestBody]{
			Param: nil,
			Query: nil,
			Body:  reqBody,
		},
	})
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, response.ResponseSuccess("login user success", userWithToken))
}
