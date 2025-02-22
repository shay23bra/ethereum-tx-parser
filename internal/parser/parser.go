package parser

import (
	"ethereum-tx-parser/internal/rpc"
	"ethereum-tx-parser/models"
	"fmt"
	"strconv"
	"time"
)

type Parser interface {
	GetCurrentBlock() int
	Subscribe(address string) bool
	GetTransactions(address string) []models.Transaction
	Listen()
}

type EthereumParser struct {
	rpcClient *rpc.RPCClient
	storage   *Storage
}

func NewEthereumParser(rpcURL string) *EthereumParser {
	rpcClient := rpc.NewRPCClient(rpcURL)
	parser := &EthereumParser{
		rpcClient: rpcClient,
	}
	block := parser.GetCurrentBlock()
	parser.storage = NewStorage(block)
	parser.Listen()
	return parser
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

func (p *EthereumParser) GetLastBlock() int {
	return p.storage.lastBlock
}

func (p *EthereumParser) Subscribe(address string) bool {
	if address == "" {
		return false
	}
	p.storage.Subscribe(address)
	return true
}

func (p *EthereumParser) GetTransactions(address string) []models.Transaction {
	return p.storage.GetTransactions(address)
}

func (p *EthereumParser) IsSubscribed(address string) bool {
	return p.storage.IsSubscribed(address)
}

// Listen starts a goroutine that continuously polls for new blocks.
// For each new block, it fetches the transactions and if a transaction involves
// a subscribed address (either as the sender or recipient), it is added to storage.
func (p *EthereumParser) Listen() {
	go func() {
		lastBlock := p.GetLastBlock()
		for {
			currentBlock := p.GetCurrentBlock()
			if currentBlock > lastBlock {
				for blockNum := lastBlock + 1; blockNum <= currentBlock; blockNum++ {
					blockHex := fmt.Sprintf("0x%x", blockNum)
					// Fetch block details along with full transaction objects.
					result, err := p.rpcClient.Call("eth_getBlockByNumber", []interface{}{blockHex, true})
					if err != nil {
						fmt.Println("Error fetching block", blockNum, ":", err)
						continue
					}
					blockData, ok := result["result"].(map[string]interface{})
					if !ok {
						fmt.Println("Invalid block data for block", blockNum)
						continue
					}
					txs, ok := blockData["transactions"].([]interface{})
					if !ok {
						fmt.Println("No transactions found in block", blockNum)
						continue
					}
					for _, rawTx := range txs {
						txMap, ok := rawTx.(map[string]interface{})
						if !ok {
							continue
						}
						tx := convertTx(txMap)
						// If the sender address is subscribed, add the transaction.
						if p.storage.IsSubscribed(tx.From) {
							p.storage.AddTransaction(tx.From, tx)
						}
						// If the receiver address is subscribed, add the transaction.
						if tx.To != "" && p.storage.IsSubscribed(tx.To) {
							p.storage.AddTransaction(tx.To, tx)
						}
					}
				}
				lastBlock = currentBlock
			}
			time.Sleep(2 * time.Second) // Adjust the polling interval as needed.
		}
	}()
}

// convertTx converts the raw transaction map into a models.Transaction.
func convertTx(txMap map[string]interface{}) models.Transaction {
	tx := models.Transaction{}
	if hash, ok := txMap["hash"].(string); ok {
		tx.Hash = hash
	}
	if from, ok := txMap["from"].(string); ok {
		tx.From = from
	}
	if to, ok := txMap["to"].(string); ok {
		tx.To = to
	}
	if value, ok := txMap["value"].(string); ok {
		tx.Value = value
	}
	if block, ok := txMap["blockNumber"].(string); ok {
		tx.Block = block
	}
	return tx
}
