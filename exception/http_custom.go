package exception

import (
	"errors"
	"net/http"

	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/repository/sql"
)

func DefaultOrHandleSHTTPStatusCode(err error, statusCode ...int) int {
	switch {
	case errors.Is(err, sql.ErrRecordNotFound):
		return http.StatusNotFound
	case errors.Is(err, sql.ErrFindOne), errors.Is(err, sql.ErrFindAll), errors.Is(err, sql.ErrCouldNotInsert),
		errors.Is(err, sql.ErrCouldNotBatchInsert), errors.Is(err, sql.ErrCouldNotUpdate),
		errors.Is(err, sql.ErrCouldNotSave), errors.Is(err, sql.ErrCouldNotDelete):
		return http.StatusInternalServerError
	default:
		if len(statusCode) < 1 {
			return http.StatusInternalServerError
		}
		return statusCode[0]
	}
}
