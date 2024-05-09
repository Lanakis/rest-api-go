package repository

import (
	"context"
	"gin-gorm/modules/user/entity"
	"gin-gorm/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	Db *gorm.DB
}

func NewUserRepository(Db *gorm.DB) entity.IUserRepository {
	return &UserRepository{Db: Db}
}

func (r *UserRepository) Create(ctx context.Context, user entity.User) (int, error) {
	result := r.Db.WithContext(ctx).Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	return user.Id, nil
}

func (r *UserRepository) FindAll(ctx context.Context) ([]entity.User, error) {
	var users []entity.User
	result := r.Db.WithContext(ctx).
		Model(&entity.User{}).
		Preload("Profile").
		Joins("LEFT JOIN profiles ON users.id = profiles.user_id").
		Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (r *UserRepository) FindOne(ctx context.Context, userId int) (entity.User, error) {
	var user entity.User
	result := r.Db.WithContext(ctx).Preload("Profile").Joins("LEFT JOIN profiles ON users.id = profiles.user_id").First(&user, userId)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}

func (r *UserRepository) Update(ctx context.Context, user entity.User, userId int) (entity.User, error) {
	// Обновление пользователя
	_ = r.Db.WithContext(ctx).Model(&user).Where("id = ?", userId).Updates(user)

	// Обновление профиля
	_ = r.Db.WithContext(ctx).Model(&user.Profile).Where("user_id = ?", userId).Updates(user.Profile)
	foundedUser, err := r.FindOne(ctx, userId)
	if err != nil {
		return foundedUser, utils.NewAppError(err, "Can't exec", "")
	}

	return foundedUser, nil
}

func (r *UserRepository) Delete(ctx context.Context, userId int) error {
	result := r.Db.WithContext(ctx).Where("id = ?", userId).Delete(&entity.User{})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (r *UserRepository) FindByUsername(ctx context.Context, username string) (entity.User, error) {
	var user entity.User
	result := r.Db.WithContext(ctx).Where("username = ?", username).First(&user)
	if result.Error != nil {
		return entity.User{}, result.Error
	}
	return user, nil
}
