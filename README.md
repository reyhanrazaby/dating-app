# Dating App

Contains RESTful API for sign-up and login feature using Go 1.21.3

## Project structure
- **server**: Sets up the HTTP server, routes, and middlewares
- **usecase**: Implements api routes (handler) and business logic (service)
- **datasource**: Handles data sources and database management
- **entity**: Defines data structure or entities
- **util**: Contains utility functions
- **errors**: Defines custom `error`

## How to run
   ```bash
   go run .
   ```

## Run all tests
   ```bash
   go test ./...
   ```