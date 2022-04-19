package main

import (
	"testing"

	"test_service.mjekson.ru/internal/data"
)

func TestWallets(t *testing.T) *data.Wallet {
	return &data.Wallet{
		ID:      1,
		Name:    "Evgeny",
		Balance: 100,
		Version: 1,
	}
}

func TestWalletModel(t *testing.T) *data.WalletModel {
	return &data.WalletModel{
		DB: nil,
	}
}
