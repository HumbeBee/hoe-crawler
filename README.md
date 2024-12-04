# Hoe Crawler
A personal research tool for analyzing escort services data in Vietnam.

### Vision & Planned Features
- **Data Collection**
    - Track available services and providers
    - Monitor price changes
    - Save detailed service information
    - Get notifications for new listings

- **Advanced Analytics**
    - Hotel-based filtering
    - Price history tracking
    - Work history across different locations
    - Visual analytics
        - Price trend charts
        - Location heat maps
        - Availability patterns
        - Custom data reports

**Note:** Currently in early development. This tool is for personal research purposes only. Please follow your local laws and regulations.

## System Requirements
- Go 1.23+
- Docker & Docker Compose
- MariaDB 11.6

## Setup
### 1. Environment Configuration
Create `.env` file from example:
```bash
cp .env.example .env
```

Update environment variables in `.env`:
```env
MYSQL_ROOT_PASSWORD=your_root_password
DB_HOST=localhost 
DB_PORT=3306
DB_USER=your_username
DB_PASSWORD=your_password  
DB_NAME=your_database

LOG_LEVEL=INFO
```

### 2. Start Database
```bash
docker-compose up -d
```

### 3. Install Dependencies
```bash
go mod download
```

## Usage
### Crawl List Page
```bash
go run cmd/list/main.go --site=gaito
```

### Crawl Detail Page
```bash
go run cmd/detail/main.go --site=gaito
```

### Parameters
- `--site`: Target site name (default: "gaito")
- `--log-level`: Logging level (default: "INFO")

## Project Structure
```
├── cmd/                 # Application entry points  
├── internal/            # Internal packages
│   ├── config/          # App configurations
│   ├── definitions/     # Interfaces and type definitions
│   ├── dto/             # Data Transfer Objects
│   ├── infrastructure/  # Infrastructure components (DB, Browser...)
│   ├── interfaces/      # Interface definitions
│   ├── models/          # Domain models
│   ├── repository/      # Database repositories 
│   ├── scrapers/        # Site-specific crawler logic
│   ├── service/         # Business logic
│   └── utils/           # Helper functions
├── docker-compose.yml   # Docker compose config
└── .env.example         # Example env file
```