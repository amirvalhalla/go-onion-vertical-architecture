package router

import (
	"github.com/amirvalhalla/go-onion-vertical-architecture/feature/order"
	"github.com/amirvalhalla/go-onion-vertical-architecture/feature/product"
	"github.com/amirvalhalla/go-onion-vertical-architecture/feature/user"
	"github.com/amirvalhalla/go-onion-vertical-architecture/infrastructure/repository/sql"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	BaseAPIRoute = "/api/v1"
)

func Setup(uow sql.UnitOfWork) *gin.Engine {
	app := gin.Default()
	app.Use(cors.Default())

	apiV1 := app.Group(BaseAPIRoute)

	userRouter := apiV1.Group("/users")
	orderRouter := apiV1.Group("/orders")
	productRouter := apiV1.Group("/products")

	// user routes
	userRouter.GET("/:userId", user.NewController(uow).Get)
	userRouter.GET("", user.NewController(uow).GetAll)
	userRouter.POST("", user.NewController(uow).Create)
	userRouter.PUT("", user.NewController(uow).Update)
	userRouter.DELETE("/:userId", user.NewController(uow).Delete)

	// order routes
	orderRouter.GET("/:orderId", order.NewController(uow).Get)
	orderRouter.GET("", order.NewController(uow).GetAll)
	orderRouter.POST("", order.NewController(uow).Create)
	orderRouter.PUT("", order.NewController(uow).Update)
	orderRouter.DELETE("/:orderId", order.NewController(uow).Delete)

	// product routes
	productRouter.GET("/:productId", product.NewController(uow).Get)
	productRouter.GET("", product.NewController(uow).GetAll)
	productRouter.POST("", product.NewController(uow).Create)
	productRouter.PUT("", product.NewController(uow).Update)
	productRouter.DELETE("/:productId", product.NewController(uow).Delete)

	return app
}
