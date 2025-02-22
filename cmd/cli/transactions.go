package cli

import (
	"ethereum-tx-parser/internal/parser"
	"fmt"

	"github.com/spf13/cobra"
)

var transactionsCmd = &cobra.Command{
	Use:   "transactions [address]",
	Short: "Get transactions for a subscribed address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]
		ethParser := parser.NewEthereumParser("https://ethereum-rpc.publicnode.com")
		transactions := ethParser.GetTransactions(address)
		if len(transactions) == 0 {
			fmt.Println("No transactions found for the address")
		} else {
			fmt.Printf("Transactions for address %s:\n", address)
			for _, tx := range transactions {
				fmt.Printf("Hash: %s, From: %s, To: %s, Value: %s, Block: %s\n", tx.Hash, tx.From, tx.To, tx.Value, tx.Block)
			}
		}
	},
}
