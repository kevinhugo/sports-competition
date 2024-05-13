package middlewares

import (
	"fmt"
	"net/http"
	"sports-competition/app/env"
	"sports-competition/app/helpers"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Check Authorization Token if provided
		authorizationHeader := c.GetHeader("Authorization")
		if authorizationHeader == "" {
			c.JSON(http.StatusUnauthorized, helpers.CreateUnauthorizedResponse("Unauthorized.", "Token not found."))
			c.Abort()
			return
		}
		tokenString := strings.Replace(authorizationHeader, "Bearer ", "", 1)
		// \ Check Authorization Token if provided

		// Check validity and decode token
		token, err := jwt.ParseWithClaims(tokenString, &helpers.AccessToken{}, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return env.ACCESS_TOKEN_SECRET_KEY, nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, helpers.CreateUnauthorizedResponse("Unauthorized.", "Invalid token."))
			c.Abort()
			return
		}
		if !token.Valid {
			c.JSON(http.StatusUnauthorized, helpers.CreateUnauthorizedResponse("Unauthorized.", "Invalid token."))
			c.Abort()
			return
		}
		// \ Check validity and decode token

		// Set claim data to struct
		tokenData, ok := token.Claims.(*helpers.AccessToken)
		if !ok {
			c.JSON(http.StatusUnauthorized, helpers.CreateUnauthorizedResponse("Unauthorized.", "Unable to parse token data."))
			c.Abort()
			return
		}
		// \ Set claim data to struct

		c.Set("tokenData", tokenData)
		c.Next()
	}
}
