package entity

import (
	"context"
	"gin-gorm/modules/profile/models"
	"gin-gorm/utils"
	"gin-gorm/utils/filter"
)

type Profile struct {
	*utils.BaseEntity
	FirstName  string `gorm:"column:first_name"`
	MiddleName string `gorm:"column:middle_name"`
	LastName   string `gorm:"column:last_name"`
	Age        int    `gorm:"column:age"`
	Head       bool   `gorm:"column:head"`
	UserId     int    `gorm:"unique;column:user_id"`
}

type IProfileService interface {
	Create(ctx context.Context, profile models.Profile, userId int) error
	FindOne(ctx context.Context, id int) (models.Profile, error)
	FindAll(ctx context.Context, option filter.Option) ([]models.Profile, error)
	//FindQuery(ctx context.Context, option filter.Option) ([]ProfileEntity, error)
	Update(ctx context.Context, id int, profile models.Profile) (Profile, error)
	Delete(ctx context.Context, id int) error
}

type IProfileRepository interface {
	Create(ctx context.Context, profile Profile, userId int) error
	FindOne(ctx context.Context, id int) (Profile, error)
	FindAll(ctx context.Context, option []filter.Field) ([]Profile, error)
	//FindQuery(ctx context.Context, option []filter.Field) ([]profile.Profile, error)
	Update(ctx context.Context, id int, profile Profile) (Profile, error)
	Delete(ctx context.Context, id int) error
}
