package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"math/rand"
	"net/http"
	"quotes_service/internal/repo"
	"quotes_service/models"
	"quotes_service/utils"
	"strconv"
)

type QuoteHandler struct {
	repo *repo.QuoteRepo
}

func NewQuoteHandler(repo *repo.QuoteRepo) *QuoteHandler {
	return &QuoteHandler{repo: repo}
}

func (h *QuoteHandler) CreateQuote(w http.ResponseWriter, r *http.Request) {
	var quote models.Quote
	err := json.NewDecoder(r.Body).Decode(&quote)
	if err != nil {
		http.Error(w, "невалидный json", http.StatusBadRequest)
		return
	}

	err = utils.ValidateQuote(quote)
	if err != nil {
		utils.SendError(w, 400, err.Error())
	}

	quoteID := h.repo.Add(quote)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quoteID)
}

func (h *QuoteHandler) GetQuotes(w http.ResponseWriter, r *http.Request) {
	author := r.URL.Query().Get("author")
	var result []models.Quote
	if author != "" {
		result = h.repo.GetByAuthor(author)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

func (h *QuoteHandler) DeleteQuote(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || !h.repo.Delete(id) {
		utils.SendError(w, 400, "цитата с этим ID не найдена")
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *QuoteHandler) GetRandom(w http.ResponseWriter, r *http.Request) {
	quotes := h.repo.GetAll()
	if len(quotes) == 0 {
		utils.SendError(w, 404, "цитаты не найдены")
		return
	}

	n := rand.Intn(len(quotes))
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(quotes[n])
}
