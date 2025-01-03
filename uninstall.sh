#!/bin/bash

echo "Uninstalling Thrylos Node..."

# Stop the service
echo "Stopping Thrylos node service..."
launchctl unload ~/Library/LaunchAgents/com.thrylos.node.plist 2>/dev/null

# Remove the plist file
echo "Removing launch agent configuration..."
rm -f ~/Library/LaunchAgents/com.thrylos.node.plist

# Optional: Remove logs
read -p "Do you want to remove log files? (y/N) " remove_logs
if [[ $remove_logs =~ ^[Yy]$ ]]; then
    echo "Removing log files..."
    rm -f ~/Library/Logs/thrylos-node.log
    rm -f ~/Library/Logs/thrylos-node.error.log
fi

echo "✅ Thrylos node has been uninstalled successfully!"