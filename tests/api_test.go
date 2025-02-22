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

	t.Run("Subscribe - Success", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/subscribe?address=0xTestAddress", nil)
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

	t.Run("GetTransactions - No Transactions", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/transactions?address=0xNonExistentAddress", nil)
		w := httptest.NewRecorder()
		server.HandleGetTransactions(w, req)

		assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 for successful transaction fetch")
		assert.Equal(t, "[]", strings.TrimSpace(w.Body.String()), "Expected empty array for no transactions")
	})

	t.Run("GetTransactions - Invalid Address", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/transactions?address=invalid-address", nil)
		w := httptest.NewRecorder()
		server.HandleGetTransactions(w, req)

		assert.Equal(t, http.StatusOK, w.Code, "Expected HTTP 200 for invalid address")
		assert.Equal(t, "[]", strings.TrimSpace(w.Body.String()), "Expected empty array for invalid address")
	})
}
