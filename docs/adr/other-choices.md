## Trade-offs

- We chose MongoDB over Elasticsearch for our log data storage needs due to its flexibility in handling varying data formats, simpler operational requirements, and better alignment with our use case of general-purpose database functionality rather than specialized search and analytics.
- Render was selected for application deployment because of its user-friendly interface, automatic scaling capabilities, cost-effectiveness for small projects, support for multiple programming languages, and built-in CI/CD features, which streamline our development and deployment processes.
- The decision to use a layered architecture (API, Operation, Service Provider, Contract) was made to separate concerns, enhance maintainability, and improve testability of the codebase. Each layer has distinct responsibilities, allowing for clearer organization and easier updates or changes in the future.
- The choice of Go as the programming language was driven by its performance, concurrency support, and strong standard library, which are well-suited for building scalable and efficient web applications. Go's simplicity and ease of deployment also contribute to faster development cycles.

## Hardest Parts to Implement
- Implementing the CQRS pattern effectively required careful design to ensure that commands and queries were properly separated, which involved creating distinct interfaces and handling data consistency between them.
- Writing a dockerized approach for local development and deployment posed challenges in ensuring that all services communicated correctly and that the environment was consistent across different setups.
- Setting up the deployment pipeline with Render

## Missing Features
- Adding migration scripts for MongoDB to handle schema changes and data transformations as the application evolves.
- Adding a cron job to auto pull data from external apis periodically.
- Implementing advanced logging and monitoring capabilities to track application performance and errors in production.
- Enhancing the API with rate limiting and authentication mechanisms to secure endpoints and manage traffic effectively.
- Integrating a more sophisticated error handling and retry mechanism for external service calls to improve resilience against transient failures.
- Implementing integration tests to ensure that the entire system works as expected when all components are combined, especially for the API and service provider interactions.
- Adding support for more complex query operations in the API layer to allow users to filter and sort data based on various criteria.
