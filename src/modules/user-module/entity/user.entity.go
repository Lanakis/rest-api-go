package entity

import (
	"authorization/src/models/user"
	"authorization/src/modules/profile-module/entity"
	"authorization/src/modules/user-module/dto"
	"authorization/src/utils"
	"authorization/src/utils/filter"
	"context"
)

type UserEntity struct {
	*utils.BaseEntity
	Username string `db:"username"`
	Password string `db:"password"`
	Role     string `db:"role"`
	Profile  entity.ProfileEntity
}

type IUserService interface {
	Create(ctx context.Context, createUserDto dto.CreateUserDTO) (int, error)
	Update(ctx context.Context, updateUserDto dto.UpdateUserDTO, userId int)
	Delete(ctx context.Context, userId int)
	FindOne(ctx context.Context, userId int) (user.User, error)
	FindAll(ctx context.Context, option filter.Option) ([]user.User, error)
	FindByUsername(ctx context.Context, username string) (user.User, error)
	//FindByToken(ctx context.Context, token string)
}
type IUserRepository interface {
	Create(ctx context.Context, user UserEntity) (int, error)
	Update(ctx context.Context, user UserEntity, userId int)
	Delete(ctx context.Context, userId int)
	FindOne(ctx context.Context, userId int) (UserEntity, error)
	FindAll(ctx context.Context, option []filter.Field) ([]UserEntity, error)
	FindByUsername(ctx context.Context, username string) (UserEntity, error)
	//FindByToken(ctx context.Context, token string)
}
