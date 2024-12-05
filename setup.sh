#!/bin/bash

# Exit configs
set -e
set -o pipefail
set -u

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
NC='\033[0m'

# Constants
SERVICES_DIR="$HOME/.local/share/hoe-crawler"
REPO_URL="https://github.com/yoori/flare-bypasser.git"
PROJECT_ROOT=$(pwd) # Store project root at the start

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

# Cleanup function for failed bypass setup
cleanup() {
    if [ -d "$SERVICES_DIR/flare-bypasser" ]; then
        log "Cleaning up failed installation..."
        rm -rf "$SERVICES_DIR/flare-bypasser"
    fi
}

# Trap signals for cleanup
trap cleanup ERR INT TERM

# Check required commands
check_requirements() {
    log "Checking requirements..."

    if ! command -v docker &> /dev/null; then
        error "Docker is required but not installed."
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
    mkdir -p "$SERVICES_DIR" || error "Failed to create services directory"

    log "Setting up Cloudflare bypass service in $SERVICES_DIR"
    cd "$SERVICES_DIR" || error "Failed to change directory"

    if [ ! -d "flare-bypasser" ]; then
        log "Cloning Cloudflare bypass repo..."
        if ! git clone "$REPO_URL" 2>&1; then
            error "Failed to clone bypass repo"
        fi
        cd flare-bypasser || error "Failed to enter directory"

        log "Building Docker image..."
        if ! docker build -t flare-bypasser .; then
            cd "$PROJECT_ROOT" || true
            cleanup
            error "Failed to build image"
        fi
    else
        warning "flare-bypasser directory exists, checking image..."
        cd flare-bypasser || error "Failed to enter directory"

        if ! docker image inspect flare-bypasser &> /dev/null; then
            log "Image not found, rebuilding..."
            if ! docker build -t flare-bypasser .; then
                cd "$PROJECT_ROOT" || true
                cleanup
                error "Failed to build image"
            fi
        fi
    fi

    cd "$PROJECT_ROOT" || error "Failed to return to project root"
}

# Start all required services
start_services() {
    log "Starting services..."

    # Start Cloudflare bypass service
    log "Starting Cloudflare bypass service..."
    cd "$HOME/.local/share/hoe-crawler/flare-bypasser" || error "Failed to change directory"

    if ! docker compose up -d; then
        warning "Docker compose failed, cleaning up..."
        cd "$PROJECT_ROOT" || true
        cleanup
        error "Failed to start bypass service"
    fi

    # Start database from project root
    cd "$PROJECT_ROOT" || error "Failed to return to project root"
    log "Starting MariaDB..."
    if ! docker compose up -d; then
        error "Failed to start database service"
    fi

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