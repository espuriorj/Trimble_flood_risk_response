# Flood Risk Management Service

## Overview
The Flood Risk Management Service is a Go-based RESTful API designed to manage and assess flood risks and reports. It allows users to retrieve and post flood risk data (based on geographic coordinates, risk levels, and discharge rates) and flood reports (detailing affected areas and actions taken). The service uses a PostgreSQL database for persistent storage and is structured to support scalability and maintainability.

## Features
- **Flood Risk Management**: Store and retrieve flood risk data including latitude, longitude, risk level, discharge rate, and assessment timestamp.
- **Flood Reports**: Record and manage flood reports with details about affected areas and response actions.
- **RESTful API**: Provides endpoints for accessing and submitting flood-related data.
- **Middleware**: Includes logging, error recovery, and CORS support for robust API operations.
- **Database Migrations**: Uses `golang-migrate/migrate` for managing database schema changes.

## Project Structure
```
├── Makefile                  # Automates database migrations
├── README.md                 # Project documentation
├── docs                      # Database-related code and migrations
│   └── postman_collection.json  # API collection
├── db                        # Database-related code and migrations
│   ├── 000001_create_flood_tables.down.sql  # Down migration script
│   ├── 000001_create_flood_tables.up.sql   # Up migration script
│   └── postgres.go           # Database connection setup
├── go.mod                    # Go module dependencies
├── go.sum                    # Dependency checksums
├── internal                  # Internal packages
│   ├── errors                # Custom error handling
│   │   └── errors.go
│   ├── handlers              # HTTP request handlers
│   │   ├── handler.go
│   │   └── types.go
│   ├── middleware            # HTTP middleware (logging, CORS, recovery)
│   │   └── middleware.go
│   ├── repository            # Database operations
│   │   ├── repository.go
│   │   └── types.go
│   ├── response              # HTTP response utilities
│   │   └── response.go
│   ├── routes.go             # API route configuration
│   ├── server.go             # HTTP server setup
│   ├── service               # Business logic
│   │   ├── service.go
│   │   └── types.go
│   └── validation            # Input validation logic
│       └── validation.go
└── main.go                   # Application entry point
```

## Prerequisites
- **Go**: Version 1.20 or higher
- **PostgreSQL**: Version 13 or higher
- **golang-migrate/migrate CLI**: For running database migrations
- **Make**: For executing Makefile commands
- **Environment Variables**: Defined in a `.env` file (see [Environment Variables](#environment-variables))

## Installation
1. **Clone the Repository**:
   ```bash
   git clone <repository-url>
   cd floodRiskManagementService
   ```

2. **Install Dependencies**:
   ```bash
   go mod tidy
   ```

3. **Set Up Environment Variables**:
   Create a `.env` file in the project root with the following variables:
   ```
   DB_USER=your_postgres_user
   DB_PASSWORD=your_postgres_password
   DB_HOST=localhost
   DB_PORT=5432
   DB_NAME=flood_risk_management
   SSL_MODE=disable
   ```

4. **Run Database Migrations**:
   Ensure the PostgreSQL database is running, then apply migrations:
   ```bash
   make migrate
   ```

5. **Run the Application**:
   ```bash
   go run main.go
   ```
   The server will start on `http://localhost:8080`.

## API Endpoints
| Method | Endpoint                     | Description                          |
|--------|------------------------------|--------------------------------------|
| GET    | `/flood`                    | Retrieve general flood information   |
| GET    | `/flood/risk`               | Retrieve all flood risks            |
| POST   | `/flood/risk`               | Submit a new flood risk             |
| GET    | `/flood/risk/{id}`          | Retrieve flood risk by ID           |
| POST   | `/flood/report`             | Submit a new flood report           |

## Environment Variables
The application requires the following environment variables in a `.env` file:
- `DB_USER`: PostgreSQL username
- `DB_PASSWORD`: PostgreSQL password
- `DB_HOST`: Database host (e.g., `localhost`)
- `DB_PORT`: Database port (e.g., `5432`)
- `DB_NAME`: Database name
- `SSL_MODE`: SSL mode for PostgreSQL (e.g., `disable`, `require`)

## Development
### Running Locally
1. Ensure PostgreSQL is running and the `.env` file is configured.
2. Apply database migrations using `make migrate`.
3. Start the server with `go run main.go`.
4. Refer [postman collection](docs/postman_collection.json) to test APIs locally 

### Testing
To add tests, place them in the respective package directories (e.g., `internal/handlers`) with filenames ending in `_test.go`. Run tests using:
```bash
go test ./...
```

## Contributing
See [CONTRIBUTION.md](CONTRIBUTION) for guidelines on contributing to this project.

