package entity

import (
	"authorization/src/models/profile"
	"authorization/src/utils"
	"authorization/src/utils/filter"
	"context"
)

type ProfileEntity struct {
	*utils.BaseEntity
	FirstName  string `db:"first_name"`
	MiddleName string `db:"middle_name"`
	LastName   string `db:"last_name"`
	Age        int    `db:"age"`
	Head       bool   `db:"head"`
}

type IProfileService interface {
	Create(ctx context.Context, profile profile.Profile, userId int)
	FindOne(ctx context.Context, id int) (profile.Profile, error)
	FindAll(ctx context.Context, option filter.Option) ([]profile.Profile, error)
	//FindQuery(ctx context.Context, option filter.Option) ([]ProfileEntity, error)
	Update(ctx context.Context, id int, profileDto profile.Profile)
	Delete(ctx context.Context, id int)
}

type IProfileRepository interface {
	Create(ctx context.Context, profile ProfileEntity, userId int)
	FindOne(ctx context.Context, id int) (ProfileEntity, error)
	FindAll(ctx context.Context, option []filter.Field) ([]ProfileEntity, error)
	//FindQuery(ctx context.Context, option []filter.Field) ([]profile.Profile, error)
	Update(ctx context.Context, id int, profile ProfileEntity)
	Delete(ctx context.Context, id int)
}
