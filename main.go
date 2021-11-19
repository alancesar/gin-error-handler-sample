package main

import (
	"github.com/alancesar/gin-error-handler-sample/api"
	"github.com/alancesar/gin-error-handler-sample/database"
	"github.com/alancesar/gin-error-handler-sample/middleware"
	"github.com/alancesar/gin-error-handler-sample/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func main() {
	db := database.NewDatabase()
	s := service.NewService(db)

	engine := gin.Default()
	engine.Use(middleware.ErrorHandlerMiddleware())
	engine.Handle(http.MethodPost, "/api", api.Handler(s))

	if err := engine.Run(":8080"); err != nil {
		log.Fatalln(err)
	}
}
