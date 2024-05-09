package user

import (
	"context"
	profile_entity "gin-gorm/modules/profile/entity"
	profile_models "gin-gorm/modules/profile/models"
	"gin-gorm/modules/user/dto"
	"gin-gorm/modules/user/entity"
	"gin-gorm/modules/user/models"
	"gin-gorm/utils"
	"gin-gorm/utils/filter"

	"time"
)

type Service struct {
	UserRepository entity.IUserRepository
	ProfileService profile_entity.IProfileService
}

func NewUserService(userRepository entity.IUserRepository, profileService profile_entity.IProfileService) entity.IUserService {
	return &Service{UserRepository: userRepository, ProfileService: profileService}
}

func (s *Service) Create(ctx context.Context, createUserDto dto.Create) (int, error) {
	userEntity := entity.User{
		Username: createUserDto.Username,
		Password: createUserDto.Password,
		Role:     createUserDto.Role,
	}

	profileEntity := profile_models.Profile{
		FirstName:  createUserDto.FirstName,
		MiddleName: createUserDto.MiddleName,
		LastName:   createUserDto.LastName,
		Age:        createUserDto.Age,
		Head:       createUserDto.Head,
	}

	userID, err := s.UserRepository.Create(ctx, userEntity)
	if err != nil {
		return 0, utils.NewAppError(err, "Can't create user", "user service")
	}

	errProfile := s.ProfileService.Create(ctx, profileEntity, userID)
	if errProfile != nil {
		return 0, utils.NewAppError(errProfile, "Can't create profile", "user service")

	}

	return userID, nil
}

func (s *Service) FindAll(ctx context.Context, option filter.Option) ([]models.User, error) {
	//	filterOption := option.Fields()
	usersResponse, _ := s.UserRepository.FindAll(ctx)
	var userResp []models.User
	for _, value := range usersResponse {
		foundedUser := models.User{
			Id:         value.Id,
			CreatedAt:  value.CreatedAt,
			UpdatedAt:  value.UpdatedAt,
			Username:   value.Username,
			Role:       value.Role,
			FirstName:  value.Profile.FirstName,
			MiddleName: value.Profile.MiddleName,
			LastName:   value.Profile.LastName,
			Age:        value.Profile.Age,
			Head:       value.Profile.Head,
		}
		userResp = append(userResp, foundedUser)
	}
	return userResp, nil
}

func (s *Service) FindOne(ctx context.Context, userId int) (models.User, error) {
	userResponse, _ := s.UserRepository.FindOne(ctx, userId)
	el := models.User{
		Id:         userResponse.Id,
		CreatedAt:  userResponse.CreatedAt,
		UpdatedAt:  userResponse.UpdatedAt,
		Username:   userResponse.Username,
		Role:       userResponse.Role,
		FirstName:  userResponse.Profile.FirstName,
		MiddleName: userResponse.Profile.MiddleName,
		LastName:   userResponse.Profile.LastName,
		Age:        userResponse.Profile.Age,
		Head:       userResponse.Profile.Head,
	}
	return el, nil

}

func (s *Service) Update(ctx context.Context, updateUserDto dto.Update, userId int) (entity.User, error) {
	currentTime := time.Now()
	userEntity := entity.User{
		BaseEntity: &utils.BaseEntity{UpdatedAt: currentTime},
		Username:   updateUserDto.Username,
		Password:   updateUserDto.Password,
		Role:       updateUserDto.Role,
		Profile: profile_entity.Profile{
			BaseEntity: &utils.BaseEntity{UpdatedAt: currentTime},
			FirstName:  updateUserDto.FirstName,
			MiddleName: updateUserDto.MiddleName,
			LastName:   updateUserDto.LastName,
			Age:        updateUserDto.Age,
			Head:       updateUserDto.Head,
		},
	}
	update, err := s.UserRepository.Update(ctx, userEntity, userId)
	if err != nil {
		return update, err
	}
	return update, nil
}

func (s *Service) Delete(ctx context.Context, userId int) error {
	err := s.UserRepository.Delete(ctx, userId)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) FindByUsername(ctx context.Context, username string) (models.User, error) {
	userResponse, _ := s.UserRepository.FindByUsername(ctx, username)
	el := models.User{
		Id:       userResponse.Id,
		Username: userResponse.Username,
		Password: userResponse.Password,
	}
	return el, nil
}
