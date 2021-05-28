package routers

import (
	"HomeCloud/src/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func CloudRouter(api *echo.Group) {
	cloud := api.Group("/cloud")

	cloud.Use(middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("secret")}))

	cloud.POST("/upload/:path?", controllers.UploadFile)
}
