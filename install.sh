#!/bin/bash

echo "Thrylos Node Setup Script"
echo "========================"

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Error: Go is not installed. Please install Go first."
    exit 1
fi

# Auto-detect Thrylos directory
THRYLOS_PATH="$(pwd)"

# Verify we're in the right directory
if [ ! -f "$THRYLOS_PATH/main.go" ]; then
    echo "Error: Please run this script from the Thrylos client directory containing main.go"
    exit 1
fi

echo "Installing Thrylos node from: $THRYLOS_PATH"

# Create logs directory
mkdir -p ~/Library/Logs

# Create launch agent plist
echo "Creating launch agent configuration..."
cat << EOF > ~/Library/LaunchAgents/com.thrylos.node.plist
<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
    <key>Label</key>
    <string>com.thrylos.node</string>
    <key>ProgramArguments</key>
    <array>
        <string>/usr/local/go/bin/go</string>
        <string>run</string>
        <string>$THRYLOS_PATH/main.go</string>
        <string>-address</string>
        <string>:8545</string>
        <string>-seed</string>
        <string>localhost:8546</string>
    </array>
    <key>WorkingDirectory</key>
    <string>$THRYLOS_PATH</string>
    <key>RunAtLoad</key>
    <true/>
    <key>KeepAlive</key>
    <true/>
    <key>StandardOutPath</key>
    <string>~/Library/Logs/thrylos-node.log</string>
    <key>StandardErrorPath</key>
    <string>~/Library/Logs/thrylos-node.error.log</string>
    <key>EnvironmentVariables</key>
    <dict>
        <key>PATH</key>
        <string>/usr/local/go/bin:/usr/bin:/bin:/usr/sbin:/sbin</string>
        <key>GOPATH</key>
        <string>$HOME/go</string>
    </dict>
</dict>
</plist>
EOF

# Set correct permissions
echo "Setting permissions..."
chmod 644 ~/Library/LaunchAgents/com.thrylos.node.plist

# Load the service
echo "Loading Thrylos node service..."
launchctl unload ~/Library/LaunchAgents/com.thrylos.node.plist 2>/dev/null || true
launchctl load ~/Library/LaunchAgents/com.thrylos.node.plist

# Check if service is running
echo "Checking service status..."
if launchctl list | grep -q "com.thrylos.node"; then
    echo "✅ Thrylos node service has been installed and started successfully!"
    echo "You can view the logs at:"
    echo "  ~/Library/Logs/thrylos-node.log"
    echo "  ~/Library/Logs/thrylos-node.error.log"
else
    echo "⚠️  Warning: Service may not have started properly. Please check the logs."
fi

# Test the node
echo -e "\nTesting node connection..."
sleep 2  # Give the node a moment to start

curl -s -X POST http://localhost:8545 \
-H "Content-Type: application/json" \
-d '{
    "jsonrpc":"2.0",
    "method":"getBlockchainInfo",
    "params":[],
    "id":1
}' > /dev/null

if [ $? -eq 0 ]; then
    echo "✅ Node is responding to requests"
else
    echo "⚠️  Node is not responding yet. Please check the logs for any issues."
fi

echo -e "\nUseful commands:"
echo "- View logs: tail -f ~/Library/Logs/thrylos-node.log"
echo "- Stop node: launchctl unload ~/Library/LaunchAgents/com.thrylos.node.plist"
echo "- Start node: launchctl load ~/Library/LaunchAgents/com.thrylos.node.plist"