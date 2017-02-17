package utils

import (
	"time"

	"twreporter.org/go-api/configs"

	"github.com/dgrijalva/jwt-go"
)

var (
	cfg          = configs.GetConfig()
	mySigningKey = []byte(cfg.APP.Token)
)

// RetrieveToken ...
func RetrieveToken(admin bool, name string, email string) string {
	// create the token
	token := jwt.New(jwt.SigningMethodHS256)

	// create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)

	// set token claims
	claims["admin"] = admin
	claims["name"] = name
	claims["email"] = email
	claims["exp"] = time.Now().Add(time.Hour * cfg.APP.Expiration).Unix()

	/* Sign the token with our secret */
	tokenString, _ := token.SignedString(mySigningKey)

	return tokenString
}