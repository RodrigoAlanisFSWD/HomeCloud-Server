package routers

import (
	"HomeCloud/src/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CloudRouter(api *echo.Group) {
	cloud := api.Group("/cloud")

	cloud.Use(middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("secret")}))

	cloud.POST("/upload/:path", controllers.UploadFile)
	cloud.POST("/mkdir/:path", controllers.CreateDir)
	cloud.DELETE("/rm/:name/:type/:path", controllers.Delete)
	cloud.GET("/read/:path", controllers.Read)
	cloud.GET("/download/:name/:path", controllers.Download)
}
