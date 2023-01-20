package resource

import (
	"github.com/doffy007/banking-api/libs/types"
	"github.com/doffy007/banking-api/pkg/model"
)

func (r *UserResource) GetUser(query *types.Query, payload *model.User) (*model.User, error) {
	return r.Postgres.UserRepository.LoginUser(query, payload)
}
