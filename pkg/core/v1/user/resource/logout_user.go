package resource

import "github.com/doffy007/banking-api/libs/types"

func (r *UserResource) GetUserStatus(query *types.Query, email string) error {
	return r.Postgres.UserRepository.LogoutUser(query, email)
}
