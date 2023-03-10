package config

import "time"

type DefaultConfig struct {
	Apps            Apps            `mapstructure:"apps"`
	Server          Server          `mapstructure:"server"`
	Database        Database        `mapstructure:"database"`
	JWT             JWT             `mapstructure:"jwt"`
	PasswordHashing PasswordHashing `mapstructure:"passwordHashing"`
}

type Apps struct {
	Name    string `mapstructure:"name"`
	Version string `mapstructure:"version"`
}

type Server struct {
	Port int `mapstructure:"port"`
}

type Datasource struct {
	Url               string        `mapstructure:"url"`
	Port              string        `mapstructure:"port"`
	DatabaseName      string        `mapstructure:"databaseName"`
	Username          string        `mapstructure:"username"`
	Password          string        `mapstructure:"password"`
	Schema            string        `mapstructure:"schema"`
	ConnectionTimeout time.Duration `mapstructure:"connectionTimeout"`
	MaxIdleConnection int           `mapstructure:"maxIdleConnection"`
	MaxOpenConnection int           `mapstructure:"maxOpenConnection"`
	DebugMode         bool          `mapstructure:"debugMode"`
}

type Database struct {
	Postgres Datasource `mapstructure:"postgres"`
}

type JWT struct {
	SecretKey string `mapstructure:"secretKey"`
}

type PasswordHashing struct {
	HashSalt int `mapstructure:"hashSalt"`
}
