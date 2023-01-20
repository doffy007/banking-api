package user

import (
	"github.com/doffy007/banking-api/libs/types"
	"github.com/doffy007/banking-api/pkg/model"
	"gorm.io/gorm"
)

type InterfaceUserRepository interface {
	CreateUser(payload *model.User) error
	LoginUser(query *types.Query, payload *model.User) (*model.User, error)
	LogoutUser(query *types.Query, email string) error
}

type UserRepository struct {
	db *gorm.DB
}

func InitUserRepository(db *gorm.DB) InterfaceUserRepository {
	return &UserRepository{
		db: db,
	}
}
