package routes

import (
	"hexagonal-sqlc/infrastructure/repository/postgres/sqlc"
	"hexagonal-sqlc/infrastructure/restapi/adapter"

	"github.com/gin-gonic/gin"
)

func ApplicationV1Router(router *gin.Engine, db *sqlc.Queries) {
	// Activity routers
	ActivityRoutes(router, adapter.ActivityAdapter(db))

	// Todo routers
	TodoRoutes(router, adapter.TodoAdapter(db))
}
