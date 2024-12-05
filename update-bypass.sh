#!/bin/bash

set -e # Exit on error
set -o pipefail # Exit on pipe failures
set -u # Exit on undefined variables

# Colors for terminal output
RED='\033[0;31m'
GREEN='\033[0;32m'
NC='\033[0m'

log() {
    echo -e "${GREEN}[+]${NC} $1"
}

error() {
    echo -e "${RED}[!]${NC} $1"
    exit 1
}

# Check bypass service directory
SERVICES_DIR="$HOME/.local/share/hoe-crawler"
if [ ! -d "$SERVICES_DIR/flare-bypasser" ]; then
    error "flare-bypasser directory not found. Run setup.sh first"
fi

# Update repository code
log "Pulling latest code..."
cd "$SERVICES_DIR/flare-bypasser"
git pull || error "Failed to pull latest code"

# Rebuild service
log "Rebuilding service..."
docker-compose down -v || error "Failed to stop existing containers"
docker-compose up -d || error "Failed to start new containers"

log "Service updated successfully!"