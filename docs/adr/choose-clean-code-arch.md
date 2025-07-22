# Architecture

We followed the design principles of clean architecture, DDD along with goodness of CQRS.

* **Clean Architecture**

  * It talks about how to define vertical boundaries within our code base and apply very neat separation of concerns.
* **DDD**

  * It advocates about how to model the domain entities encapsulating intra-aggregate invariants and behaviours
* **CQRS**

  * Help to see create/update operations differently from query operations
  * It is not always about event sourcing and having a separate database

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
