package usecase

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ErrorResponse struct {
	Message string `json:"message"`
	Code    string `json:"code"`
	Err     error  `json:"-"`
}

func (e *ErrorResponse) Error() string {
	return fmt.Sprintf("%s %s: %v", e.Code, e.Message, e.Err)
}

func (e *ErrorResponse) Unwrap() error {
	return e.Err
}

func newErrorResponse(c echo.Context, httpStatusCode int, errorCode string, errorMessage string, err error) error {
	return echo.NewHTTPError(httpStatusCode, &ErrorResponse{
		Code:    errorCode,
		Message: errorMessage,
		Err:     err,
	})
}

func badRequest(c echo.Context, message string, err error) error {
	return newErrorResponse(c, http.StatusBadRequest, "BadRequest", message, err)
}

func internalServerError(c echo.Context, message string, err error) error {
	return newErrorResponse(c, http.StatusInternalServerError, "InternalServerError", message, err)
}
