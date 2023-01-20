package filter

import (
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"

	"github.com/doffy007/banking-api/libs/response"
	"github.com/go-playground/validator"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func CustomHTTPErrorHandler(err error, ctx echo.Context) {
	logrus.Infof("error type %T\n", err)
	he, ok := err.(*echo.HTTPError)
	if ok {
		ctx.JSON(he.Code, response.ResponseFail(fmt.Sprint(he.Message), he.Internal))
		return
	}
	ve, ok := err.(validator.ValidationErrors)
	if ok {
		ctx.JSON(http.StatusBadRequest, response.ResponseFail(ve.Error(), ve))
		return
	}
	me, ok := err.(*json.MarshalerError)
	if ok {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResponseFail(me.Error(), me.Unwrap()))
		return
	}
	re, ok := err.(runtime.Error)
	if ok {
		ctx.JSON(http.StatusUnprocessableEntity, response.ResponseFail(re.Error(), re))
		return
	}
	ctx.JSON(http.StatusInternalServerError, response.ResponseFail(err.Error(), err))
}
