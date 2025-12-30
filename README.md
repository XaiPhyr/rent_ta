# RENT TA

Connect with available warehouse and open gym spaces near you and book them instantly for short- or long-term use. Our platform makes it easy to find storage, or covered court in one place.

## Tech Stack

- **Go (Golang)**: Programming language used to build the API.
- **Gin**: Web framework for building fast and efficient web applications.
- **Bun**: PostgreSQL ORM for Go, enabling interaction with a PostgreSQL database.
- **PostgreSQL**: Relational database used for data persistence.

## Prerequisites

Before running the project, ensure you have the following installed:

- [Go (1.25+)](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/download/)
- [Git](https://git-scm.com/)
- [Go Modules](https://blog.golang.org/using-go-modules)
- (Optional, coming soon): [Docker](https://www.docker.com/)

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/XaiPhyr/rent_ta
cd rent_ta
```

### 2. Install Dependencies

```bash
cd api
go mod tidy
```

### 3. Create Database

```sql
CREATE DATABASE your_db_name;
```

### 4. Configure config.yml file

remove .example on conf/config.yml.example

```yml
env: 'dev'

server:
  host: 'localhost:8200'
  endpoint: '/v1/api'
  jwt_key: ''

database:
  dsn: 'postgres://postgres:postgres@localhost:5432/your_db?sslmode=disable&TimeZone=Asia/Singapore'
  bundebug: true
```

### 5. SQL tables

SQL files can be found on sql/

### 6. Run the API

```bash
go run .
```

By default, the API will be accessible at http://localhost:8200.
