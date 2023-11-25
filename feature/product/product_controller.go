package product

import (
	"net/http"

	product "github.com/amirvalhalla/go-onion-vertical-architecture/feature/product/dto"

	"github.com/amirvalhalla/go-onion-vertical-architecture/exception"
	mapper "github.com/amirvalhalla/go-onion-vertical-architecture/feature/product/mapper"
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
	productService Service
}

func NewController(uow sql.UnitOfWork) Controller {
	return &controller{
		productService: NewService(uow),
	}
}

func (c *controller) Get(ctx *gin.Context) {
	errs := make([]string, 0)
	statusCode := http.StatusOK
	var id uuid.UUID

	if err := govrHTTP.BaseResponseShouldBindParamToUUID("productId", &id, ctx); err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			base.NewResponse[product.GetDto](nil, http.StatusBadRequest, errs),
		)
		return
	}

	productEntity, err := c.productService.Get(ctx, id)
	if err != nil {
		statusCode = exception.DefaultOrHandleSHTTPStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode,
			base.NewResponse[product.GetDto](nil, statusCode, []string{err.Error()}),
		)
		return
	}

	ctx.JSON(
		statusCode,
		base.NewResponse[product.GetDto](mapper.ToGetDto(productEntity), statusCode, nil),
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
			base.NewPaginationResponse[[]product.GetDto](nil, 0, 0, 0,
				http.StatusBadRequest, errs),
		)
		return
	}

	productEntities, totalRecords, err := c.productService.GetAll(ctx, pageIndex, pageSize, search)
	if err != nil {
		statusCode = exception.DefaultOrHandleSHTTPStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode,
			base.NewResponse[product.GetDto](nil, statusCode, []string{err.Error()}),
		)
		return
	}

	ctx.JSON(
		statusCode,
		base.NewPaginationResponse[[]product.GetDto](mapper.ToGetDtos(productEntities),
			pageIndex, pageSize, totalRecords, statusCode, nil),
	)
}

func (c *controller) Create(ctx *gin.Context) {
	var createDto product.CreateDto
	statusCode := http.StatusCreated

	if err := govrHTTP.BaseResponseShouldBindJSON[product.CreateDto](&createDto, ctx); err != nil {
		return
	}

	productEntity, err := c.productService.Create(ctx, createDto)
	if err != nil {
		statusCode = exception.DefaultOrHandleSHTTPStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode,
			base.NewResponse[product.GetDto](nil, statusCode, []string{err.Error()}),
		)
		return
	}

	ctx.JSON(
		statusCode,
		base.NewResponse[product.GetDto](mapper.ToGetDto(productEntity), statusCode, nil),
	)
}

func (c *controller) Update(ctx *gin.Context) {
	var updateDto product.UpdateDto
	statusCode := http.StatusOK

	if err := govrHTTP.BaseResponseShouldBindJSON[product.UpdateDto](&updateDto, ctx); err != nil {
		return
	}

	productEntity, err := c.productService.Update(ctx, updateDto)
	if err != nil {
		statusCode = exception.DefaultOrHandleSHTTPStatusCode(err)
		ctx.AbortWithStatusJSON(statusCode,
			base.NewResponse[product.GetDto](nil, statusCode, []string{err.Error()}),
		)
		return
	}

	ctx.JSON(
		statusCode,
		base.NewResponse[product.GetDto](mapper.ToGetDto(productEntity), statusCode, nil),
	)
}

func (c *controller) Delete(ctx *gin.Context) {
	errs := make([]string, 0)
	statusCode := http.StatusOK
	var id uuid.UUID

	if err := govrHTTP.BaseResponseShouldBindParamToUUID("productId", &id, ctx); err != nil {
		errs = append(errs, err.Error())
	}

	if len(errs) > 0 {
		ctx.AbortWithStatusJSON(http.StatusBadRequest,
			base.NewPaginationResponse[[]product.GetDto](nil, 0, 0, 0,
				http.StatusBadRequest, errs),
		)
		return
	}

	err := c.productService.Delete(ctx, id)
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
