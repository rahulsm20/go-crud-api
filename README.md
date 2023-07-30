# GO CRUD API 
This is a CRUD (Create, Read, Update, Delete) API built in Go using the Gin web framework. The API allows users to perform basic CRUD operations on posts and supports user signup and authentication. It uses a PostgreSQL database to store post and user data.
### Dependencies
* Framework
    * Gin 
* Authentication
  * Crypto
  * JWT
* Database Utilities
  * GORM 
  * GORM - Postgres Driver
  
### Database
* PostgreSQL

## Requirements

- Go (v1.16 or higher)
- PostgreSQL database
- Postman (optional, for testing the API)

## Dependencies
- Gin
- Authentication
  - jwt-go
- ORM
  - GORM
  - GORM/Postgres Driver

## Installation

1. Clone the repository to your local machine:

```
git clone https://github.com/rahulsm20/go-crud-api.git
cd go-crud-api
```

2. Set up the PostgreSQL database:
   - Create a .env file in the root directory.
   - Define the database connection string as DB_URL.

3. Install the dependencies and run the API server:

```
go mod download
go run main.go
```
Alternatively, you can use the following command to automatically recompile and restart the server whenever changes are made to the source code.

```
CompileDaemon -command="./go-crud-api"
```

The API server will run on `http://localhost:3000`.