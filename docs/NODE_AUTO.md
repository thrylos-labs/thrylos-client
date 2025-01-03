Running a Thrylos Node - Simple Setup Guide
This guide helps you set up a Thrylos node that automatically starts when you turn on your computer.

Prerequisites

Go installed on your computer (go version to verify)

Quick Installation

Clone the repository:

git clone <repository-url>
cd thrylos-client

Run the installation script:

./install.sh
That's it! Your node will start automatically and restart when you reboot your computer.
Managing Your Node
Common commands:
bashCopy# View logs
tail -f ~/Library/Logs/thrylos-node.log

# Stop node
./uninstall.sh

# Start node again
./install.sh
Verify It's Working

The installation script will automatically test the connection. You can manually test it with:

curl -X POST http://localhost:8545 \
-H "Content-Type: application/json" \
-d '{
    "jsonrpc":"2.0",
    "method":"getBlockchainInfo",
    "params":[],
    "id":1
}'

You should see blockchain information in the response.

# Uninstalling
To remove the node completely:
./uninstall.sh
Troubleshooting
If you encounter issues:

Verify Go is installed correctly: go version
Check the logs: tail -f ~/Library/Logs/thrylos-node.log
Ensure port 8545 isn't being used by another application
Make sure you have necessary permissions in your working directory

For additional support:

Join our Discord
Check the documentation
View detailed logs at ~/Library/Logs/thrylos-node.error.log