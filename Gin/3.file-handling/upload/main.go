package main

import (
	"fmt"
	"log"
	"mime/multipart"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Persons struct {
	Fname string                  `form:"fname" json:"fname"`
	Lname string                  `form:"lname" json:"lname"`
	Files []*multipart.FileHeader `form:"files" json:"files"`
}
type Person struct {
	Fname string                `form:"fname" json:"fname"`
	Lname string                `form:"lname" json:"lname"`
	Files *multipart.FileHeader `form:"file" json:"file"`
}

// only files
// files with other forms fields
// in postman and browser
// large files - using buffer
func main() {
	r := gin.Default()

	r.LoadHTMLGlob("./*.html")

	r.GET("form", form)
	r.POST("/saveform", formhandler)

	r.Run()
}
func form(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", gin.H{})
}

// get only specific key from form  c.PostForm("fname")
// get only specific file from form c.FormFile("filename")
// get entire map (all keys, files) from form c.MultipartForm()
func formhandler(c *gin.Context) {
	form, _ := c.MultipartForm()
	resjson := form.Value

	fmt.Println("form--------------------------------")
	fmt.Println(form.Value)
	fmt.Println(form.File)
	fmt.Println("form--------------------------------")

	for formkey, files := range form.File {
		resjson[formkey] = []string{}
		for _, file := range files { //if multiple files for one html input tag
			resjson[formkey] = append(resjson[formkey], file.Filename)
			err := c.SaveUploadedFile(file, "./"+file.Filename)
			if err != nil {
				log.Println("err saving the file: ", file.Filename, err.Error())
				c.String(400, "err saving the files")
				return
			}
		}
	}
	c.JSON(200, resjson)
}
