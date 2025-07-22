# Data Pipeline API

## Introduction

This is a Go-based application following a clean architecture pattern with clear separation of concerns. The application implements a modular design with distinct layers for API handling, business operations, and service provision.

## Architecture

The project follows a layered architecture with the following key components:

### Core Structure

- **`server/`** - HTTP server configuration and startup logic
- **`api/`** - REST API layer with handlers and routing
- **`operation/`** - Business logic layer containing commands and queries
- **`serviceProvider/`** - External service integrations and data providers
- **`contract/`** - Shared contracts and interfaces

### Key Patterns

- **CQRS (Command Query Responsibility Segregation)** - Separate command and query operations
- **Factory Pattern** - Object creation abstraction in `factory/` directories
- **Dependency Injection** - Configuration-based service resolution
- **Clean Architecture** - Domain-driven design with clear boundaries

### Layer Responsibilities

- **API Layer**: HTTP request handling, routing, and response formatting
- **Operation Layer**: Business logic execution, validation, and orchestration
- **Service Provider Layer**: External integrations and data access
- **Contract Layer**: Interface definitions and data transfer objects

## Setup Steps

### Prerequisites

- Go 1.24.4 or higher
- Git

### Install dependencies
`- make deps`

### Format and validate code
`- make fmt`

### Run tests
`- make test`

### Start local development environment
` - make start-server-locally`

## Deployment
**Local Deployment**
- **Start with Docker Compose**: `make start-server-locally`
- **Compile and run locally**: `make compile ./out/api`
- **Stop services**: `make end-server-locally`

## Services Overview
When running locally with Docker Compose, the following services are available:

- **Data Pipeline API**: http://localhost:13000
- **MongoDB**: localhost:27017 (admin/password)
- **Mongo Express**: http://localhost:8081 (admin/pass)
- **ReDoc API Documentation**: http://localhost:9194

## API Endpoints

The Data Pipeline API exposes the following REST endpoints:

#### Create Data
```http
POST /data
{
	"message": "Successfully fetched and stored posts in db"
}
```

#### Get All Data
```http
GET /data
{
  "data": [
    {
      "id": "string",
      "created_at": "string",
      "source": "string",
      "posts": [
        {
          "userId": 1,
          "title": "string",
          "body": "string"
        }
      ]
    }
  ]
}
```

#### Get Data By Id
```http
GET /data?id="string"
{
  "data": [
    {
      "id": "string",
      "created_at": "string",
      "source": "string",
      "posts": [
        {
          "userId": 1,
          "title": "string",
          "body": "string"
        }
      ]
    }
  ]
}
```

#### Get Data By Created At
```http
GET /data?created_at="string"
{
  "data": [
    {
      "id": "string",
      "created_at": "string",
      "source": "string",
      "posts": [
        {
          "userId": 1,
          "title": "string",
          "body": "string"
        }
      ]
    }
  ]
}
```