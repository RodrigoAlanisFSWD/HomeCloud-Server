package funtions

import (
	"HomeCloud/src/database/models"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func CreateToken(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	return token.SignedString([]byte("secret"))
}

func CreateRefresh(user models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = user.ID
	claims["exp"] = time.Now().Add(time.Hour * 24 * 365).Unix()

	return token.SignedString([]byte("secret"))
}
