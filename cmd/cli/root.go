package cli

import (
	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "ethereum-tx-parser",
	Short: "A CLI for parsing Ethereum transactions",
	Long:  `A CLI tool to interact with the Ethereum blockchain and parse transactions for subscribed addresses.`,
}

func init() {
	RootCmd.AddCommand(subscribeCmd)
	RootCmd.AddCommand(blockCmd)
	RootCmd.AddCommand(transactionsCmd)
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		panic(err)
	}
}
