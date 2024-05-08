package main

import (
	"authorization/api"
	"authorization/database"
	"authorization/modules"
)

func main() {
	api.Init()
	db := database.Connection()
	modules.InitModule(db)
	select {}
}
