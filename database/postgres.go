package database

import (
	"context"
	"fmt"
	"gin-gorm/config"
	"gin-gorm/database/seed"
	profile_entity "gin-gorm/modules/profile/entity"
	user_entity "gin-gorm/modules/user/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

func Connection() *gorm.DB {
	cfg := config.GetConfig()

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgreSQL.Host, cfg.PostgreSQL.Port, cfg.PostgreSQL.Username, cfg.PostgreSQL.Password, cfg.PostgreSQL.Database)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Ошибка открытия соединения с базой данных: ", err)
	}

	// Проверка соединения
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Ошибка получения объекта базы данных: ", err)
	}
	err = sqlDB.PingContext(context.Background())
	if err != nil {
		log.Fatal("Ошибка проверки соединения с базой данных: ", err)
	}

	if !db.Migrator().HasTable(&user_entity.User{}) || !db.Migrator().HasTable(&profile_entity.Profile{}) {
		// Если таблицы не существуют, создаем их
		if err := db.AutoMigrate(&user_entity.User{}, &profile_entity.Profile{}); err != nil {
			log.Fatal("Ошибка создания таблиц базы данных: ", err)
		}

		// Seeding начальных данных
		if err := seed.SeedUsersProfiles(db); err != nil {
			log.Fatal("Ошибка при заполнении базы данных начальными данными: ", err)
		}
	}
	return db
}
