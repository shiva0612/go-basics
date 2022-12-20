package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.Use(middleware)
	r.GET("/", handler)
	r.Run()
}

func handler(c *gin.Context) {
	c.String(200, "ok")
}

func middleware(c *gin.Context) {
	fmt.Println("in middleware")
	c.Next()
}
