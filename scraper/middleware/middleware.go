package middleware

import (
	"errors"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getApiKey() string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv("API_KEY")
}

func ValidateApiKey() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("X-API-KEY")
		value := getApiKey()
		if key != value {
			c.AbortWithError(http.StatusUnauthorized, errors.New("invalid API key"))
			return
		}
	}
}
