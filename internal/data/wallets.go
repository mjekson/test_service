package data

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"test_service.mjekson.ru/internal/validator"
)

type Wallet struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Balance int    `json:"balance"`
	Version int32  `json:"version"`
}

type WalletModel struct {
	DB *sql.DB
}

func (m WalletModel) Get(id int64) (*Wallet, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	//pg_sleep(10),
	query := `
		SELECT id, name, balance, version 
		FROM users
		WHERE id = $1`

	var wallet Wallet

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	//&[]byte{},
	err := m.DB.QueryRowContext(ctx, query, id).Scan(
		&wallet.ID,
		&wallet.Name,
		&wallet.Balance,
		&wallet.Version,
	)

	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}
	return &wallet, nil
}

func (m WalletModel) Update(wallet *Wallet) error {
	query := `
	UPDATE users
	SET balance = $1, version = version + 1 
	WHERE id = $2 AND version = $3
	RETURNING version`

	args := []interface{}{
		wallet.Balance,
		wallet.ID,
		wallet.Version,
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	err := m.DB.QueryRowContext(ctx, query, args...).Scan(&wallet.Version)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return ErrEditConflict
		default:
			return err
		}
	}

	return nil

}

func ValidateWallet(v *validator.ValidatorWallet, input *Wallet) {
	//v.Check(input.Balance != 0, "balance", "must be provided")
	v.CheckWallet(input.Balance >= 0, "balance", "must be a positive integer")
}
