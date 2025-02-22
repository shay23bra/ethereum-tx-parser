package main

import (
	"ethereum-tx-parser/cmd/api"
	"ethereum-tx-parser/cmd/cli"
	"flag"
	"os"
)

const RpcURL = "https://ethereum-rpc.publicnode.com"

func main() {
	mode := flag.String("mode", "cli", "Run mode: cli or api")
	flag.Parse()

	switch *mode {
	case "cli":
		os.Args = append([]string{os.Args[0]}, flag.Args()...)
		cli.Execute()
	case "api":
		server := api.NewServer(RpcURL)
		if err := server.Start(":8080"); err != nil {
			panic(err)
		}
	default:
		panic("invalid mode. use 'cli' or 'api'")
	}
}
