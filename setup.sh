#!/bin/bash

# Exit on error. Append "|| true" if you expect an error.
set -e

# Exit on error in pipe
set -o pipefail

# Exit on undefined variable
set -u

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m' # No Color

# Log functions
log() {
    echo -e "${GREEN}[+]${NC} $1"
}

error() {
    echo -e "${RED}[!]${NC} $1"
    exit 1
}

warning() {
    echo -e "${YELLOW}[!]${NC} $1"
}

# Check required commands
check_requirements() {
    log "Checking requirements..."

    if ! command -v docker &> /dev/null; then
        error "Docker is required but not installed."
    fi

    if ! command -v docker-compose &> /dev/null; then
        error "Docker Compose is required but not installed."
    fi

    if ! command -v git &> /dev/null; then
        error "Git is required but not installed."
    fi

    if ! docker info &> /dev/null; then
        error "Docker daemon is not running."
    fi
}

# Setup Cloudflare bypass service
setup_bypass() {
    # Create services directory
    SERVICES_DIR="$HOME/.local/share/hoe-crawler"
    mkdir -p "$SERVICES_DIR" || error "Failed to create services directory"

    log "Setting up Cloudflare bypass service in $SERVICES_DIR"

    # Change to services directory
    cd "$SERVICES_DIR" || error "Failed to change directory to $SERVICES_DIR"

    if [ ! -d "flare-bypasser" ]; then
        log "Cloning Cloudflare bypass repo..."
        git clone https://github.com/yoori/flare-bypasser.git || error "Failed to clone bypass repo"
        cd flare-bypasser || error "Failed to enter flare-bypasser directory"

        log "Building Docker image for bypass service..."
        if ! docker build -t flare-bypasser .; then
            error "Failed to build bypass service Docker image"
        fi
    else
        warning "flare-bypasser directory already exists, checking service..."
        cd flare-bypasser || error "Failed to enter existing flare-bypasser directory"

        # Check if image exists
        if ! docker image inspect flare-bypasser &> /dev/null; then
            log "Bypass service image not found, rebuilding..."
            if ! docker build -t flare-bypasser .; then
                error "Failed to build bypass service Docker image"
            fi
        fi
    fi

    # Return to original directory
    cd - > /dev/null || error "Failed to return to project directory"
}

# Start all required services
start_services() {
    log "Starting services..."

    # Start Cloudflare bypass service
    log "Starting Cloudflare bypass service..."
    cd "$HOME/.local/share/hoe-crawler/flare-bypasser" || error "Failed to change to bypass service directory"
    if ! docker-compose up -d; then
        error "Failed to start bypass service"
    fi
    cd - > /dev/null || error "Failed to return to project directory"

    # Start database
    log "Starting MariaDB..."
    if ! docker-compose up -d; then
        error "Failed to start database service"
    fi

    # Wait and verify services
    log "Waiting for services to be ready..."
    sleep 10

    log "All services are up and running"
}

main() {
    log "Starting setup..."

    check_requirements
    setup_bypass
    start_services

    log "Setup completed successfully!"
}

main