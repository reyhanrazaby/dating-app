# Dating App

Contains RESTful API for sign-up and login feature using Go 1.21.3

## Project structure
- **server**: Sets up the HTTP server, routes, and middlewares
- **usecase**: Implements api routes (handler) and business logic (service)
- **datasource**: Handles data sources and database management
- **entity**: Defines data structure or entities
- **util**: Contains utility functions
- **errors**: Defines custom `error`
- **scripts**: Contains scripts for development or testing

## How to run
   ```bash
   go run .
   ```

## Run all tests
   ```bash
   go test ./...
   ```

## Test with curl
- With terminal
   ```bash
   bash scripts/[sh file]
   ```
   OR
   ```bash
   sh scripts/[sh file]
   ```
- With Postman
  - Click import
  - Paste the `.sh` file content OR choose `scripts` folder to import all tests

  