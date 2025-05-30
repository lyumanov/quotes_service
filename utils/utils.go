package utils

import (
	"encoding/json"
	"errors"
	"net/http"
	"quotes_service/models"
)

type ErrorResponse struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	} `json:"error"`
}

func ValidateQuote(quote models.Quote) error {
	if quote.Author == "" || quote.Quote == "" {
		return errors.New("автор и/или цитата не заполнены")
	}
	return nil
}

func SendError(w http.ResponseWriter, code int, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	errorResponse := ErrorResponse{}
	errorResponse.Error.Code = code
	errorResponse.Error.Message = message

	json.NewEncoder(w).Encode(errorResponse)
}
