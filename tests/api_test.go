package tests

import (
	"ethereum-tx-parser/cmd/api"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPI(t *testing.T) {
	server := api.NewServer("https://ethereum-rpc.publicnode.com")

	t.Run("Subscribe - Success", func(t *testing.T) { // address of Binance Hot Wallet
		req := httptest.NewRequest("GET", "/subscribe?address=0xF977814e90dA44bFA03b6295A0616a897441aceC", nil)
		w := httptest.NewRecorder()
		server.HandleSubscribe(w, req)

		assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 for successful subscription")
		assert.Contains(t, w.Body.String(), "subscribed", "Response should contain 'subscribed'")
	})

	t.Run("Subscribe - Missing Address", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/subscribe", nil)
		w := httptest.NewRecorder()
		server.HandleSubscribe(w, req)

		assert.Equal(t, http.StatusBadRequest, w.Code, "Expected HTTP 400 for missing address")
		assert.Contains(t, w.Body.String(), "address is required", "Response should contain error message")
	})

	t.Run("GetCurrentBlock - Success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/block", nil)
		w := httptest.NewRecorder()
		server.HandleGetCurrentBlock(w, req)

		assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 for successful block fetch")
		assert.Contains(t, w.Body.String(), "current_block", "Response should contain 'current_block'")
	})

	t.Run("GetTransactions - Invalid Address", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/transactions?address=0xNonExistentAddress", nil)
		w := httptest.NewRecorder()
		server.HandleGetTransactions(w, req)

		assert.Equal(t, http.StatusNotFound, w.Code, "Expected HTTP 404 for not found address")
		assert.Contains(t, w.Body.String(), "address is not subscribed", strings.TrimSpace(w.Body.String()), "Response should indicate address is not subscribed to")
	})
}
