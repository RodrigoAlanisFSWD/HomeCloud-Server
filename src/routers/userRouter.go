package routers

import (
	"HomeCloud/src/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func UserRouter(e *echo.Group) {
	user := e.Group("/user")

	user.Use(middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("secret")}))

	user.GET("/profile", controllers.Profile)
	user.POST("/avatar", controllers.Avatar)
}
