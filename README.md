# Webhook in Go

- exercise in goroutines
- exercise in docker compose (containerise the hook, queue, and db)
- Dockerfiles for all


## Plan
1. HTTP payload sample sender
    - mock the app, small python api
    - fastapi
    - make a mock payload to send to the msg queue
2. Queue (ideally distributed, but just a simple container for now)
3. Webhook service sends to multiple other APIs
4. Destinations
    - DB (only need something small, MySQL/Postgres/somethingelse
    - Mock Martech service

## Questions
- What messages will this take? Transactions? Martech Stuff?
- Queue details, ordered? TTL/retry? DEDUPLICATING A MUST
- DB: will need multiple tables for different events



