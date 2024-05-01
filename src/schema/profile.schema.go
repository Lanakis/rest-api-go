package schema

import "context"

func (b *PostgresDb) ProfileSchema(ctx context.Context) error {

	tx, err := b.Db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()

	SQL := `
		CREATE TABLE IF NOT EXISTS profiles (
			id SERIAL PRIMARY KEY,
			first_name VARCHAR(20) NOT NULL,
		    middle_name VARCHAR(20),
		    last_name VARCHAR(20),
		    age INT,
		    head BOOLEAN,
		    user_id INT,
			created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
			updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
		)
	`

	_, err = tx.ExecContext(ctx, SQL)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}
