package respositoryutil

import (
	"context"
	"project-management-demo-backend/ent"

	"github.com/pkg/errors"
)

// WithTx queries database safely with transactions.
func WithTx(ctx context.Context, client *ent.Client, fn func(tx *ent.Tx) (res interface{}, error error)) (res interface{}, error error) {
	tx, err := client.Tx(ctx)
	if err != nil {
		return nil, err
	}

	defer func() {
		if v := recover(); v != nil {
			_ = tx.Rollback()
		}
	}()

	res, err = fn(tx)

	if err != nil {
		if rerr := tx.Rollback(); rerr != nil {
			err = errors.Wrapf(err, "rolling back transaction: %v", rerr)
		}
		return nil, err
	}
	if err = tx.Commit(); err != nil {
		return nil, errors.Wrapf(err, "committing transaction: %v", err)
	}

	return res, nil
}
