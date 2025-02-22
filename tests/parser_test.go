package tests

import (
	"ethereum-tx-parser/internal/parser"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEthereumParser(t *testing.T) {
	rpcURL := "https://ethereum-rpc.publicnode.com"
	ethParser := parser.NewEthereumParser(rpcURL)

	t.Run("GetCurrentBlock - Success", func(t *testing.T) {
		block := ethParser.GetCurrentBlock()
		assert.Greater(t, block, 0, "Current block should be greater than 0")
	})

	t.Run("GetCurrentBlock - Invalid RPC URL", func(t *testing.T) {
		invalidParser := parser.NewEthereumParser("https://invalid-rpc-url")
		block := invalidParser.GetCurrentBlock()
		assert.Equal(t, 0, block, "Expected 0 for invalid RPC URL")
	})

	t.Run("Subscribe - Success", func(t *testing.T) {
		address := "0xTestAddress"
		success := ethParser.Subscribe(address)
		assert.True(t, success, "Subscription should succeed")
	})

	t.Run("Subscribe - Empty Address", func(t *testing.T) {
		success := ethParser.Subscribe("")
		assert.False(t, success, "Subscription should fail for empty address")
	})

	t.Run("GetTransactions - No Transactions", func(t *testing.T) {
		address := "0xNonExistentAddress"
		transactions := ethParser.GetTransactions(address)
		assert.Equal(t, 0, len(transactions), "Expected 0 transactions for non-existent address")
	})

	t.Run("GetTransactions - Invalid Address", func(t *testing.T) {
		transactions := ethParser.GetTransactions("invalid-address")
		assert.Equal(t, 0, len(transactions), "Expected 0 transactions for invalid address")
	})
}
