package main

import (
	"io"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
}

func best_practice(c *gin.Context) {
	r := gin.Default()
	r.Use(gin.Logger())

	logfile, _ := os.Open("someifle")
	gin.DefaultWriter = io.MultiWriter(logfile, os.Stdout)

	// when using concurrent handler make sure to copy the context
	c.Copy()
	// else the context will be terminated

}
