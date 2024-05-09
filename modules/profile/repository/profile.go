package repository

import (
	"context"
	"gin-gorm/modules/profile/entity"
	"gin-gorm/utils/filter"
	"gorm.io/gorm"
)

type Repository struct {
	Db *gorm.DB
}

func NewProfileRepository(Db *gorm.DB) entity.IProfileRepository {
	return &Repository{Db: Db}
}

func (r *Repository) Create(ctx context.Context, profile entity.Profile, userId int) error {
	profile.UserId = userId
	result := r.Db.WithContext(ctx).Create(&profile)
	return result.Error
}

func (r *Repository) FindAll(ctx context.Context, option []filter.Field) ([]entity.Profile, error) {
	var profiles []entity.Profile
	result := r.Db.WithContext(ctx).Where(option).Find(&profiles)
	if result.Error != nil {
		return nil, result.Error
	}
	return profiles, nil
}

func (r *Repository) FindOne(ctx context.Context, profileId int) (entity.Profile, error) {
	var profile entity.Profile
	result := r.Db.WithContext(ctx).First(&profile, profileId)
	if result.Error != nil {
		return entity.Profile{}, result.Error
	}
	return profile, nil
}

func (r *Repository) Update(ctx context.Context, profileId int, profile entity.Profile) (entity.Profile, error) {
	_ = r.Db.WithContext(ctx).Model(&entity.Profile{}).Where("id = ?", profileId).Updates(profile)
	foundedProfile, _ := r.FindOne(ctx, profileId)
	return foundedProfile, nil
}

func (r *Repository) Delete(ctx context.Context, profileId int) error {
	result := r.Db.WithContext(ctx).Delete(&entity.Profile{}, profileId)
	return result.Error
}
