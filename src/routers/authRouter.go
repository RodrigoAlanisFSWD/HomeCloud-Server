package routers

import (
	"HomeCloud/src/controllers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func AuthRouter(e *echo.Group) {
	auth := e.Group("/auth")

	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)
	auth.GET("/refresh", controllers.Refresh, middleware.JWTWithConfig(middleware.JWTConfig{SigningKey: []byte("secret")}))
}
