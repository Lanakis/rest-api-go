package main

import (
	"gin-gorm/api"
	"gin-gorm/database"
	"gin-gorm/modules"
)

func main() {
	router := api.Init()
	db := database.Connection()
	modules.InitModule(db, router)
	select {}
}
