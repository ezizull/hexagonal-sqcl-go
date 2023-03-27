package routes

import (
	"github.com/gin-gonic/gin"
	"skyshi-gethired.go/infrastructure/repository/postgres/sqlc"
	"skyshi-gethired.go/infrastructure/restapi/adapter"
)

func ApplicationV1Router(router *gin.Engine, db *sqlc.Queries) {
	routerV1 := router.Group("/v1")

	{
		TodoRoutes(routerV1, adapter.TodoAdapter(db))
	}
}
