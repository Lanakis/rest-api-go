package config

import (
	"authorization/src/schema"
	"authorization/src/utils"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // Postgres goalng driver
	"log"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "ais"
	password = "896325"
	dbName   = "buro_nest"
)

func DatabaseConnection() *sql.DB {
	sqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbName)

	db, err := sql.Open("postgres", sqlInfo)
	utils.PanicIfError(err)

	err = db.Ping()
	utils.PanicIfError(err)
	pgDB := &schema.PostgresDb{Db: db}
	if err := pgDB.LoadSchema(context.Background()); err != nil {
		log.Fatal("error loading database schema: ", err)
	}
	fmt.Println("Connected to database!!")
	return db
}
