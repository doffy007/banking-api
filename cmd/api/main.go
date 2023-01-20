package main

import (
	"fmt"

	"github.com/doffy007/banking-api/config"
	"github.com/doffy007/banking-api/libs/filter"
	"github.com/doffy007/banking-api/libs/utils"
	val "github.com/doffy007/banking-api/libs/validator"
	v1 "github.com/doffy007/banking-api/pkg/core/v1"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func init() {
	config.ConfigApps()
}

func main() {
	e := echo.New()
	e.Validator = &val.CustomValidator{Validator: validator.New()}
	e.HTTPErrorHandler = filter.CustomHTTPErrorHandler
	e.Use(
		middleware.Logger(),
		middleware.Recover(),
		middleware.RequestID(),
	)

	v1.InitRouter(e)

	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", utils.GetPortApp())))
}
