package helpers

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

func CreateToken(userid string) (string, int64, error) {
	var err error
	//Creating Access Token

	expiresIn := time.Now().Add(time.Minute * 15).Unix()

	os.Setenv("ACCESS_SECRET", "jdnfksdmfksd") //this should be in an env file
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userid
	atClaims["exp"] = expiresIn
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)

	token, err := at.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
	if err != nil {
		return "", 0, err
	}
	return token, expiresIn, nil
}

func ValidateToken(token string) (bool, error) {
	at, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return false, err
	}
	if _, ok := at.Claims.(jwt.Claims); !ok && !at.Valid {
		return false, err
	}
	return true, nil
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
