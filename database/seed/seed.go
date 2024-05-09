package seed

import (
	profile_entity "gin-gorm/modules/profile/entity"
	user_entity "gin-gorm/modules/user/entity"
	"gorm.io/gorm"
)

func SeedUsersProfiles(db *gorm.DB) error {
	// Создание начальных пользователей
	users := []user_entity.User{
		{Username: "admin", Password: "123456", Role: "admin"},
		{Username: "user2", Password: "password2", Role: "user"},
		{Username: "user3", Password: "password3", Role: "user"},
	}
	for _, user := range users {
		if err := db.Create(&user).Error; err != nil {
			return err
		}
	}

	// Создание профилей для начальных пользователей
	profiles := []profile_entity.Profile{
		{FirstName: "John", LastName: "Doe", Age: 30, Head: true, UserId: 1},
		{FirstName: "Jane", LastName: "Smith", Age: 25, Head: false, UserId: 2},
		{FirstName: "Michael", LastName: "Johnson", Age: 35, Head: true, UserId: 3},
	}
	for _, profile := range profiles {
		if err := db.Create(&profile).Error; err != nil {
			return err
		}
	}

	return nil
}
