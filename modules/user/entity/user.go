package entity

import (
	"authorization/modules/profile/entity"
	"authorization/modules/user/dto"
	"authorization/modules/user/models"
	"authorization/utils"
	"authorization/utils/filter"
	"context"
)

type User struct {
	*utils.BaseEntity
	Username string         `db:"username"`
	Password string         `db:"password"`
	Role     string         `db:"role"`
	Profile  entity.Profile `db:"foreignKey:UserId"`
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
	FindAll(ctx context.Context, option []filter.Field) ([]User, error)
	FindByUsername(ctx context.Context, username string) (User, error)
	//FindByToken(ctx context.Context, token string)
}
