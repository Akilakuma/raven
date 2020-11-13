package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func test(c *gin.Context) {
	fmt.Println("🐛 hello")
	c.Status(http.StatusOK)
}
