package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"web-service-gin/controllers"
	_ "web-service-gin/docs"
)

var Router *gin.Engine

func CreateRoutes() {
	Router = gin.Default()

	Router.Use(errorHandler)

	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	g1 := Router.Group("api/v1/albums")
	{
		// g1.Use(auth())
		g1.GET("", controllers.GetAlbums)
		g1.POST("", auth(), controllers.PostAlbums)
		g1.GET("/:id", controllers.GetAlbumById)
	}
}

func errorHandler(c *gin.Context) {
	c.Next()

	for _, err := range c.Errors {
		fmt.Println(err)
	}

	if len(c.Errors.Errors()) != 0 {
		c.IndentedJSON(http.StatusInternalServerError, c.Errors)
	}
}

func auth() gin.HandlerFunc { // Функция промежуточного ПО
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization") // Получаем заголовок
		if len(authHeader) == 0 {                  // Проверяем, установлен ли
			// httputil.NewError(c, http.StatusUnauthorized, errors.New(""))
			err := fmt.Errorf("Authorization is required Header")
			c.Error(err)
			c.Abort()
		}
		if authHeader != "secret_key" { // Проверяем, подходит ли API-ключ
			// httputil.NewError(c, http.StatusUnauthorized, fmt.Errorf("this user isn't authorized to this operation: api_key=%s", authHeader))
			err := fmt.Errorf("Invalid Authorization key")
			c.Error(err)
			c.Abort()
		}
		c.Next()
	}
}
