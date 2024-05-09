package profile

import (
	"context"
	"fmt"
	"gin-gorm/modules/profile/entity"
	"gin-gorm/modules/profile/models"
	"gin-gorm/utils"
	"gin-gorm/utils/filter"
	"time"
)

type Service struct {
	ProfileRepository entity.IProfileRepository
}

func NewProfileService(profileRepository entity.IProfileRepository) entity.IProfileService {
	return &Service{ProfileRepository: profileRepository}
}

func (s *Service) Create(ctx context.Context, profile models.Profile, userId int) error {

	profileEntity := entity.Profile{
		FirstName:  profile.FirstName,
		MiddleName: profile.MiddleName,
		LastName:   profile.LastName,
		Age:        profile.Age,
		Head:       profile.Head,
	}
	err := s.ProfileRepository.Create(ctx, profileEntity, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) FindAll(ctx context.Context, filterOptions filter.Option) ([]models.Profile, error) {
	a := filterOptions.Fields()
	fmt.Printf("%+v\n Поля ", a)
	profilesResponse, _ := s.ProfileRepository.FindAll(ctx, a)

	var profileResp []models.Profile

	for _, value := range profilesResponse {
		profile := models.Profile{

			Id:         value.Id,
			CreatedAt:  value.CreatedAt,
			UpdatedAt:  value.UpdatedAt,
			FirstName:  value.FirstName,
			MiddleName: value.MiddleName,
			LastName:   value.LastName,
			Age:        value.Age,
			Head:       value.Head,
		}
		profileResp = append(profileResp, profile)
	}
	return profileResp, nil
}

func (s *Service) FindQuery(ctx context.Context, filterOptions filter.Option) ([]entity.Profile, error) {
	a := filterOptions.Fields()
	fmt.Printf("%+v\n", a)
	profilesResponse, _ := s.ProfileRepository.FindAll(ctx, a)

	var profileResp []entity.Profile

	for _, value := range profilesResponse {
		profile := entity.Profile{
			BaseEntity: &utils.BaseEntity{
				Id:        value.Id,
				CreatedAt: value.CreatedAt,
				UpdatedAt: value.UpdatedAt,
			},
			FirstName:  value.FirstName,
			MiddleName: value.MiddleName,
			LastName:   value.LastName,
			Age:        value.Age,
			Head:       value.Head,
		}
		profileResp = append(profileResp, profile)
	}
	return profileResp, nil
}

func (s *Service) FindOne(ctx context.Context, id int) (models.Profile, error) {
	profileResponse, _ := s.ProfileRepository.FindOne(ctx, id)
	el := models.Profile{

		Id:         profileResponse.Id,
		CreatedAt:  profileResponse.CreatedAt,
		UpdatedAt:  profileResponse.UpdatedAt,
		FirstName:  profileResponse.FirstName,
		MiddleName: profileResponse.MiddleName,
		LastName:   profileResponse.LastName,
		Age:        profileResponse.Age,
		Head:       profileResponse.Head,
	}
	return el, nil

}

func (s *Service) Update(ctx context.Context, id int, profile models.Profile) (entity.Profile, error) {
	currentTime := time.Now()
	profileEntity := entity.Profile{
		BaseEntity: &utils.BaseEntity{
			UpdatedAt: currentTime},
		FirstName:  profile.FirstName,
		MiddleName: profile.MiddleName,
		LastName:   profile.LastName,
		Age:        profile.Age,
		Head:       profile.Head,
	}
	update, err := s.ProfileRepository.Update(ctx, id, profileEntity)
	if err != nil {
		return profileEntity, utils.NewAppError(err, "Can't Update", "Update profile service")
	}
	return update, nil
}

func (s *Service) Delete(ctx context.Context, id int) error {
	err := s.ProfileRepository.Delete(ctx, id)
	if err != nil {
		return err
	}
	return nil
}
