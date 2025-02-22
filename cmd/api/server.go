package api

import (
	"encoding/json"
	"ethereum-tx-parser/internal/models"
	"ethereum-tx-parser/internal/parser"
	"net/http"
)

type Server struct {
	parser *parser.EthereumParser
}

func NewServer(rpcURL string) *Server {
	return &Server{
		parser: parser.NewEthereumParser(rpcURL),
	}
}

func (s *Server) Start(addr string) error {
	http.HandleFunc("/subscribe", s.HandleSubscribe)
	http.HandleFunc("/block", s.HandleGetCurrentBlock)
	http.HandleFunc("/transactions", s.HandleGetTransactions)
	return http.ListenAndServe(addr, nil)
}

func (s *Server) HandleSubscribe(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "address is required", http.StatusBadRequest)
		return
	}

	success := s.parser.Subscribe(address)
	if !success {
		http.Error(w, "failed to subscribe", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"status": "subscribed", "address": address})
}

func (s *Server) HandleGetCurrentBlock(w http.ResponseWriter, r *http.Request) {
	block := s.parser.GetCurrentBlock()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]int{"current_block": block})
}

func (s *Server) HandleGetTransactions(w http.ResponseWriter, r *http.Request) {
	address := r.URL.Query().Get("address")
	if address == "" {
		http.Error(w, "address is required", http.StatusBadRequest)
		return
	}

	transactions := s.parser.GetTransactions(address)
	if transactions == nil {
		transactions = []models.Transaction{} // Return an empty array instead of nil
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(transactions)
}
