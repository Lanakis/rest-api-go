package profile

import (
	"gin-gorm/modules/profile/entity"
	"gin-gorm/modules/profile/repository"
	"gorm.io/gorm"
)

func InitModule(db *gorm.DB) entity.IProfileService {

	profileRepository := repository.NewProfileRepository(db)
	profileService := NewProfileService(profileRepository)
	return profileService
}
