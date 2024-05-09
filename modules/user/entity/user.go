package entity

import (
	"context"
	"gin-gorm/modules/profile/entity"
	"gin-gorm/modules/user/dto"
	"gin-gorm/modules/user/models"
	"gin-gorm/utils"
	"gin-gorm/utils/filter"
)

type User struct {
	*utils.BaseEntity
	Username string         `gorm:"column:username"`
	Password string         `gorm:"column:password"`
	Role     string         `gorm:"column:role"`
	Profile  entity.Profile `gorm:"foreignKey:UserId"`
}

type IUserService interface {
	Create(ctx context.Context, user dto.Create) (int, error)
	Update(ctx context.Context, user dto.Update, userId int) (User, error)
	Delete(ctx context.Context, userId int) error
	FindOne(ctx context.Context, userId int) (models.User, error)
	FindAll(ctx context.Context, option filter.Option) ([]models.User, error)
	FindByUsername(ctx context.Context, username string) (models.User, error)
	//FindByToken(ctx context.Context, token string)
}
type IUserRepository interface {
	Create(ctx context.Context, user User) (int, error)
	Update(ctx context.Context, user User, userId int) (User, error)
	Delete(ctx context.Context, userId int) error
	FindOne(ctx context.Context, userId int) (User, error)
	FindAll(ctx context.Context) ([]User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	//FindByToken(ctx context.Context, token string)
}
