package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/fahad-md-kamal/go-jwt/initializers"
	"github.com/fahad-md-kamal/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

func RequrieAuth(c *gin.Context) {
	// Get the cookie off req
	tokenString, err := c.Cookie("Authorization")
	
	if err != nil  {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// Decode/validate it
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
	
		return []byte(os.Getenv("SECRET")), nil
	})

	if token == nil {
		// fmt.Println(token.Valid)
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok{

		if float64(time.Now().Unix()) > claims["exp"].(float64){
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// Find the user with token sub
		var user models.User
		initializers.DB.First(&user,claims["sub"])
		
		if user.ID == 0 {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// Attach to req
		c.Set("user", user)

		// Continue
		c.Next()

		} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}