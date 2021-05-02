package web

import (
	"github.com/gin-gonic/gin"
)

type Router struct {
}

var engine *gin.Engine

func init() {
	engine = gin.Default()

	// defer engine.Run(":" + port)
}

func createGinFunc(callback func(*Controller)) gin.HandlerFunc {
	return func(c *gin.Context) {
		if Application.Context == nil {
			Application.Context = c
		}
		callback(&Application)
	}
}

func (r *Router) AddRoute(path string, ctrlr IController) {
	idPath := path + "/:id"
	engine.GET(path+"/*id", createGinFunc(ctrlr.Get))
	engine.POST(path, createGinFunc(ctrlr.Post))
	engine.PUT(idPath, createGinFunc(ctrlr.Put))
	engine.PATCH(idPath, createGinFunc(ctrlr.Patch))
	engine.DELETE(idPath, createGinFunc(ctrlr.Delete))
}

func (r *Router) Run(port string) {
	engine.Run(":" + port)
}
