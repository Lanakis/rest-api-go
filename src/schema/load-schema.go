package schema

import (
	"context"
	_ "database/sql"
)

func (b *PostgresDb) LoadSchema(ctx context.Context) error {
	err := b.UserSchema(ctx)
	if err != nil {
		return err
	}

	err = b.ProfileSchema(ctx)
	if err != nil {
		return err
	}

	err = b.UserData(ctx)
	if err != nil {
		return err
	}

	return nil
}
