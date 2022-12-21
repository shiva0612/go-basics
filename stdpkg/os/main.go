package main

import (
	"fmt"
	"os"

	"github.com/subosito/gotenv"
)

func main() {
	// create_append_file()

	executable, _ := os.Executable() //while starting the program - just to be in sync
	dir, _ := os.Getwd()
	host, _ := os.Hostname()
	fmt.Println(host) //name of the machine as per the router
	fmt.Println(executable)
	fmt.Println(dir)
	// os.Getuid()
	// os.Getpid()

	working_with_env()

}

func working_with_env() {
	gotenv.Load(".env")
	// os.Getenv()
	// os.Clearenv()
	// os.Environ()
	// os.Setenv("NAME", "gopher")
	// os.Setenv("BURROW", "/usr/gopher")
	// os.Unsetenv()
	// fmt.Println(os.ExpandEnv("$NAME lives in ${BURROW}."))

}

func create_append_file() {
	_, err := os.OpenFile("op.go", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}
}
