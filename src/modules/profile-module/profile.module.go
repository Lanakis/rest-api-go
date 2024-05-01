package profile_module

import (
	"authorization/src/modules/profile-module/entity"
	"authorization/src/modules/profile-module/repository"
	"database/sql"
)

func InitModule(db *sql.DB) entity.IProfileService {

	profileRepository := repository.NewProfileRepository(db)
	profileService := NewProfileService(profileRepository)
	return profileService
}
