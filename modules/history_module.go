package modules

import (
	"errors"
	"koda-b6-weekly1-go/models"
	"time"
)

type HistoryModule struct {
	Transactions []models.Transaction
	lastID       int
}

func NewHistoryModule() *HistoryModule {
	return &HistoryModule{
		Transactions: []models.Transaction{},
		lastID:       0,
	}
}

func (h *HistoryModule) AddTransaction(items []models.CartItem, total int) error {
	if len(items) == 0 {
		return errors.New("keranjang kosong, tidak bisa checkout")
	}

	h.lastID++

	newItems := []models.CartItem{}
	for i := 0; i < len(items); i++ {
		newItems = append(newItems, items[i])
	}

	transaction := models.Transaction{
		ID:        h.lastID,
		Items:     newItems,
		Total:     total,
		CreatedAt: time.Now(),
	}

	h.Transactions = append(h.Transactions, transaction)
	return nil
}

func (h *HistoryModule) GetAll() []models.Transaction {
	return h.Transactions
}

func (h *HistoryModule) IsEmpty() bool {
	return len(h.Transactions) == 0
}