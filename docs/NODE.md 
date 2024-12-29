# Running a Thrylos Node

This guide explains how to run a Thrylos node, either as a validator or a regular network participant.

## Types of Node Participation

### 1. Regular Node (Non-Validator)
Running a regular node helps support the network by:
- Verifying transactions
- Relaying blockchain data
- Increasing network decentralization
- Providing network redundancy

To run a regular node:
```bash
# Start the light client
go run main.go -address :8545 -seed localhost:8546 -network mainnet
```

Configuration options:
- `-address`: Your node's listening address
- `-seed`: Initial seed node to connect to
- `-network`: Choose 'mainnet' or 'testnet'

### 2. Validator Node
Validators play a crucial role by:
- Producing new blocks
- Securing the network
- Earning staking rewards
- Participating in consensus

To become a validator:
```bash
# First, check current validators
curl -X POST http://localhost:8545 \
-H "Content-Type: application/json" \
-H "Origin: http://localhost:8545" \
-d '{
    "jsonrpc":"2.0",
    "method":"getValidators",
    "params":[],
    "id":1
}'

# Then stake tokens (minimum 40 THRYLOS required)
curl -X POST http://localhost:8545 \
-H "Content-Type: application/json" \
-H "Origin: http://localhost:8545" \
-d '{
    "jsonrpc":"2.0",
    "method":"stake",
    "params":[{
        "amount": "40",
        "validatorAddress": "YOUR_ADDRESS"
    }],
    "id":1
}'
```

## Network Participation Benefits

### Regular Node Benefits:
- Access to real-time blockchain data
- Support network decentralization
- Run your own infrastructure
- Direct network connection

### Validator Benefits:
- Earn staking rewards (currently 4.8M THRYLOS yearly)
- Participate in network governance
- Help secure the network
- Earn transaction fees

## System Requirements

### Minimum Requirements (Regular Node):
- 2 CPU cores
- 4GB RAM
- 50GB SSD
- Stable internet connection

### Recommended Requirements (Validator):
- 4+ CPU cores
- 8GB+ RAM
- 100GB+ SSD
- High-speed internet connection
- 99.9% uptime

## Monitoring Your Node

Check node status:
```bash
curl -X POST http://localhost:8545 \
-H "Content-Type: application/json" \
-H "Origin: http://localhost:8545" \
-d '{
    "jsonrpc":"2.0",
    "method":"getNetworkHealth",
    "params":[],
    "id":1
}'
```

Monitor validator performance:
```bash
curl -X POST http://localhost:8545 \
-H "Content-Type: application/json" \
-H "Origin: http://localhost:8545" \
-d '{
    "jsonrpc":"2.0",
    "method":"getStakingStats",
    "params":[],
    "id":1
}'
```

## Best Practices

1. Security:
   - Use a firewall
   - Keep software updated
   - Use secure key management
   - Enable SSL/TLS

2. Maintenance:
   - Regular backups
   - Monitor system resources
   - Check logs regularly
   - Stay informed about updates

3. Network:
   - Use static IP
   - Configure proper ports
   - Ensure stable connection
   - Monitor bandwidth usage

## Getting Started Tutorial

1. Set up your node:
```bash
# Create directory
mkdir thrylos-node
cd thrylos-node

# Start the node
go run main.go -address :8545 -seed mainnet.thrylos.org:8546
```

2. Check sync status:
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
```

3. For validators, stake tokens:
```bash
# First ensure you have sufficient balance
# Then stake tokens using the stake method
```

4. Monitor your participation:
```bash
# Check node status regularly
# Monitor system resources
# Keep track of validator performance if applicable
```

## Troubleshooting

Common issues and solutions:
1. Connection issues:
   - Check network connectivity
   - Verify seed node address
   - Check firewall settings

2. Sync issues:
   - Ensure sufficient disk space
   - Check system resources
   - Verify blockchain data integrity

3. Validator issues:
   - Confirm stake amount
   - Check validator status
   - Monitor uptime

## Support

For technical support:
- Join our Discord community
- Check GitHub issues
- Read documentation
- Contact support team

Remember: Running a node contributes to network decentralization and security, whether you're a validator or not!