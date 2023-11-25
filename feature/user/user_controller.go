package user

import (
	"net/http"

	"github.com/amirvalhalla/go-onion-vertical-architecture/feature/user/dto"
	mapper "github.com/amirvalhalla/go-onion-vertical-architecture/feature/user/mapper"

	"github.com/amirvalhalla/go-onion-vertical-architecture/exception"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/base"
	govrHTTP "github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/http"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/repository/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Controller interface {
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	userService Service
}

func NewController(uow sql.UnitOfWork) Controller {
	return &controller{
		userService: NewService(uow),
	}
}

func (c *controller) Get(ctx *gin.Context) {
	errs := make([]string, 0)
	statusCode := http.StatusOK
	var id uuid.UUID

	if err := govrHTTP.BaseResponseShouldBindParamToUUID("userId", &id, ctx); err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			base.NewResponse[user.GetDto](nil, http.StatusBadRequest, errs),
		)
		return
	}

	userEntity, err := c.userService.Get(ctx, id)
	if err != nil {
		statusCode = exception.DefaultOrHandleSHTTPStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode,
			base.NewResponse[user.GetDto](nil, statusCode, []string{err.Error()}),
		)
		return
	}

	ctx.JSON(
		statusCode,
		base.NewResponse[user.GetDto](mapper.ToGetDto(userEntity), statusCode, nil),
	)
}

func (c *controller) GetAll(ctx *gin.Context) {
	errs := make([]string, 0)
	statusCode := http.StatusOK
	var search string
	var pageIndex int
	var pageSize int

	if err := govrHTTP.BaseResponseShouldBindQuery("search", &search, true, ctx); err != nil {
		errs = append(errs, err.Error())
	}

	if err := govrHTTP.BaseResponseShouldBindPageIndex(&pageIndex, ctx); err != nil {
		errs = append(errs, err.Error())
	}

	if err := govrHTTP.BaseResponseShouldBindPageSize(&pageSize, ctx); err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			base.NewPaginationResponse[[]user.GetDto](nil, 0, 0, 0,
				http.StatusBadRequest, errs),
		)
		return
	}

	userEntities, totalRecords, err := c.userService.GetAll(ctx, pageIndex, pageSize, search)
	if err != nil {
		statusCode = exception.DefaultOrHandleSHTTPStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode,
			base.NewResponse[user.GetDto](nil, statusCode, []string{err.Error()}),
		)
		return
	}

	ctx.JSON(
		statusCode,
		base.NewPaginationResponse[[]user.GetDto](mapper.ToGetDtos(userEntities),
			pageIndex, pageSize, totalRecords, statusCode, nil),
	)
}

func (c *controller) Create(ctx *gin.Context) {
	var createDto user.CreateDto
	statusCode := http.StatusCreated

	if err := govrHTTP.BaseResponseShouldBindJSON[user.CreateDto](&createDto, ctx); err != nil {
		return
	}

	userEntity, err := c.userService.Create(ctx, createDto)
	if err != nil {
		statusCode = exception.DefaultOrHandleSHTTPStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode,
			base.NewResponse[user.GetDto](nil, statusCode, []string{err.Error()}),
		)
		return
	}

	ctx.JSON(
		statusCode,
		base.NewResponse[user.GetDto](mapper.ToGetDto(userEntity), statusCode, nil),
	)
}

func (c *controller) Update(ctx *gin.Context) {
	var updateDto user.UpdateDto
	statusCode := http.StatusOK

	if err := govrHTTP.BaseResponseShouldBindJSON[user.UpdateDto](&updateDto, ctx); err != nil {
		return
	}

	userEntity, err := c.userService.Update(ctx, updateDto)
	if err != nil {
		statusCode = exception.DefaultOrHandleSHTTPStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode,
			base.NewResponse[user.GetDto](nil, statusCode, []string{err.Error()}),
		)
		return
	}

	ctx.JSON(
		statusCode,
		base.NewResponse[user.GetDto](mapper.ToGetDto(userEntity), statusCode, nil),
	)
}

func (c *controller) Delete(ctx *gin.Context) {
	errs := make([]string, 0)
	statusCode := http.StatusOK
	var id uuid.UUID

	if err := govrHTTP.BaseResponseShouldBindParamToUUID("userId", &id, ctx); err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			base.NewPaginationResponse[[]user.GetDto](nil, 0, 0, 0,
				http.StatusBadRequest, errs),
		)
		return
	}

	err := c.userService.Delete(ctx, id)
	if err != nil {
		statusCode = exception.DefaultOrHandleSHTTPStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode,
			base.NewResponse[user.GetDto](nil, statusCode, []string{err.Error()}),
		)
		return
	}

	ctx.JSON(
		statusCode,
		base.NewResponse[struct{}](nil, statusCode, nil),
	)
}
