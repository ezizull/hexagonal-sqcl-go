package routes

import (
	"github.com/gin-gonic/gin"
	"skyshi-gethired.go/infrastructure/repository/postgres/sqlc"
	"skyshi-gethired.go/infrastructure/restapi/adapter"
)

func ApplicationV1Router(router *gin.Engine, db *sqlc.Queries) {
	// Activity routers
	ActivityRoutes(router, adapter.ActivityAdapter(db))

	// Todo routers
	TodoRoutes(router, adapter.TodoAdapter(db))
}
