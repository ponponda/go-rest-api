package server

import (
	"to-do-api/internal/core/web"
	"to-do-api/internal/db"
	_ "to-do-api/internal/server/routers"
)

func init() {
	web.Run()
	db.Migrate(web.Application.DbContext)
}
