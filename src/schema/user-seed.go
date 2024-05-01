package schema

import (
	"authorization/src/utils"
	"context"
)

func (b *PostgresDb) UserData(ctx context.Context) error {
	tx, err := b.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// Проверяем существует ли п
	//ользователь с заданным именем пользователя
	var count int
	var userId int
	err = tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM users WHERE username = $1", "admin").Scan(&count)
	if err != nil {
		return err
	}
	if count == 0 {
		// Если пользователь не существует, выполняем вставку данных
		SQL := `INSERT INTO users (username, password, role) VALUES ($1, $2, $3) RETURNING id`
		password := utils.HashPassword("123456")
		err = tx.QueryRowContext(ctx, SQL, "admin", password, "admin").Scan(&userId)

		if err != nil {
			return err
		}
	}
	errProf := tx.QueryRowContext(ctx, "SELECT COUNT(*) FROM profiles WHERE first_name = $1 AND middle_name = $2 AND last_name = $3 ", "Администратор", "Тест", "Создание").Scan(&count)
	if errProf != nil {
		return err
	}
	if count == 0 {
		// Если пользователь не существует, выполняем вставку данных
		SQL := `INSERT INTO profiles (first_name, middle_name, last_name, age, head, user_id) VALUES ($1, $2, $3,$4,$5,$6)`

		_, err = tx.ExecContext(ctx, SQL, "Администратор", "Тест", "Создание", 25, true, userId)
		if err != nil {
			return err
		}

	}
	err = tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
