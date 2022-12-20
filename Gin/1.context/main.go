package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Person struct {
	Name string
	Age  string
}

func all_ctx_functions(c *gin.Context) {

	router := gin.Default()

	c.Set("key", "value") //key values for this context

	c.Param("name")
	c.Query("name")

	c.PostForm("name")
	c.GetPostForm("name") // value, true
	c.FormFile("pic")     // give the first file
	c.MultipartForm()     //

	c.DefaultQuery("name", "default value")
	c.DefaultPostForm("name", "default value")

	//not imp
	c.QueryMap("name")
	c.PostFormMap("name")

	c.String(400, "message", "values", "values")
	c.JSON(400, gin.H{"key": "value"})
	c.JSON(200, Person{"shiva", "24"}) //check the json tag for Person struct
	c.XML(400, gin.H{"key": "value"})
	c.File("/path/filename/file.go")
	c.Data(200, "application/json", []byte("shiva"))
	// c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)

	c.FullPath()
	c.Cookie("key")
	// c.SetCookie("key","value",...)
	c.Abort()                  //for aborting in middlewares
	c.AddParam("key", "value") //for internal routing
	c.ClientIP()
	c.Copy() //copy context if passing in goroutine

	c.Redirect(504, "newurl")

	ctxTime, _ := c.Deadline()
	<-c.Done()
	fmt.Println(ctxTime)

	router.GET("/long_async", func(c *gin.Context) {
		// create copy to be used inside the goroutine
		cCp := c.Copy()
		go func() {
			// simulate a long task with time.Sleep(). 5 seconds
			time.Sleep(5 * time.Second)

			// note that you are using the copied context "cCp", IMPORTANT
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	router.Run("8080")
	http.ListenAndServe(":8080", router)

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
