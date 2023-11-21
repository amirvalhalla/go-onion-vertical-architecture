package user

import (
	"github.com/gin-gonic/gin"
)

type Controller interface {
	Get(ctx *gin.Context)
	GetAll(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}

type controller struct {
}

func NewController() Controller {
	return &controller{}
}

func (c *controller) Get(ctx *gin.Context) {

}

func (c *controller) GetAll(ctx *gin.Context) {

}

func (c *controller) Create(ctx *gin.Context) {

}

func (c *controller) Update(ctx *gin.Context) {

}

func (c *controller) Delete(ctx *gin.Context) {

}
