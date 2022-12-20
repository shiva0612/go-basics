package main

import (
	"shiva/packaging/database"
	db "shiva/packaging/database/db"
	db2 "shiva/packaging/database/db2"
)

var (
	//import module & main.Main_var
	Main_var = "main var"
)

func main() {
	_ = db.Name1
	_ = database.Name
	_ = db2.Name2
}
