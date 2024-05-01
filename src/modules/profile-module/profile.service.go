package profile_module

import (
	"authorization/src/models/profile"
	"authorization/src/modules/profile-module/entity"
	"authorization/src/utils"
	"authorization/src/utils/filter"
	"context"
	"fmt"
	"time"
)

type ProfileService struct {
	ProfileRepository entity.IProfileRepository
}

func NewProfileService(profileRepository entity.IProfileRepository) entity.IProfileService {
	return &ProfileService{ProfileRepository: profileRepository}
}

func (s *ProfileService) Create(ctx context.Context, profile profile.Profile, userId int) {

	profileEntity := entity.ProfileEntity{
		FirstName:  profile.FirstName,
		MiddleName: profile.MiddleName,
		LastName:   profile.LastName,
		Age:        profile.Age,
		Head:       profile.Head,
	}
	s.ProfileRepository.Create(ctx, profileEntity, userId)

}

func (s *ProfileService) FindAll(ctx context.Context, filterOptions filter.Option) ([]profile.Profile, error) {
	a := filterOptions.Fields()
	fmt.Printf("%+v\n Поля ", a)
	profilesResponse, _ := s.ProfileRepository.FindAll(ctx, a)

	var profileResp []profile.Profile

	for _, value := range profilesResponse {
		profile := profile.Profile{

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

func (s *ProfileService) FindQuery(ctx context.Context, filterOptions filter.Option) ([]entity.ProfileEntity, error) {
	a := filterOptions.Fields()
	fmt.Printf("%+v\n", a)
	profilesResponse, _ := s.ProfileRepository.FindAll(ctx, a)

	var profileResp []entity.ProfileEntity

	for _, value := range profilesResponse {
		profile := entity.ProfileEntity{
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

func (s *ProfileService) FindOne(ctx context.Context, id int) (profile.Profile, error) {
	profileResponse, _ := s.ProfileRepository.FindOne(ctx, id)
	el := profile.Profile{

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

func (s *ProfileService) Update(ctx context.Context, id int, profile profile.Profile) {
	currentTime := time.Now()
	profileEntity := entity.ProfileEntity{
		BaseEntity: &utils.BaseEntity{
			UpdatedAt: currentTime},
		FirstName:  profile.FirstName,
		MiddleName: profile.MiddleName,
		LastName:   profile.LastName,
		Age:        profile.Age,
		Head:       profile.Head,
	}
	s.ProfileRepository.Update(ctx, id, profileEntity)
}

func (s *ProfileService) Delete(ctx context.Context, id int) {
	s.ProfileRepository.Delete(ctx, id)
}
