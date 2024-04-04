package utils

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/tarun-kavipurapu/gin-gorm/db/models"
)

var privateKey []byte = []byte("cEwHkXr2u5x8A/B?D(G+KbPeShVmYq3t6v9y$B&E)H@McQfTjWnZr4u7x!A%C*F-JaN")

func GenerateJWT(user models.User) (string, error) {
	tokenTTL, _ := strconv.Atoi("3600")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   user.UserID,
		"user_role": user.UserRoleID,
		"iat":       time.Now().Unix(),
		"eat":       time.Now().Add(time.Second * time.Duration(tokenTTL)).Unix(),
	})
	return token.SignedString(privateKey)
}

func ValidateJWT(context *gin.Context) error {
	token, err := ExtractJWT(context)
	if err != nil {
		return err
	}
	_, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		return nil
	}
	return errors.New("invalid token provided")
}

// /validate Restraunt
func ValidateRestaurantRoleJWT(context *gin.Context) error {
	token, err := ExtractJWT(context)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["user_role"].(float64))
	if ok && token.Valid && userRole == 2 {
		return nil
	}
	return errors.New("invalid Restaurant token provided")
}

// validate Customer Role
func ValidateCustomerRoleJWT(context *gin.Context) error {
	token, err := ExtractJWT(context)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	userRole := uint(claims["user_role"].(float64))
	if ok && token.Valid && userRole == 1 {
		return nil
	}
	return errors.New("invalid auth token provided")
}

//

func CurrentUser(context *gin.Context) models.User {
	err := ValidateJWT(context)
	if err != nil {
		return models.User{}
	}
	token, _ := ExtractJWT(context)
	claims, _ := token.Claims.(jwt.MapClaims)
	userId := uint(claims["user_id"].(float64))

	user, err := models.GetUserById(userId)
	if err != nil {
		return models.User{}
	}
	return user
}
func ExtractJWT(context *gin.Context) (*jwt.Token, error) {
	tokenString := ExtractFromRequest(context)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return privateKey, nil
	})
	return token, err
}

func ExtractFromRequest(context *gin.Context) string {
	bearerToken := context.Request.Header.Get("Authorization")
	splitToken := strings.Split(bearerToken, " ")
	if len(splitToken) == 2 {
		return splitToken[1]
	}
	return ""
}
