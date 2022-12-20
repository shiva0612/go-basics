package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

var (
	router *gin.Engine
)

func main() {

	ech := make(chan os.Signal, 1)
	signal.Notify(ech, os.Interrupt)

	readconfig()

	createGinEngine()
	server := createServer()

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {

		fmt.Printf("\nfile: %s \noperation: %s\n", in.Name, in.Op.String())
		fmt.Printf("\nnew config:\n%v\n", viper.AllSettings())

		err := server.Shutdown(context.Background())
		if err != nil {
			log.Fatalln("error while shutting down: ", err.Error())
		}
		fmt.Printf("\nserver running on address:[%s] stopped\n", server.Addr)

		time.Sleep(time.Second)
		server = createServer()
		go startServer(server)

	})

	go startServer(server)

	<-ech
	os.Exit(1)

}

func startServer(server *http.Server) {
	fmt.Println("starting server on address: ", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		log.Println("server stopped listening: ", err.Error())
	}
}

func createGinEngine() {
	router = gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "home")
	})
}
func createServer() *http.Server {

	return &http.Server{
		Addr:    ":" + viper.GetString("port"),
		Handler: router,
	}

}
func readconfig() {
	viper.SetConfigName("config") // name of config file (without extension)
	viper.SetConfigType("json")   // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err.Error())
	}
}
