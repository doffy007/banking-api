package postgres

import (
	user "github.com/doffy007/banking-api/pkg/repository/postgres/user"
	"gorm.io/gorm"
)

type PostgresRepository struct {
	UserRepository user.InterfaceUserRepository
}

func InitPostgresRepository(db *gorm.DB) *PostgresRepository {
	var (
		userRepository = user.InitUserRepository(db)
	)

	return &PostgresRepository{
		UserRepository: userRepository,
	}
}
