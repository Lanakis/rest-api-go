package profile

import (
	"authorization/modules/profile/entity"
	"authorization/modules/profile/repository"
	"database/sql"
)

func InitModule(db *sql.DB) entity.IProfileService {

	profileRepository := repository.NewProfileRepository(db)
	profileService := NewProfileService(profileRepository)
	return profileService
}
