package schema

import (
	"context"
)

func (b *PostgresDb) UserSchema(ctx context.Context) error {

	tx, err := b.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	SQL := `
		CREATE TABLE IF NOT EXISTS users (
			id SERIAL PRIMARY KEY,
			username VARCHAR(50) NOT NULL,
			password VARCHAR(100) NOT NULL,
			role VARCHAR(20) NOT NULL,
			refresh_token VARCHAR (100),
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	_, err = tx.ExecContext(ctx, SQL)
	if err != nil {
		return err
	}

	// Фиксируем транзакцию
	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
