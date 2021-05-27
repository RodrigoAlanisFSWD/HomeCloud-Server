package controllers

import (
	"HomeCloud/src/database/models"
	"HomeCloud/src/database/services"
	funtions "HomeCloud/src/functions"
	"encoding/json"
	"fmt"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func Register(c echo.Context) error {
	var data models.User

	if err := json.NewDecoder(c.Request().Body).Decode(&data); err != nil {
		fmt.Println(err.Error())
		return err
	}

	err, user := services.FindByUsername(data)
	fmt.Println(user)

	if user.Username == data.Username {
		return c.JSON(200, echo.Map{"res": 101, "auth": false, "token": ""})
	}

	password, err := bcrypt.GenerateFromPassword([]byte(data.Password), 5)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	data.Password = string(password)

	err = services.CreateUser(data)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	signedToken, err := funtions.CreateToken(user)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	refreshToken, err := funtions.CreateRefresh(user)

	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	return c.JSON(200, echo.Map{"res": 100, "auth": true, "token": signedToken, "refresh": refreshToken})
}

func Login(c echo.Context) error {
	var data models.User

	if err := json.NewDecoder(c.Request().Body).Decode(&data); err != nil {
		return err
	}

	err, user := services.FindByUsername(data)

	if err != nil {
		return c.JSON(200, echo.Map{"res": 101, "auth": false, "token": ""})
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password))

	if err != nil {
		return c.JSON(200, echo.Map{"res": 102, "auth": false, "token": ""})
	}

	signedToken, err := funtions.CreateToken(user)

	if err != nil {
		return err
	}

	refreshToken, err := funtions.CreateRefresh(user)

	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{"res": 100, "auth": true, "token": signedToken, "refresh": refreshToken})
}

func Refresh(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(string)

	objectId, err := primitive.ObjectIDFromHex(id)

	signedToken, err := funtions.CreateToken(models.User{ID: objectId})

	if err != nil {
		return err
	}

	return c.JSON(200, echo.Map{"res": 100, "auth": true, "token": signedToken})
}
