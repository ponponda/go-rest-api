package web

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type Controller struct {
	Logger    *logrus.Logger
	DbContext *gorm.DB
	Context   *gin.Context
	Router    *Router
}

type IController interface {
	Get(c *Controller)
	Post(c *Controller)
	Put(c *Controller)
	Delete(c *Controller)
	Patch(c *Controller)
}

func (c *Controller) Get(ctx *Controller) {
	ctx.Context.JSON(http.StatusMethodNotAllowed, nil)
}
func (c *Controller) Post(ctx *Controller) {
	ctx.Context.JSON(http.StatusMethodNotAllowed, nil)
}
func (c *Controller) Put(ctx *Controller) {
	ctx.Context.JSON(http.StatusMethodNotAllowed, nil)
}
func (c *Controller) Delete(ctx *Controller) {
	ctx.Context.JSON(http.StatusMethodNotAllowed, nil)
}
func (c *Controller) Patch(ctx *Controller) {
	ctx.Context.JSON(http.StatusMethodNotAllowed, nil)
}
