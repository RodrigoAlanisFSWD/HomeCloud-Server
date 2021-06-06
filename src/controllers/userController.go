package controllers

import (
	"HomeCloud/src/database/services"
	"fmt"
	"io"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func Profile(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)
	fmt.Println(id)

	objectId, err := primitive.ObjectIDFromHex(id)

	user, err := services.FindById(objectId)

	if err != nil {
		fmt.Println(err.Error())
		return c.JSON(200, echo.Map{"res": 99, "auth": false, "data": false})
	}

	return c.JSON(200, echo.Map{"res": 100, "auth": true, "data": user})
}

func Avatar(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	id := claims["id"].(string)

	fmt.Println(id)

	file, err := c.FormFile("avatar")

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	src, err := file.Open()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer src.Close()

	dst, err := os.Create("uploads/" + file.Filename)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		fmt.Println(err.Error())
		return err
	}

	err = services.UpdateAvatar(id, file.Filename)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return c.JSON(200, echo.Map{"res": 100, "auth": true, "data": "UPLOADED"})
}
