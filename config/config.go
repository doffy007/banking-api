package config

import (
	"errors"
	"flag"
	"fmt"

	"github.com/doffy007/banking-api/libs/constant"
	"github.com/doffy007/banking-api/libs/types"
	"github.com/doffy007/banking-api/pkg/model"
	"github.com/labstack/echo/middleware"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

var (
	AppConfig *DefaultConfig
)

func ConfigApps() {
	var (
		environtment = flag.String("env", "", "input the environment type")
	)

	flag.Parse()

	switch constant.EnvironmentType(*environtment) {
	case constant.DEV:
		viper.SetConfigFile("./resource/dev.app.yaml")
	case constant.STAG:
		viper.SetConfigFile("./resource/stag.app.yaml")
	case constant.PROD:
		viper.SetConfigFile("./resource/prod.app.yaml")
	case constant.TEST:
		viper.SetConfigFile("./resource/test.app.yaml")
	default:
		panic(errors.New("missing input environtment type [ dev | stag | prod | test]"))
	}

	if err := viper.ReadInConfig(); err != nil {
		logrus.Panic(err)
	}

	var conf DefaultConfig
	if err := viper.Unmarshal(&conf); err != nil {
		logrus.Panic(err)
	}

	AppConfig = &conf
}

func Postgres(ds Datasource, dialect constant.DialectDatabaseType) *gorm.DB {
	var (
		db  *gorm.DB
		err error
	)

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=require search_path=%s",
		ds.Url,
		ds.Username,
		ds.Password,
		ds.DatabaseName,
		ds.Port,
		ds.Schema,
	)
	cfg := &gorm.Config{
		Logger: logger.Default,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   ds.Schema,
			SingularTable: false,
		},
	}

	if ds.DebugMode {
		cfg.Logger = logger.Default.LogMode(logger.Info)
	}

	switch dialect {
	case constant.POSTGRES:
		db, err = gorm.Open(postgres.Open(dsn), cfg)
		if err != nil {
			logrus.Panic(err)
		}
	}

	// Auto Migration Models
	db.AutoMigrate(
		&model.User{},
		&model.UserBalance{},
		&model.UserBalanceHistory{},
		&model.BankBalace{},
		&model.BankBalaceHistory{},
	)

	sqlDB, err := db.DB()
	if err != nil {
		logrus.Panic(err)
	}

	sqlDB.SetConnMaxIdleTime(ds.ConnectionTimeout)
	sqlDB.SetMaxIdleConns(ds.MaxIdleConnection)
	sqlDB.SetMaxOpenConns(ds.MaxOpenConnection)

	return db
}

func ConfigJWT() middleware.JWTConfig {
	return middleware.JWTConfig{
		Claims:     &types.JWTCustomClaims{},
		SigningKey: []byte(AppConfig.JWT.SecretKey),
	}
}
