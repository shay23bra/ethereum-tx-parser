package rpc

import (
	"bytes"
	"encoding/json"
	"net/http"
)

type RPCClient struct {
	URL string
}

func NewRPCClient(url string) *RPCClient {
	return &RPCClient{URL: url}
}

func (c *RPCClient) Call(method string, params []interface{}) (map[string]interface{}, error) {
	request := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  method,
		"params":  params,
		"id":      1,
	}

	requestBody, _ := json.Marshal(request)
	resp, err := http.Post(c.URL, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var result map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result, nil
}
