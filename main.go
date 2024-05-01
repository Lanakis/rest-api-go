package main

import (
	"authorization/src/api"
	"authorization/src/modules"
)

func main() {
	api.Init()
	modules.InitModule()
	select {}
}
