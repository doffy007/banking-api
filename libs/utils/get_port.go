package utils

import (
	"strconv"

	"github.com/doffy007/banking-api/config"
)

func GetPortApp() string {
	return strconv.Itoa(config.AppConfig.Server.Port)
}
