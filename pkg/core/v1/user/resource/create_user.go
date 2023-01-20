package resource

import "github.com/doffy007/banking-api/pkg/model"

func (r *UserResource) CreateUser(payload *model.User) error {
	return r.Postgres.UserRepository.CreateUser(payload)
}
