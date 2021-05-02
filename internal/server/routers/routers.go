package routers

import (
	"to-do-api/internal/core/web"
	"to-do-api/internal/server/controllers"
)

func init() {
	prefix := "/api/v1"
	contollerPrefix := prefix + "/todos"

	web.RegisterRoutes(contollerPrefix, &controllers.TodoController{})
}
