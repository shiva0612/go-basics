package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	simple()
	complex()
}
func simple() {
	router := gin.Default()
	router.Static("/assets", "./") //serving the folder

	router.StaticFile("/profilepic", "./shiva.jpg")

	router.GET("/file", func(c *gin.Context) {
		c.File("/path/file.go")
	})

}

// streaming a video or large file
func complex() {
	router := gin.Default()
	router.GET("/someDataFromReader", func(c *gin.Context) {
		response, _ := http.Get("https://raw.githubusercontent.com/gin-gonic/logo/master/color.png")

		reader := response.Body
		defer reader.Close()
		contentLength := response.ContentLength
		contentType := response.Header.Get("Content-Type")

		extraHeaders := map[string]string{
			"Content-Disposition": `attachment; filename="gopher.png"`,
		}

		c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
	})

	router.Run(":8080")
}

//to display file in browser - Content-Disposition: inline
//save as dialogue with defaultfilename - Content-Disposition: attachment
//save as dialogue with custom filename - Content-Disposition: attachment; filename="filename.jpg"
