package controllers

import (
	"HomeCloud/src/database/services"
	"HomeCloud/src/functions"
	"fmt"
	"io"
	"os"
	"path"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UploadFile(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	fmt.Println(claims)
	id := claims["id"].(string)

	userId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	user, err := services.FindById(userId)
	fmt.Println(user)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	form, err := c.MultipartForm()

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	files := form.File["file"]

	for _, file := range files {
		formatedPath := functions.FormatPath(c.Param("path"))

		src, err := file.Open()

		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		defer src.Close()

		newPath := path.Join("cloud/" + user.Username + formatedPath + "/")
		finalPath := path.Join(newPath + "/" + file.Filename)

		dst, err := os.Create(finalPath)

		if err != nil {
			fmt.Println(err.Error())
			return err
		}

		defer dst.Close()

		if _, err := io.Copy(dst, src); err != nil {
			fmt.Println(err.Error())
			return err
		}
	}

	return c.JSON(200, echo.Map{"res": 100, "auth": true, "data": "UPLOADED"})
}
