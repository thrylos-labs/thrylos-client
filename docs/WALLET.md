# Thrylos Light Client

The Thrylos Light Client is a lightweight node implementation that allows easy interaction with the Thrylos blockchain without running a full node. 

This is particularly useful for wallet development and other applications that need blockchain interaction.

The Thrylos light client serves several important purposes:

Easy Network Access:
Provides a lightweight way to query blockchain data

Acts as a proxy between users and full nodes

Resource Efficiency:
Doesn't store the full blockchain
Uses minimal storage and memory

## Getting Started

### Installation
```bash
# Clone the repository
git clone https://github.com/thrylos-labs/thrylos-client
cd thrylos-client

# Run the light client
go run main.go -address :8545 -seed localhost:8546
```

### Basic Usage for Wallet Development

Here are the main JSON-RPC endpoints you'll need for wallet development:

#### 1. Check Balance
```bash
curl -X POST http://localhost:8545 \
-H "Content-Type: application/json" \
-H "Origin: http://localhost:8545" \
-d '{
    "jsonrpc":"2.0",
    "method":"getBalance",
    "params":["ADDRESS_HERE"],
    "id":1
}'

# Example Response:
{
    "jsonrpc": "2.0",
    "result": {
        "balance": 700000000,
        "balanceThrylos": 70
    },
    "id": 1
}
```

#### 2. Send Transaction
```bash
curl -X POST http://localhost:8545 \
-H "Content-Type: application/json" \
-H "Origin: http://localhost:8545" \
-d '{
    "jsonrpc":"2.0",
    "method":"submitTransaction",
    "params":[{
        "from": "SENDER_ADDRESS",
        "to": "RECIPIENT_ADDRESS",
        "amount": "100"
    }],
    "id":1
}'
```

#### 3. Get Blockchain Info
```bash
curl -X POST http://localhost:8545 \
-H "Content-Type: application/json" \
-H "Origin: http://localhost:8545" \
-d '{
    "jsonrpc":"2.0",
    "method":"getBlockchainInfo",
    "params":[],
    "id":1
}'

# Example Response:
{
    "jsonrpc": "2.0",
    "result": {
        "chainId": "0x5",
        "height": 1,
        "isSyncing": false,
        "nodeCount": 0,
        "nodeVersion": "1.0.0"
    },
    "id": 1
}
```

#### 4. Get Validators
```bash
curl -X POST http://localhost:8545 \
-H "Content-Type: application/json" \
-H "Origin: http://localhost:8545" \
-d '{
    "jsonrpc":"2.0",
    "method":"getValidators",
    "params":[],
    "id":1
}'
```

#### 5. Check Transaction Status
```bash
curl -X POST http://localhost:8545 \
-H "Content-Type: application/json" \
-H "Origin: http://localhost:8545" \
-d '{
    "jsonrpc":"2.0",
    "method":"getTransaction",
    "params":["TRANSACTION_HASH"],
    "id":1
}'
```

### Example Wallet Integration

Here's a simple example of how to integrate the light client with a JavaScript wallet:

```javascript
class ThrylosWallet {
    constructor(lightClientUrl = 'http://localhost:8545') {
        this.clientUrl = lightClientUrl;
    }

    async makeRequest(method, params) {
        const response = await fetch(this.clientUrl, {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Origin': window.location.origin
            },
            body: JSON.stringify({
                jsonrpc: '2.0',
                method: method,
                params: params,
                id: 1
            })
        });
        return await response.json();
    }

    async getBalance(address) {
        return await this.makeRequest('getBalance', [address]);
    }

    async sendTransaction(from, to, amount) {
        return await this.makeRequest('submitTransaction', [{
            from: from,
            to: to,
            amount: amount
        }]);
    }

    async getTransactionHistory(address) {
        return await this.makeRequest('getTransactions', [address]);
    }
}

// Usage example:
const wallet = new ThrylosWallet();
wallet.getBalance('ADDRESS_HERE')
    .then(balance => console.log('Balance:', balance))
    .catch(error => console.error('Error:', error));
```

## Important Notes

1. **Security**
   - Always use HTTPS in production
   - Implement proper key management
   - Never send private keys to the light client

2. **Rate Limiting**
   - Be mindful of request frequency
   - Implement caching where appropriate
   - Handle connection errors gracefully

3. **Best Practices**
   - Validate all inputs before sending
   - Implement proper error handling
   - Use WebSocket for real-time updates
   - Keep track of transaction history locally

## Error Handling

The light client uses standard JSON-RPC error codes:
- -32700: Parse error
- -32600: Invalid request
- -32601: Method not found
- -32602: Invalid params
- -32603: Internal error

Example error response:
```json
{
    "jsonrpc": "2.0",
    "error": {
        "code": -32603,
        "message": "Internal error",
        "data": null
    },
    "id": 1
}
```

## Support

For issues and feature requests, please open an issue in the GitHub repository.