package parser

import (
	"ethereum-tx-parser/internal/models"
	"ethereum-tx-parser/internal/rpc"
	"fmt"
	"strconv"
)

type Parser interface {
	GetCurrentBlock() int
	Subscribe(address string) bool
	GetTransactions(address string) []models.Transaction
}

type EthereumParser struct {
	rpcClient *rpc.RPCClient
	storage   *Storage
}

func NewEthereumParser(rpcURL string) *EthereumParser {
	return &EthereumParser{
		rpcClient: rpc.NewRPCClient(rpcURL),
		storage:   NewStorage(),
	}
}

func (p *EthereumParser) GetCurrentBlock() int {
	result, err := p.rpcClient.Call("eth_blockNumber", []interface{}{})
	if err != nil {
		fmt.Println("Error fetching current block:", err)
		return 0
	}

	blockHex := result["result"].(string)
	block, _ := strconv.ParseInt(blockHex[2:], 16, 64)
	return int(block)
}

func (p *EthereumParser) Subscribe(address string) bool {
	p.storage.Subscribe(address)
	return true
}

func (p *EthereumParser) GetTransactions(address string) []models.Transaction {
	return p.storage.GetTransactions(address)
}
