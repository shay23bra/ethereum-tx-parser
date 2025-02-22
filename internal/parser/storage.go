package parser

import (
	"ethereum-tx-parser/internal/models"
	"sync"
)

type Storage struct {
	subscribedAddresses map[string]bool
	transactions        map[string][]models.Transaction
	mu                  sync.RWMutex
}

func NewStorage() *Storage {
	return &Storage{
		subscribedAddresses: make(map[string]bool),
		transactions:        make(map[string][]models.Transaction),
	}
}

func (s *Storage) Subscribe(address string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.subscribedAddresses[address] = true
}

func (s *Storage) IsSubscribed(address string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.subscribedAddresses[address]
}

func (s *Storage) AddTransaction(address string, tx models.Transaction) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.transactions[address] = append(s.transactions[address], tx)
}

func (s *Storage) GetTransactions(address string) []models.Transaction {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.transactions[address]
}
