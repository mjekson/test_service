package validator

type ValidatorWallet struct {
	Errors map[string]string
}

func NewValidatorWallet() *ValidatorWallet {
	return &ValidatorWallet{Errors: make(map[string]string)}
}

func (v *ValidatorWallet) ValidWallet() bool {
	return len(v.Errors) == 0
}

func (v *ValidatorWallet) AddErrorWallet(key, message string) {
	if _, exists := v.Errors[key]; !exists {
		v.Errors[key] = message
	}
}

// Check adds an error message to the map only if a validation check is not 'ok'
func (v *ValidatorWallet) CheckWallet(ok bool, key, message string) {
	if !ok {
		v.AddErrorWallet(key, message)
	}
}
