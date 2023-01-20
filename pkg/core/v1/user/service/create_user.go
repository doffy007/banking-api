package service

import (
	"net/http"

	"github.com/doffy007/banking-api/libs/utils"
	"github.com/doffy007/banking-api/pkg/core/v1/user/dto"
	"github.com/doffy007/banking-api/pkg/core/v1/user/request"
	"github.com/doffy007/banking-api/pkg/model"
	"github.com/labstack/echo"
)

func (s *UserService) CreateUser(payload *dto.CreateUserDTO) error {
	userModelRes, err := s.UserResource.GetUser(nil, &model.User{
		Username: payload.Body.Username,
	})
	if err != nil {
		if err.Error() != "record not found" {
			return err
		}
	}
	if userModelRes != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "user already registered")
	}

	userModel, err := utils.ConvertToStruct[model.User](struct {
		*request.CreateUserRequestBody
	}{payload.Body})
	if err != nil {
		return err
	}

	hashedPassword, err := utils.HashPassword(userModel.Password)
	if err != nil {
		return err
	}
	userModel.Password = hashedPassword

	err = s.UserResource.CreateUser(&userModel)
	if err != nil {
		return err
	}

	return nil
}
