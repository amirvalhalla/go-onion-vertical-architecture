package order

import (
	"github.com/amirvalhalla/go-onion-vertical-architecture/exception"
	order "github.com/amirvalhalla/go-onion-vertical-architecture/feature/order/dto"
	mapper "github.com/amirvalhalla/go-onion-vertical-architecture/feature/order/mapper"
	product "github.com/amirvalhalla/go-onion-vertical-architecture/feature/product/dto"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/base"
	govrHTTP "github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/http"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/repository/sql"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"strconv"
)

type Controller interface {
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
	orderService Service
}

func NewController(uow sql.UnitOfWork) Controller {
	return &controller{
		orderService: NewService(uow),
	}
}

func (c *controller) Get(ctx *gin.Context) {
	errs := make([]string, 0)
	statusCode := http.StatusOK
	var id uuid.UUID

	if err := govrHTTP.BaseResponseShouldBindParamToUUID("orderId", &id, ctx); err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			base.NewResponse[product.GetDto](nil, http.StatusBadRequest, errs),
		)
		return
	}

	orderEntity, err := c.orderService.Get(ctx, id)
	if err != nil {
		statusCode = exception.DefaultOrHandleSHTTPStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode,
			base.NewResponse[product.GetDto](nil, statusCode, []string{err.Error()}),
		)
		return
	}

	ctx.JSON(
		statusCode,
		base.NewResponse[order.GetDto](mapper.ToGetDto(orderEntity), statusCode, nil),
	)
}

func (c *controller) GetAll(ctx *gin.Context) {
	errs := make([]string, 0)
	statusCode := http.StatusOK
	var search string
	var parsedSearch int
	var pageIndex int
	var pageSize int
	var err error

	if err := govrHTTP.BaseResponseShouldBindQuery("search", &search, true, ctx); err != nil {
		errs = append(errs, err.Error())
	}

	if err := govrHTTP.BaseResponseShouldBindPageIndex(&pageIndex, ctx); err != nil {
		errs = append(errs, err.Error())
	}

	if err := govrHTTP.BaseResponseShouldBindPageSize(&pageSize, ctx); err != nil {
		errs = append(errs, err.Error())
	}

	if parsedSearch, err = strconv.Atoi(search); err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			base.NewPaginationResponse[[]product.GetDto](nil, 0, 0, 0,
				http.StatusBadRequest, errs),
		)
		return
	}

	orderEntities, totalRecords, err := c.orderService.GetAll(ctx, pageIndex, pageSize, parsedSearch)
	if err != nil {
		statusCode = exception.DefaultOrHandleSHTTPStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode,
			base.NewResponse[product.GetDto](nil, statusCode, []string{err.Error()}),
		)
		return
	}

	ctx.JSON(
		statusCode,
		base.NewPaginationResponse[[]order.GetDto](mapper.ToGetDtos(orderEntities),
			pageIndex, pageSize, totalRecords, statusCode, nil),
	)
}

func (c *controller) Create(ctx *gin.Context) {
	var createDto order.CreateDto
	statusCode := http.StatusCreated

	if err := govrHTTP.BaseResponseShouldBindJSON[order.CreateDto](&createDto, ctx); err != nil {
		return
	}

	productEntity, err := c.orderService.Create(ctx, createDto)
	if err != nil {
		statusCode = exception.DefaultOrHandleSHTTPStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode,
			base.NewResponse[product.GetDto](nil, statusCode, []string{err.Error()}),
		)
		return
	}

	ctx.JSON(
		statusCode,
		base.NewResponse[order.GetDto](mapper.ToGetDto(productEntity), statusCode, nil),
	)
}

func (c *controller) Update(ctx *gin.Context) {
	var updateDto order.UpdateDto
	statusCode := http.StatusOK

	if err := govrHTTP.BaseResponseShouldBindJSON[order.UpdateDto](&updateDto, ctx); err != nil {
		return
	}

	productEntity, err := c.orderService.Update(ctx, updateDto)
	if err != nil {
		statusCode = exception.DefaultOrHandleSHTTPStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode,
			base.NewResponse[product.GetDto](nil, statusCode, []string{err.Error()}),
		)
		return
	}

	ctx.JSON(
		statusCode,
		base.NewResponse[order.GetDto](mapper.ToGetDto(productEntity), statusCode, nil),
	)
}

func (c *controller) Delete(ctx *gin.Context) {
	errs := make([]string, 0)
	statusCode := http.StatusOK
	var id uuid.UUID

	if err := govrHTTP.BaseResponseShouldBindParamToUUID("orderId", &id, ctx); err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			base.NewPaginationResponse[[]product.GetDto](nil, 0, 0, 0,
				http.StatusBadRequest, errs),
		)
		return
	}

	err := c.orderService.Delete(ctx, id)
	if err != nil {
		statusCode = exception.DefaultOrHandleSHTTPStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode,
			base.NewResponse[product.GetDto](nil, statusCode, []string{err.Error()}),
		)
		return
	}

	ctx.JSON(
		statusCode,
		base.NewResponse[struct{}](nil, statusCode, nil),
	)
}
