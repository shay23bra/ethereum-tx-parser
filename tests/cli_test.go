package tests

import (
	"bytes"
	"os/exec"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCLI(t *testing.T) {
	// Create a temporary directory for the binary
	tempDir := t.TempDir()
	binaryPath := filepath.Join(tempDir, "ethereum-tx-parser")

	// Build the CLI binary
	buildCmd := exec.Command("go", "build", "-o", binaryPath, "./cmd/main.go")
	var buildOut bytes.Buffer
	buildCmd.Stdout = &buildOut
	buildCmd.Stderr = &buildOut
	err := buildCmd.Run()
	if err != nil {
		t.Fatalf("Failed to build CLI binary: %v\nOutput: %s", err, buildOut.String())
	}

	t.Run("Subscribe - Success", func(t *testing.T) {
		cmd := exec.Command(binaryPath, "-mode", "cli", "subscribe", "0xTestAddress")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		err := cmd.Run()
		assert.NoError(t, err, "CLI command should succeed")
		assert.Contains(t, out.String(), "Subscribed to address", "Output should indicate successful subscription")
	})

	t.Run("Subscribe - Missing Address", func(t *testing.T) {
		cmd := exec.Command(binaryPath, "-mode", "cli", "subscribe")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		err := cmd.Run()
		assert.Error(t, err, "CLI command should fail for missing address")
		assert.Contains(t, out.String(), "accepts 1 arg(s)", "Output should indicate missing address")
	})

	t.Run("GetCurrentBlock - Success", func(t *testing.T) {
		cmd := exec.Command(binaryPath, "-mode", "cli", "block")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		err := cmd.Run()
		assert.NoError(t, err, "CLI command should succeed")
		assert.Contains(t, out.String(), "Current block", "Output should contain current block number")
	})

	t.Run("GetTransactions - No Transactions", func(t *testing.T) {
		cmd := exec.Command(binaryPath, "-mode", "cli", "transactions", "0xNonExistentAddress")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		err := cmd.Run()
		assert.NoError(t, err, "CLI command should succeed")
		assert.Contains(t, out.String(), "No transactions found", "Output should indicate no transactions")
	})

	t.Run("GetTransactions - Invalid Address", func(t *testing.T) {
		cmd := exec.Command(binaryPath, "-mode", "cli", "transactions", "invalid-address")
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		err := cmd.Run()
		assert.NoError(t, err, "CLI command should succeed")
		assert.Contains(t, out.String(), "No transactions found", "Output should indicate no transactions")
	})
}
