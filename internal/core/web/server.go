package web

import (
	"os"
	"to-do-api/internal/core/dataContext"
	"to-do-api/internal/core/logger"

	"github.com/joho/godotenv"
)

var Application Controller

func init() {
	godotenv.Load()
	Application = Controller{
		Logger: logger.Default(),
		DbContext: dataContext.Init(
			os.Getenv("DB_HOST"),
			os.Getenv("DB_PORT"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
		),
		// Context: &gin.Context{},
		Router: &Router{},
	}
}

func RegisterRoutes(path string, c IController) {
	Application.Router.AddRoute(path, c)
}

func Run() {
	Application.Router.Run(os.Getenv("API_PORT"))
}
