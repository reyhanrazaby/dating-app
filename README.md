# Dating App

Contains RESTful API for sign-up and login feature using Go 1.21.3

## Project structure
- **server**: Sets up the HTTP server, routes, and middlewares
- **usecase**: Implements api routes (handler) and business logic (service)
- **datasource**: Handles data sources and database management
- **entity**: Defines data structure or entities
- **util**: Contains utility functions
- **errors**: Defines custom `error`
- **scripts**: Contains scripts for deployment, testing, etc

## How to run
   ```bash
   go run .
   ```

## Run all tests
   ```bash
   go test ./...
   ```

## Test with curl
#### With terminal
   ```bash
   bash scripts/sign_up_curl.sh
   bash scripts/login_curl.sh
   ```
#### With Postman
  - Click import
  - Paste the content of `_curl.sh` file you want to test

## Deployment (to docker)
Make sure docker is running before executing this script
```bash
bash scripts/deploy.sh
```

## Linter
Make sure docker is running before executing this script
```bash
bash scripts/linter.sh
```
  