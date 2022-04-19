package main

import (
	"errors"
	"net/http"

	"test_service.mjekson.ru/internal/data"
	"test_service.mjekson.ru/internal/validator"
)

func (app *application) addBalanceHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParamWallet(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	wallet, err := app.models.Wallets.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Balance int `json:"balance"`
	}

	// err = app.readJSON(w, r, &input)
	// if err != nil {
	// 	app.badRequestResponse(w, r, err)
	// 	return
	// }
	v := validator.NewValidatorWallet()

	qs := r.URL.Query()

	input.Balance = app.readIntWallet(qs, "amount", 0, v)

	wallet.Balance = wallet.Balance + input.Balance

	if data.ValidateWallet(v, wallet); !v.ValidWallet() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Wallets.Update(wallet)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSONWallet(w, http.StatusOK, envelopeWallet{"wallet": wallet}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}

func (app *application) addTransferHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParamWallet(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	wallet1, err := app.models.Wallets.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	v := validator.NewValidatorWallet()
	qs := r.URL.Query()
	id, err = app.readIDParamTransfer(qs, "user_id_to", v)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	wallet2, err := app.models.Wallets.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		Balance int `json:"balance"`
	}

	input.Balance = app.readIntWallet(qs, "amount", 0, v)

	wallet1.Balance = wallet1.Balance - input.Balance

	if data.ValidateWallet(v, wallet1); !v.ValidWallet() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Wallets.Update(wallet1)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	wallet2.Balance = wallet2.Balance + input.Balance

	if data.ValidateWallet(v, wallet2); !v.ValidWallet() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Wallets.Update(wallet2)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	wallets := make([]data.Wallet, 0)
	wallets = append(wallets, *wallet1)
	wallets = append(wallets, *wallet2)

	err = app.writeJSONWallet(w, http.StatusOK, envelopeWallet{"wallets": wallets}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}

}
