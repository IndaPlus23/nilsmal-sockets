package actions

import (
	"net/http"
	"fmt"

	"github.com/gin-gonic/gin"
)


func GetBlogFeed(c *gin.Context) {
	// Set headers to allow all origins
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
	c.Header("Access-Control-Allow-Headers", "Content-Type")

	// If it's a preflight request, return immediately with a 200 status
	if c.Request.Method == http.MethodOptions {
		fmt.Println("Handling preflight request")
		c.AbortWithStatus(http.StatusOK)
		return
	}

    fmt.Println("Hello world!")

    c.String(http.StatusOK, "")
}