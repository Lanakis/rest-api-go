package entity

import (
	"authorization/modules/profile/models"
	"authorization/utils"
	"authorization/utils/filter"
	"context"
)

type Profile struct {
	*utils.BaseEntity
	FirstName  string `db:"first_name"`
	MiddleName string `db:"middle_name"`
	LastName   string `db:"last_name"`
	Age        int    `db:"age"`
	Head       bool   `db:"head"`
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
