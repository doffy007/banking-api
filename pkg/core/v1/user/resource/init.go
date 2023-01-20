package resource

import (
	"github.com/doffy007/banking-api/libs/types"
	"github.com/doffy007/banking-api/pkg/model"
	"github.com/doffy007/banking-api/pkg/repository/postgres"
	"github.com/doffy007/banking-api/shared/database"
)

type InterfaceUserResource interface {
	CreateUser(payload *model.User) error
	GetUser(query *types.Query, payload *model.User) (*model.User, error)
}

type UserResource struct {
	Postgres *postgres.PostgresRepository
}

func InitUserResource() InterfaceUserResource {
	var (
		postgres = database.InitPSQL()
	)

	return &UserResource{
		Postgres: postgres,
	}
}
