package cli

import (
	"ethereum-tx-parser/internal/parser"
	"fmt"

	"github.com/spf13/cobra"
)

var subscribeCmd = &cobra.Command{
	Use:   "subscribe [address]",
	Short: "Subscribe to an Ethereum address",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		address := args[0]
		ethParser := parser.NewEthereumParser("https://ethereum-rpc.publicnode.com")
		success := ethParser.Subscribe(address)
		if success {
			fmt.Printf("Subscribed to address: %s\n", address)
		} else {
			fmt.Println("Failed to subscribe to address")
		}
	},
}
