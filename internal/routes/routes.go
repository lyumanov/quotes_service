package routes

import (
	"github.com/gorilla/mux"
	"net/http"
	"quotes_service/internal/handler"
)

func RegisterRoute(h *handler.QuoteHandler) http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/quotes", h.CreateQuote).Methods("POST")
	r.HandleFunc("/quotes", h.GetQuotesByAuthor).Methods("GET").Queries("author", "{author}")
	r.HandleFunc("/quotes", h.GetAllQuotes).Methods("GET")
	r.HandleFunc("/quotes/random", h.GetRandom).Methods("GET")
	r.HandleFunc("/quotes/{id:[0-9]+}", h.DeleteQuote).Methods("DELETE")

	return r

}
