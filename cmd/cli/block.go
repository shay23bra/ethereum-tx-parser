package cli

import (
	"ethereum-tx-parser/internal/parser"
	"fmt"

	"github.com/spf13/cobra"
)

var blockCmd = &cobra.Command{
	Use:   "block",
	Short: "Get the current block number",
	Run: func(cmd *cobra.Command, args []string) {
		ethParser := parser.NewEthereumParser("https://ethereum-rpc.publicnode.com")
		block := ethParser.GetCurrentBlock()
		fmt.Printf("Current block: %d\n", block)
	},
}

func init() {
	RootCmd.AddCommand(blockCmd)
}
