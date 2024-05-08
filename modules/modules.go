package modules

import (
	"authorization/modules/auth"
	"authorization/modules/user"
	"database/sql"
)

func InitModule(db *sql.DB) {
	userService := user.InitModule(db)
	auth.InitModule(userService)
}
