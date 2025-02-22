# Ethereum Transaction Parser

A Go-based Ethereum blockchain parser that allows users to subscribe to Ethereum addresses and retrieve transactions. The project supports both CLI and HTTP API modes for interacting with the parser.

- **Subscribe to Ethereum Addresses**: Add addresses to monitor for transactions.
- **Fetch Current Block Number**: Get the latest block number from the Ethereum blockchain.
- **Retrieve Transactions**: Get all inbound or outbound transactions for a subscribed address.
- **CLI and HTTP API**: Interact with the parser via command-line interface or HTTP API.

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/yourusername/ethereum-tx-parser.git
2. Navigate to the project directory:
    ```bash
    cd ethereum-tx-parser
3. Build the application:
    ```bash
    go build -o ethereum-tx-parser cmd/main.go

## Usage

### CLI Mode

Run the application in CLI mode to interact with the parser via the command line.

**Subscribe to an Ethereum Address:**
```bash
./ethereum-tx-parser -mode cli subscribe 0xYourEthereumAddress
```
**Get the Current Block Number:**
```bash
./ethereum-tx-parser -mode cli block
```
**Get Transactions for a Subscribed Address:**
```bash
./ethereum-tx-parser -mode cli transactions 0xYourEthereumAddress
```


### HTTP API Mode
Run the application in HTTP API mode to expose the parser functionality via REST endpoints.

1. Start the HTTP server:

    ```bash
    ./ethereum-tx-parser -mode api
2. Use the following endpoints to interact with the API:

**Subscribe to an Ethereum Address**
```bash
GET /subscribe?address=0xYourEthereumAddress
```
**Get the Current Block Number**
```bash
GET /block
```
**Get Transactions for a Subscribed Address**
```bash
GET /transactions?address=0xYourEthereumAddress
```
**Example Requests**
```bash
curl http://localhost:8080/subscribe?address=0xYourEthereumAddress
curl http://localhost:8080/block
curl http://localhost:8080/transactions?address=0xYourEthereumAddress
```

## Testing

### Automated Testing

This project uses GitHub Actions to run automated tests on every push to the `main` branch. The tests ensure that the parser, API endpoints, and CLI functionality work as expected.

### Running Tests Locally

To run the tests locally, use the following command:
```bash
go test ./tests/...
```

## Project Structure
```
ethereum-tx-parser/
├── cmd/
│   ├── main.go          # Entry point for the application
│   ├── cli/             # CLI commands
│   └── api/             # HTTP API server
├── internal/
│   ├── parser/          # Parser implementation
│   ├── rpc/             # Ethereum JSON-RPC client
│   └── models/          # Data models
├── tests/               # Automated tests
│   ├── parser_test.go   # Unit tests for the parser
│   ├── cli_test.go      # CLI integration tests
│   └── api_test.go      # HTTP API tests
├── README.md            # Project documentation
└── go.mod               # Go module file
```

## Dependencies
- Go Modules: Used for dependency management.
- Cobra: A library for building powerful CLI applications.
- net/http: Used for creating the HTTP API server.

## Limitations
- In-Memory Storage: The project uses in-memory storage for simplicity. This can be extended to support databases like PostgreSQL or Redis.
- No Authentication: The HTTP API does not include authentication. This can be added for production use.
- Single RPC Endpoint: The project uses a single Ethereum RPC endpoint. Load balancing or failover mechanisms can be added for robustness.

## Future Enhancements
- Database Integration: Replace in-memory storage with a persistent database.
- Authentication: Add API key or JWT-based authentication for the HTTP API.
- WebSocket Support: Add WebSocket support for real-time transaction notifications.
- Dockerization: Create a Docker image for easy deployment.

## References
-   [Ethereum JSON-RPC Documentation](https://ethereum.org/en/developers/docs/apis/json-rpc/)
-   [Cobra CLI Library](https://github.com/spf13/cobra)
-   [Go net/http Package](https://pkg.go.dev/net/http)

## License
This project is licensed under the terms of the MIT license.