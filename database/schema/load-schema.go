package schema

import (
	"context"
	"database/sql"
	_ "database/sql"
	"os"
	"path/filepath"
)

type PostgresDb struct {
	Db *sql.DB
}

func (b *PostgresDb) LoadSchema(ctx context.Context) error {
	// Чтение содержимого файла SQL
	absPath, err := filepath.Abs("database/schema/schema.sql")
	content, err := os.ReadFile(absPath)
	if err != nil {
		return err
	}

	// Начало транзакции
	tx, err := b.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	// Выполнение SQL-запросов из файла
	queries := string(content)
	_, err = tx.ExecContext(ctx, queries)
	if err != nil {
		// Откат транзакции в случае ошибки
		tx.Rollback()
		return err
	}

	// Фиксация транзакции
	if err := tx.Commit(); err != nil {
		return err
	}

	return nil

}
