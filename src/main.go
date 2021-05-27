package main

import (
	"HomeCloud/src/database"
	"HomeCloud/src/routers"
	"fmt"
	"log"

	"github.com/labstack/echo/v4"
)

func main() {
	err := database.Open()

	if err != nil {
		log.Fatal("Error By Connecting The Database")
		return
	}

	fmt.Println("Database Connected")

	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Welcome To HomeCloud Server")
	})

	api := e.Group("/api")

	routers.AuthRouter(api)
	routers.UserRouter(api)

	e.Logger.Fatal(e.Start(":3000"))
}
