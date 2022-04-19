package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"test_service.mjekson.ru/internal/validator"
)

type envelopeWallet map[string]interface{}

func (app *application) readIDParamWallet(r *http.Request) (int64, error) {
	params := httprouter.ParamsFromContext(r.Context())
	id, err := strconv.ParseInt(params.ByName("id"), 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) readIDParamTransfer(qs url.Values, key string, v *validator.ValidatorWallet) (int64, error) {
	s := qs.Get(key)
	if s == "" {
		v.AddErrorWallet(key, "must be user_id_to transfer parameter")
		return 0, nil
	}
	id, err := strconv.ParseInt(s, 10, 64)
	if err != nil || id < 1 {
		return 0, errors.New("invalid id parameter")
	}
	return id, nil
}

func (app *application) readIntWallet(qs url.Values, key string, defaultValue int, v *validator.ValidatorWallet) int {
	s := qs.Get(key)

	if s == "" {
		return defaultValue
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		v.AddErrorWallet(key, "must be an integer value")
		return defaultValue
	}

	return i
}

// func (app *application) readKeyWallet(qs url.Values, key string) bool {
// 	s := qs.Get(key)

// 	// if s == "" {
// 	// 	return false
// 	// }

// 	return s != ""
// }

func (app *application) writeJSONWallet(w http.ResponseWriter, status int, data envelopeWallet, headers http.Header) error {

	js, err := json.MarshalIndent(data, "", "\t")
	if err != nil {
		return err
	}
	// Append a newline to make it easier to view in terminal applications.
	js = append(js, '\n')

	for key, value := range headers {
		w.Header()[key] = value
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}
