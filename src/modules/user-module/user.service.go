package user_module

import (
	"authorization/src/models/profile"
	"authorization/src/models/user"
	"authorization/src/utils/filter"

	//profile_module "authorization/src/modules/profile-module"
	profile_entity "authorization/src/modules/profile-module/entity"
	"authorization/src/modules/user-module/dto"
	"authorization/src/modules/user-module/entity"
	"authorization/src/utils"
	"context"
	"fmt"
	"time"
)

type UserService struct {
	UserRepository entity.IUserRepository
	ProfileService profile_entity.IProfileService
}

func NewUserService(userRepository entity.IUserRepository, profileService profile_entity.IProfileService) entity.IUserService {
	return &UserService{UserRepository: userRepository, ProfileService: profileService}
}

func (s *UserService) Create(ctx context.Context, createUserDto dto.CreateUserDTO) (int, error) {
	userEntity := entity.UserEntity{
		Username: createUserDto.Username,
		Password: createUserDto.Password,
		Role:     createUserDto.Role,
	}

	profileEntity := profile.Profile{
		FirstName:  createUserDto.FirstName,
		MiddleName: createUserDto.MiddleName,
		LastName:   createUserDto.LastName,
		Age:        createUserDto.Age,
		Head:       createUserDto.Head,
	}

	// Создать пользователя и получить его id
	userID, err := s.UserRepository.Create(ctx, userEntity)
	if err != nil {
		return 0, err
	}
	fmt.Printf("%v\n UserId ", userID)

	// Создать профиль

	s.ProfileService.Create(ctx, profileEntity, userID)

	return userID, nil
}

func (s *UserService) FindAll(ctx context.Context, option filter.Option) ([]user.User, error) {
	filterOption := option.Fields()
	usersResponse, _ := s.UserRepository.FindAll(ctx, filterOption)
	var userResp []user.User
	for _, value := range usersResponse {
		user := user.User{
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
		userResp = append(userResp, user)
	}
	return userResp, nil
}

func (s *UserService) FindOne(ctx context.Context, userId int) (user.User, error) {
	userResponse, _ := s.UserRepository.FindOne(ctx, userId)
	el := user.User{
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

func (s *UserService) Update(ctx context.Context, updateUserDto dto.UpdateUserDTO, userId int) {
	currentTime := time.Now()
	userEntity := entity.UserEntity{
		BaseEntity: &utils.BaseEntity{UpdatedAt: currentTime},
		Username:   updateUserDto.Username,
		Password:   updateUserDto.Password,
		Role:       updateUserDto.Role,
		Profile: profile_entity.ProfileEntity{
			BaseEntity: &utils.BaseEntity{UpdatedAt: currentTime},
			FirstName:  updateUserDto.FirstName,
			MiddleName: updateUserDto.MiddleName,
			LastName:   updateUserDto.LastName,
			Age:        updateUserDto.Age,
			Head:       updateUserDto.Head,
		},
	}
	s.UserRepository.Update(ctx, userEntity, userId)
}

func (s *UserService) Delete(ctx context.Context, userId int) {
	s.UserRepository.Delete(ctx, userId)
}

func (s *UserService) FindByUsername(ctx context.Context, username string) (user.User, error) {
	userResponse, _ := s.UserRepository.FindByUsername(ctx, username)
	el := user.User{
		Id:       userResponse.Id,
		Username: userResponse.Username,
		Password: userResponse.Password,
	}
	return el, nil
}
