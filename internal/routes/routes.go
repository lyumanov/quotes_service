package routes

import (
	"github.com/gorilla/mux"
	"quotes_service/internal/handler"
)

func RegisterRoute(h *handler.QuoteHandler) {
	r := mux.NewRouter()

	r.HandleFunc("/quotes", h.CreateQuote).Methods("POST")

}
