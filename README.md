# First Go HTTP Server

This project is a foundational Go-based HTTP server designed to manage user accounts and RSS feed subscriptions. It provides RESTful APIs for user registration, authentication, feed management, and subscription handling.
This project could help you understand basic http-server cereation along with middleware integration as well as the go routines used with waitGroup. 

## Features

- **User Management**: Register and authenticate users securely.
- **Feed Handling**: Add, list, and manage RSS feeds.
- **Subscription Management**: Subscribe to feeds and manage user-specific subscriptions.
- **Authentication Middleware**: Protect endpoints using middleware for token-based authentication.
- **Health Checks**: Readiness endpoint to monitor server health.
- **Database Integration**: Utilizes PostgreSQL with SQLC for type-safe database interactions.
- **RSS Scraper**: Fetch and parse RSS feeds to keep content up-to-date.

## Project Structure

- `internal/` – Internal packages
- `sql/` – SQL migrations and SQLC queries
- `vendor/` – Go vendor dependencies
- `handler_feeds.go` – Feed creation and listing handlers
- `handler_feeds_follow.go` – Feed follow/unfollow handlers
- `handler_readiness.go` – Readiness check handler
- `handler_user.go` – User registration and login handlers
- `json.go` – JSON utility functions
- `main.go` – Application entry point
- `middleware_auth.go` – Auth middleware
- `models.go` – Data models
- `rss.go` – RSS feed parser
- `scrapper.go` – RSS scraping logic
- `sqlc.yml` – SQLC config file
- `go.mod` – Go module metadata
- `go.sum` – Dependency checksums
- `RSS AGG.postman_collection.json` – Postman collection for API testing


## Getting Started

### Prerequisites

- Go 1.20 or higher
- PostgreSQL database
- goose
- sqlc

### Installation

1. **Clone the repository**:
   ```bash
   git clone https://github.com/Rushikesh-purohit-0503/first-go-http-server.git
   cd first-go-http-server
2. **go mod vednor **:
   ```bash
   go mod vendor
3. **go mod tidy **:
   ```bash
   go mod tidy 


# Assuming you have a migration tool installed
  ```bash
  goose postgres postgres://user:password@localhost:5432/dbname up

  sqlc generate

  go build -o app
  ./app
