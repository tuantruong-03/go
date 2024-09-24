package middlewares

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				// Log the panic error for debugging
				log.Printf("Recovered from panic: %v", r)

				// Respond with a 500 Internal Server Error
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Internal Server Error",
				})
				c.Abort() // Prevent further processing of the request
			}
		}()
		c.Next() // Continue to the next handler
	}
}
