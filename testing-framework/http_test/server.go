package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func run_gin_server() {
	router := setup()
	log.Fatalln(router.Run(":8087"))
}

func setup() *gin.Engine {
	router := gin.Default()
	router.GET("/", Homegin)
	return router
}
func Homegin(c *gin.Context) {
	c.String(http.StatusOK, "hello from gin")
}

// ----------------------------------------------------------------
func run_normal_server() {
	http.HandleFunc("/", Homenormal)
	log.Fatalln(http.ListenAndServe(":8081", nil))
}

func Homenormal(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("hello from normal"))
}
