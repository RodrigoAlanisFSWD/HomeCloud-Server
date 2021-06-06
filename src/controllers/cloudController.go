package controllers

import (
	"HomeCloud/src/database/services"
	"HomeCloud/src/functions"
	"encoding/json"
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

		newPath := path.Join("cloud/" + user.Username + "/" + formatedPath + "/")
		finalPath := path.Join(newPath + "/" + file.Filename)

		fmt.Println(finalPath)

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

func CreateDir(c echo.Context) error {
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

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	var body map[string]string

	if err := json.NewDecoder(c.Request().Body).Decode(&body); err != nil {
		fmt.Println(err.Error())
		return err
	}

	formatedPath := functions.FormatPath(c.Param("path"))
	fmt.Println(formatedPath)
	newPath := path.Join("cloud/" + user.Username + "/" + formatedPath + "/" + body["name"] + "/")

	err = os.Mkdir(newPath, 0777)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return c.JSON(200, echo.Map{"res": 100, "auth": true, "data": "CREATED"})
}

func Delete(c echo.Context) error {
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

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	formatedPath := functions.FormatPath(c.Param("path"))
	newPath := path.Join("cloud/" + user.Username + "/" + formatedPath + "/" + c.Param("name"))
	fmt.Println(newPath)

	if c.Param("type") == "dir" {
		err = os.RemoveAll(newPath)
	} else {
		err = os.Remove(newPath)
	}

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return c.JSON(200, echo.Map{"res": 100, "auth": true, "data": "DELETED"})
}

func Read(c echo.Context) error {
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

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	formatedPath := functions.FormatPath(c.Param("path"))
	newPath := path.Join("cloud/" + user.Username + "/" + formatedPath + "/")

	read, err := os.ReadDir(newPath)

	var files []string

	for _, item := range read {
		files = append(files, item.Name())
	}

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return c.JSON(200, echo.Map{"res": 100, "auth": true, "data": files})
}

func Download(c echo.Context) error {
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

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	formatedPath := functions.FormatPath(c.Param("path"))
	newPath := path.Join("cloud/" + user.Username + "/" + formatedPath + "/")
	finalPath := path.Join(newPath + "/" + c.Param("name"))

	return c.File(finalPath)
}
