package repo

import (
	"quotes_service/models"
	"sync"
)

type QuoteRepo struct {
	sync.Mutex
	quotes map[int]models.Quote
	lastID int
}

func NewQuoteRepo() *QuoteRepo {
	return &QuoteRepo{
		quotes: make(map[int]models.Quote),
	}
}

func (r *QuoteRepo) Add(quote models.Quote) int {
	r.Lock()
	defer r.Unlock()
	r.lastID++
	quote.ID = r.lastID
	r.quotes[quote.ID] = quote
	return r.lastID
}

func (r *QuoteRepo) GetAll() []models.Quote {
	r.Lock()
	defer r.Unlock()
	var result []models.Quote
	for _, q := range r.quotes {
		result = append(result, q)
	}
	return result
}

func (r *QuoteRepo) GetByAuthor(author string) []models.Quote {
	r.Lock()
	defer r.Unlock()
	var result []models.Quote
	for _, q := range r.quotes {
		if q.Author == author {
			result = append(result, q)
		}
	}
	return result
}

func (r *QuoteRepo) Delete(id int) bool {
	r.Lock()
	defer r.Unlock()
	if _, ok := r.quotes[id]; ok {
		delete(r.quotes, id)
		return true
	}
	return false
}
