package utils

import "time"

type BaseEntity struct {
	Id        int       `db:"id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
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
