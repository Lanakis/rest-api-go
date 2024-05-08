package database

import (
	"authorization/config"
	"authorization/database/schema"
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func Connection() *sql.DB {

	cfg := config.GetConfig()

	sqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgreSQL.Host, cfg.PostgreSQL.Port, cfg.PostgreSQL.Username, cfg.PostgreSQL.Password, cfg.PostgreSQL.Database)

	db, err := sql.Open("postgres", sqlInfo)
	if err != nil {
		log.Fatal("ошибка открытия соединения с базой данных: ", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("ошибка проверки соединения с базой данных: ", err)
	}

	pgDB := &schema.PostgresDb{Db: db}

	if err := pgDB.LoadSchema(context.Background()); err != nil {
		log.Fatal("ошибка загрузки схемы базы данных: ", err)
	}

	fmt.Println("Соединение с базой данных установлено!!")
	return db
}
