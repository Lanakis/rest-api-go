package modules

import (
	"authorization/src/config"
	auth "authorization/src/modules/auth-module"
	user "authorization/src/modules/user-module"
)

func InitModule() {
	db := config.DatabaseConnection()
	userService := user.InitModule(db)
	auth.InitModule(userService)
	//	profile_module.InitModule(db)

}
