Running a Thrylos Node - Simple Setup Guide
This guide helps you set up a Thrylos node that automatically starts when you turn on your computer.
Prerequisites

Go installed on your computer
Thrylos client code downloaded to your preferred directory

Setup Instructions
MacOS Setup

First, identify your key paths:

Note where Go is installed (typically /usr/local/go/bin/go)
Note where you downloaded the Thrylos client code
Create a logs directory if desired


Create the launch agent configuration:

Create and edit the plist file - replace PATHS below with your actual paths
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
        <string>PATH_TO_THRYLOS/main.go</string>
        <string>-address</string>
        <string>:8545</string>
        <string>-seed</string>
        <string>localhost:8546</string>
    </array>
    <key>WorkingDirectory</key>
    <string>PATH_TO_THRYLOS</string>
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
        <string>~/go</string>
    </dict>
</dict>
</plist>
EOF

Before proceeding, edit the plist file and replace:

PATH_TO_THRYLOS with your actual path to the Thrylos client directory
Adjust the GOPATH if yours is different
Modify log file locations if desired


Set up the service:

bashCopy# Set permissions
chmod 644 ~/Library/LaunchAgents/com.thrylos.node.plist

# Load the service
launchctl unload ~/Library/LaunchAgents/com.thrylos.node.plist 2>/dev/null || true
launchctl load ~/Library/LaunchAgents/com.thrylos.node.plist
Managing Your Node
Basic commands to manage your node:
bashCopy# Start the node
launchctl load ~/Library/LaunchAgents/com.thrylos.node.plist

# Check if running
launchctl list | grep thrylos

# Stop the node
launchctl unload ~/Library/LaunchAgents/com.thrylos.node.plist

# View logs
tail -f ~/Library/Logs/thrylos-node.log
tail -f ~/Library/Logs/thrylos-node.error.log
Uninstall
To remove the node service:
bashCopy# 1. Stop the service
launchctl unload ~/Library/LaunchAgents/com.thrylos.node.plist

# 2. Remove the plist file
rm ~/Library/LaunchAgents/com.thrylos.node.plist

# 3. Remove logs (optional)
rm ~/Library/Logs/thrylos-node.log
rm ~/Library/Logs/thrylos-node.error.log
Verify It's Working
Test your node:
bashCopycurl -X POST http://localhost:8545 \
-H "Content-Type: application/json" \
-d '{
    "jsonrpc":"2.0",
    "method":"getBlockchainInfo",
    "params":[],
    "id":1
}'
You should see blockchain information in the response.
Troubleshooting
If you encounter issues:

Verify Go is installed correctly: go version
Check the log files for specific error messages
Ensure port 8545 isn't being used by another application
Verify all paths in the plist file are correct for your system
Make sure you have necessary permissions in your working directory

For additional support, join our Discord or consult the documentation.