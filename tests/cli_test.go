package tests

import (
	"bytes"
	"os"
	"os/exec"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCLI(t *testing.T) {
	// Build the CLI binary
	buildCmd := exec.Command("go", "build", "-o", "ethereum-tx-parser", "cmd/main.go")
	err := buildCmd.Run()
	assert.NoError(t, err, "Failed to build CLI binary")

	t.Run("Subscribe - Success", func(t *testing.T) {
		cmd := exec.Command("./ethereum-tx-parser", "-mode", "cli", "subscribe", "0xTestAddress")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		assert.NoError(t, err, "CLI command should succeed")
		assert.Contains(t, out.String(), "Subscribed to address", "Output should indicate successful subscription")
	})

	t.Run("Subscribe - Missing Address", func(t *testing.T) {
		cmd := exec.Command("./ethereum-tx-parser", "-mode", "cli", "subscribe")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		err := cmd.Run()
		assert.Error(t, err, "CLI command should fail for missing address")
		assert.Contains(t, out.String(), "accepts 1 arg(s)", "Output should indicate missing address")
	})

	t.Run("GetCurrentBlock - Success", func(t *testing.T) {
		cmd := exec.Command("./ethereum-tx-parser", "-mode", "cli", "block")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		assert.NoError(t, err, "CLI command should succeed")
		assert.Contains(t, out.String(), "Current block", "Output should contain current block number")
	})

	t.Run("GetTransactions - No Transactions", func(t *testing.T) {
		cmd := exec.Command("./ethereum-tx-parser", "-mode", "cli", "transactions", "0xNonExistentAddress")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		assert.NoError(t, err, "CLI command should succeed")
		assert.Contains(t, out.String(), "No transactions found", "Output should indicate no transactions")
	})

	t.Run("GetTransactions - Invalid Address", func(t *testing.T) {
		cmd := exec.Command("./ethereum-tx-parser", "-mode", "cli", "transactions", "invalid-address")
		var out bytes.Buffer
		cmd.Stdout = &out
		err := cmd.Run()
		assert.NoError(t, err, "CLI command should succeed")
		assert.Contains(t, out.String(), "No transactions found", "Output should indicate no transactions")
	})

	// Clean up the binary after tests
	t.Cleanup(func() {
		os.Remove("ethereum-tx-parser")
	})
}
