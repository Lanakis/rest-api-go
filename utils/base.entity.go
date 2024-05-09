package utils

import "time"

type BaseEntity struct {
	Id        int       `gorm:"primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime"`
	UpdatedAt time.Time `gorm:"autoUpdateTime"`
}

// BeforeInsert вызывается перед вставкой новой записи в базу данных
func (b *BaseEntity) BeforeInsert() error {
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
	return nil
}

// BeforeUpdate вызывается перед обновлением существующей записи в базе данных
func (b *BaseEntity) BeforeUpdate() error {
	b.UpdatedAt = time.Now()
	return nil
}
