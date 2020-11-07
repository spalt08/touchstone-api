package middleware

import (
	"fmt"
	"jsbnch/pkg/model"
	"jsbnch/pkg/utils/env"
	"strconv"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var hmacSecret = []byte(env.GetJwtSecret())

// CreateToken for authorized API calls
func CreateToken(user *model.User) (string, *GenericAPIError) {
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": strconv.FormatInt(user.ID, 10),
	})

	tokenString, err := jwtToken.SignedString(hmacSecret)

	if err != nil {
		return tokenString, NewInternalError(err)
	}

	return tokenString, nil
}

// ParseToken checks header token and returns the userId
func ParseToken(ctx *gin.Context) (int64, *GenericAPIError) {
	userID := int64(0)
	bearToken := ctx.GetHeader("Authorization")
	bearAttributes := strings.Split(bearToken, " ")

	if len(bearAttributes) != 2 {
		return userID, NewBindingError(fmt.Errorf("Invalid Authorization Header"))
	}

	tokenString := bearAttributes[1]

	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecret, nil
	})

	if claims, ok := jwtToken.Claims.(jwt.MapClaims); ok && jwtToken.Valid {
		userID, err = strconv.ParseInt(claims["userId"].(string), 10, 64)
	}

	if err != nil {
		return userID, NewUnauthorizedError(err)
	}

	return userID, nil
}
