package main

import (
	_ "web-service-gin/docs"
	"web-service-gin/routers"
)

// @title My first Swagger API
// @version 1.0
// @description My first Swagger API for Golang/Gin
// @termsOfService http://swagger.io/terms/

// @contact.name muratom
// @contact.email muratom73@gmail.com

// @BasePath /api/v1

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	routers.CreateRoutes()
	routers.Router.Run(":8080")
}
