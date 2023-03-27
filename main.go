package main

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	limit "github.com/aviddiviner/gin-limit"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"skyshi-gethired.go/infrastructure/repository/postgres"
	"skyshi-gethired.go/infrastructure/restapi/middlewares"
	"skyshi-gethired.go/infrastructure/restapi/routes"

	errorsController "skyshi-gethired.go/infrastructure/restapi/controllers/errors"
)

func main() {
	router := gin.Default()
	router.Use(limit.MaxAllowed(200))
	router.Use(cors.Default())

	// postgres connection
	postgresDB, err := postgres.NewSqlc()
	if err != nil {
		_ = fmt.Errorf("Cannot connect to PostgresDB:: %s", err)
		panic(err)
	}

	router.Use(middlewares.GinBodyLogMiddleware)
	router.Use(errorsController.Handler)

	// postgres routes
	routes.ApplicationV1Router(router, postgresDB)

	startServer(router)
}

func startServer(router http.Handler) {
	viper.SetConfigFile("config.json")
	if err := viper.ReadInConfig(); err != nil {
		_ = fmt.Errorf("fatal error in config file: %s", err.Error())
		panic(err)

	}
	serverPort := fmt.Sprintf(":%s", viper.GetString("ServerPort"))
	s := &http.Server{
		Addr:           serverPort,
		Handler:        router,
		ReadTimeout:    18000 * time.Second,
		WriteTimeout:   18000 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	if err := s.ListenAndServe(); err != nil {
		_ = fmt.Errorf("fatal error description: %s", strings.ToLower(err.Error()))
		panic(err)

	}
}
