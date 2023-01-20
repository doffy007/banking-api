package dto

import (
	"github.com/doffy007/banking-api/libs/dto"
	"github.com/doffy007/banking-api/pkg/core/v1/user/request"
)

type LogoutUserDto struct {
	dto.DTO[any, any, *request.LogoutUserRequestBody]
}
