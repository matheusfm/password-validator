package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/matheusfm/password-validator/pkg/validator"
)

func PasswordValidator(w http.ResponseWriter, r *http.Request) {
	var body *ValidationRequest
	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	pwdValidator := validator.New(validator.MinLength(9),
		validator.MinDigits(1),
		validator.MinLowercase(1),
		validator.MinUppercase(1),
		validator.MinSpecialChar(1))
	err := pwdValidator.IsValid(body.Password)
	response := &ValidationResponse{Valid: err == nil}
	if err != nil {
		response.Message = err.Error()
	}

	bytes, err := json.Marshal(response)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write(bytes)
	return
}
