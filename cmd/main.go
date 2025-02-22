package main

import (
	"ethereum-tx-parser/cmd/api"
	"ethereum-tx-parser/cmd/cli"
	"flag"
)

func main() {
	mode := flag.String("mode", "cli", "Run mode: cli or api")
	flag.Parse()

	rpcURL := "https://ethereum-rpc.publicnode.com"

	switch *mode {
	case "cli":
		cli.Execute()
	case "api":
		server := api.NewServer(rpcURL)
		if err := server.Start(":8080"); err != nil {
			panic(err)
		}
	default:
		panic("invalid mode. use 'cli' or 'api'")
	}
}
