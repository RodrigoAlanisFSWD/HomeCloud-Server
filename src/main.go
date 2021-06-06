package main

import (
	"HomeCloud/src/database"
	"HomeCloud/src/routers"
	"fmt"
	"log"
	"path"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	err := database.Open()

	if err != nil {
		log.Fatal("Error By Connecting The Database")
		return
	}

	fmt.Println("Database Connected")

	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
	}))

	e.GET("assets/:file", func(c echo.Context) error {
		file := path.Join("uploads/" + c.Param("file"))
		return c.File(file)
	})

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome To HomeCloud Server")
	})

	api := e.Group("/api")

	routers.AuthRouter(api)
	routers.UserRouter(api)
	routers.CloudRouter(api)

	e.Logger.Fatal(e.Start(":5000"))
}
