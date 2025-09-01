# DynamoDB Go - Test Project

A Go application for testing and demonstrating DynamoDB operations using AWS SDK for Go.

## Features

- **Local DynamoDB Testing**: Uses Docker Compose to run DynamoDB Local for development and testing
- **DynamoDB Operations**: Demonstrates common DynamoDB operations including:
  - Table creation and management
  - Item insertion with conditional checks
  - Item retrieval and querying
  - Batch operations and item deletion
- **User Management**: Includes a user store interface for managing user records
- **Testing Suite**: Comprehensive tests for database operations
- **Error Handling**: Examples of handling DynamoDB-specific errors like conditional check failures

## Project Structure

```
├── main.go              # Main application entry point with examples
├── database/
│   ├── database.go      # DynamoDB client and user operations
│   └── db.go           # Database utilities and table management
├── types/
│   └── types.go        # Type definitions
├── api/
│   └── api.go          # API layer
├── tests/
│   └── dynamo_test.go  # Unit tests for DynamoDB operations
├── docker-compose.yml   # DynamoDB Local container configuration
├── go.mod              # Go module dependencies
└── go.sum              # Go module checksums
```

## Dependencies

- **AWS SDK for Go v1** (`github.com/aws/aws-sdk-go`)
- **AWS SDK for Go v2** (`github.com/aws/aws-sdk-go-v2`)
- **AWS Lambda Go** (`github.com/aws/aws-lambda-go`)

## Getting Started

1. **Start DynamoDB Local**:
   ```bash
   docker-compose up -d
   ```

2. **Run the application**:
   ```bash
   go run main.go
   ```

3. **Run tests**:
   ```bash
   go test ./tests/
   ```

## Examples Included

- **Conditional Check Failures**: Demonstrates how to handle duplicate item insertion attempts
- **Bulk Data Operations**: Seeds the database with 500 test items
- **Data Cleanup**: Shows how to delete all items from a table
- **User Management**: CRUD operations for user records

## Configuration

- DynamoDB Local runs on port 8000 (configurable in docker-compose.yml)
- Default table name: `go-dynamodb-reference-table`
- Test table name: `test-db`

This project serves as a reference implementation for Go developers working with DynamoDB, providing practical examples of common patterns and operations.
