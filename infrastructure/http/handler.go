package http

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/amirvalhalla/go-onion-vertical-architecture/config"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/base"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/util"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

var ErrCouldNotBindJSON = errors.New("could not bind json")

func BaseResponseShouldBindJSON[T any](model *T, ctx *gin.Context) error {
	if err := ctx.ShouldBindJSON(model); err != nil {
		var ve validator.ValidationErrors
		errs := make([]string, 0)
		if errors.As(err, &ve) {
			for _, fe := range ve {
				errs = append(errs, fe.Field()+" "+util.GetErrorMessage(fe))
			}
		}

		if len(errs) == 0 {
			errs = append(errs, err.Error())
		}

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			base.NewResponse[T](nil, http.StatusBadRequest, errs),
		)
		return ErrCouldNotBindJSON
	}
	return nil
}

func BasePaginationResponseShouldBindJSON[T any](model *T, ctx *gin.Context) error {
	if err := ctx.ShouldBindJSON(model); err != nil {
		var ve validator.ValidationErrors
		errs := make([]string, 0)
		if errors.As(err, &ve) {
			for _, fe := range ve {
				errs = append(errs, fe.Field()+" "+util.GetErrorMessage(fe))
			}
		}

		if len(errs) == 0 {
			errs = append(errs, err.Error())
		}

		ctx.AbortWithStatusJSON(
			http.StatusBadRequest,
			base.NewPaginationResponse[T](nil, 0, 0, 0,
				http.StatusBadRequest, errs),
		)
		return ErrCouldNotBindJSON
	}
	return nil
}

func BaseResponseShouldBindParamToUUID(paramName string, model *uuid.UUID, ctx *gin.Context) error {
	v := ctx.Param(paramName)

	id, err := uuid.Parse(v)
	if err != nil {
		return fmt.Errorf("%s param has invalid uuid", paramName)
	}

	*model = id
	return nil
}

func BaseResponseShouldBindParam(paramName string, model *string, ctx *gin.Context) error {
	v := ctx.Param(paramName)

	*model = v
	return nil
}

func BaseResponseShouldBindQuery(queryParamName string, model *string, canEmpty bool, ctx *gin.Context) error {
	v, ok := ctx.GetQuery(queryParamName)

	if !ok && !canEmpty {
		return fmt.Errorf("%s query param is required", queryParamName)
	}

	*model = v
	return nil
}

func BaseResponseShouldBindQueryToInt(queryParamName string, model *int, canEmpty bool, ctx *gin.Context) error {
	v, ok := ctx.GetQuery(queryParamName)

	if !ok && !canEmpty {
		return fmt.Errorf("%s query param is required", queryParamName)
	}

	if canEmpty && util.IsStringEmpty(v) {
		*model = -1
		return nil
	}

	c, err := strconv.Atoi(v)
	if err != nil {
		return fmt.Errorf("%s query param has invalid type", queryParamName)
	}

	if queryParamName == "pageIndex" && c < config.MinPageIndex {
		return fmt.Errorf("%s query param should be greater or equal than %d",
			queryParamName, config.MinPageIndex)
	}

	if queryParamName == "pageSize" && c > config.MaxPageSize {
		return fmt.Errorf("%s query param should be less or equal than %d",
			queryParamName, config.MaxPageSize)
	}

	*model = c
	return nil
}

func BaseResponseShouldBindPageIndex(model *int, ctx *gin.Context) error {
	queryParamName := "pageIndex"
	v, ok := ctx.GetQuery(queryParamName)

	if !ok {
		return fmt.Errorf("%s query param is required", queryParamName)
	}

	c, err := strconv.Atoi(v)
	if err != nil {
		return fmt.Errorf("%s query param has invalid type", queryParamName)
	}

	if c < config.MinPageIndex {
		return fmt.Errorf("%s query param should be greater or equal than %d",
			queryParamName, config.MinPageIndex)
	}

	*model = c
	return nil
}

func BaseResponseShouldBindPageSize(model *int, ctx *gin.Context) error {
	queryParamName := "pageSize"
	v, ok := ctx.GetQuery(queryParamName)

	if !ok {
		return fmt.Errorf("%s query param is required", queryParamName)
	}

	c, err := strconv.Atoi(v)
	if err != nil {
		return fmt.Errorf("%s query param has invalid type", queryParamName)
	}

	if c > config.MaxPageSize {
		return fmt.Errorf("%s query param should be less or equal than %d",
			queryParamName, config.MaxPageSize)
	}

	*model = c
	return nil
}
