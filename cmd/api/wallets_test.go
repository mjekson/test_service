package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"test_service.mjekson.ru/internal/data"
	"test_service.mjekson.ru/internal/validator"
)

func TestWallets_Validate(t *testing.T) {
	testCases := []struct {
		name    string
		v       func() *validator.ValidatorWallet
		w       func() *data.Wallet
		isValid bool
	}{
		{
			name: "Valid",
			v: func() *validator.ValidatorWallet {
				return validator.NewValidatorWallet()
			},
			w: func() *data.Wallet {
				return TestWallets(t)
			},
			isValid: true,
		},
		{
			name: "invalid balance",
			v: func() *validator.ValidatorWallet {
				return validator.NewValidatorWallet()
			},
			w: func() *data.Wallet {
				u := TestWallets(t)
				u.Balance = -100
				return u
			},
			isValid: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.isValid {
				data.ValidateWallet(tc.v(), tc.w())
				assert.NoError(t, nil)
			} else {
				data.ValidateWallet(tc.v(), tc.w())
				assert.Error(t, errors.New("Error ValidWallet"))
			}
		})
	}
}
