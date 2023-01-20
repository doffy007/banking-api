package database

import (
	"github.com/doffy007/banking-api/config"
	"github.com/doffy007/banking-api/libs/constant"
	"github.com/doffy007/banking-api/pkg/repository/postgres"
)

func InitPSQL() *postgres.PostgresRepository {
	db := config.Postgres(config.AppConfig.Database.Postgres, constant.POSTGRES)
	return postgres.InitPostgresRepository(db)
}
