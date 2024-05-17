# Blog Aggregator

The Blog Aggregator API retrieves blog content from various web sources, including RSS feeds. Data is fetched concurrently by the worker function and parsed for storage in the database. Authentication is managed through API keys provided during user creation. This authentication process is implemented using authenticate middleware, which is added to handlers requiring authentication. PostgreSQL is the chosen database, Goose serves as the migration tool, and Sqlc is utilized to generate type-safe code from SQL.

## Features

- Retrieves blog content from various web sources
- Fetches data concurrently
- Parses data for storage in PostgreSQL database
- Authentication managed through API keys
- Middleware for authentication
- Utilizes PostgreSQL as the database
- Goose for database migrations
- Sqlc for generating type-safe code from SQL

## Technologies Used

- Go (Golang)
- PostgreSQL
- Goose (Database Migration Tool)
- Sqlc (SQL Compiler)
- Middleware

## Installation

1. Clone the repository:
```
git clone github.com/friday1602/blog-aggregator
```
2. Install dependencies:
```
go mod download
```
3. Install Goose for database migrations:
```
go install github.com/pressly/goose/v3/cmd/goose@latest
```
4.Configure environment variables:
- Edit the `.env` file with your configurations.
5.Run database migrations:
```
goose -dir sql/schema postgres <database_connection_string> up
```
6.Build and run the application:
```
go build -o blog-aggregator && ./blog-aggregator
```
