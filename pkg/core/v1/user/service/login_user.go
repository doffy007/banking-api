package service

import (
	"net/http"
	"time"

	"github.com/doffy007/banking-api/config"
	"github.com/doffy007/banking-api/libs/types"
	"github.com/doffy007/banking-api/libs/utils"
	"github.com/doffy007/banking-api/pkg/core/v1/user/dto"
	"github.com/doffy007/banking-api/pkg/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

func (s *UserService) GetUser(payload *dto.GetUserDTO) (*dto.UserWithTokenDTO, error) {
	userModelRes, err := s.UserResource.GetUser(nil, &model.User{
		Username: payload.Body.Username,
	})
	if err != nil {
		return nil, err
	}

	match := utils.CheckPasswordHash(payload.Body.Password, userModelRes.Password)
	if !match {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "password not match")
	}

	claims := &types.JWTCustomClaims{
		ID:       userModelRes.ID,
		Username: userModelRes.Username,
		Email:    userModelRes.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(config.AppConfig.JWT.SecretKey))
	if err != nil {
		return nil, err
	}

	userWithTokenDto, err := utils.ConvertToStruct[*dto.UserWithTokenDTO](&dto.UserWithTokenDTO{
		User: dto.UserDTOForLogin{
			ID: userModelRes.ID,
		},
		Token: t,
	})
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	return userWithTokenDto, nil
}
