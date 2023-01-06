package middleware

import (
	"chatty/pkg/http/jwt"
	"net/http"

	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/labstack/echo/v4"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"

	"chatty/chatty/delivery"
	"chatty/chatty/usecase"
)

//ErrorHandlerMiddleware handle responses which ended with not nil error
var ErrorHandlerMiddleware = func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(eCtx echo.Context) error {
		if err := next(eCtx); err != nil {
			return getHttpErr(err)
		}
		return nil
	}
}

//errResponse log error and return new http error with custom status code and message
func getHttpErr(err error) error {
	log.Error().Msg(errors.Wrap(err, "Err in ServerHandler: ").Error())

	return echo.NewHTTPError(errToHttpErr(err))
}

//getErrStatusCode returns http status code based on err and appropriate error message
func errToHttpErr(err error) (int, interface{}) {
	//delivery errors
	if errors.Is(err, delivery.ErrBadRequestBody) {
		return http.StatusBadRequest, delivery.ErrBadRequestBody.Error()
	}
	if errors.Is(err, delivery.ErrBadContentType) {
		return http.StatusUnsupportedMediaType, delivery.ErrBadContentType.Error()
	}
	if errors.Is(err, echo.ErrNotFound) {
		return http.StatusNotFound, err.Error()
	}

	//JWT errors
	if errors.Is(err, jwt.ErrUnableGenerateToken) {
		return http.StatusUnauthorized, jwt.ErrUnableGenerateToken.Error()
	}

	//validation errors
	if _, ok := err.(validation.Errors); ok {
		return http.StatusBadRequest, err
	}

	//usecase errors
	if errors.Is(err, usecase.ErrNoUser) {
		return http.StatusNotFound, usecase.ErrNoUser.Error()
	}
	if errors.Is(err, usecase.ErrUserDuplication) {
		return http.StatusConflict, usecase.ErrUserDuplication.Error()
	}
	if errors.Is(err, usecase.ErrDataDuplication) {
		err = errors.Unwrap(err)
		return http.StatusConflict, usecase.ErrDataDuplication.Error()
	}
	if errors.Is(err, usecase.ErrUserUnauthorized) {
		return http.StatusUnauthorized, usecase.ErrUserUnauthorized.Error()
	}

	return 0, usecase.ErrInternalError
}
