package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/talk", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hello Miley! Let's talk for hours.",
		})
	})

	err := r.Run()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
