# Title
**Choosing MongoDB as db**

## Status
We proposed 2 database which is provided by GCP.
1. Elastic Seacrch
2. MongoDB
   We have accepted MongoDB for our use case.

## Context
We have to save different logs data in database. The data from each log could have different data format and body.
We should be able to use same database to store different data from each log data pipeline.

## Decision
**Data Structure Flexibility**
- MongoDB's document-based structure handles varying log formats more naturally
- No need to define strict schemas upfront, unlike Elasticsearch's mapping requirements
- Better support for nested and complex data structures
**Operational Simplicity**
- Simpler deployment and maintenance compared to Elasticsearch clusters
- Less memory and resource intensive for general database operations
**Development Experience**
- More intuitive query language (MQL) compared to Elasticsearch's Query DSL
- Better support for traditional database operations (CRUD)
- Stronger consistency guarantees for transactional operations
**Use Case Alignment**
- Elasticsearch is optimized for search and analytics, which may be overkill for log storage
- MongoDB provides better general-purpose database functionality

## Consequences
It will help with storing different data structure in same db.