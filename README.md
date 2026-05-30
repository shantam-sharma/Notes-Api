# Notes API

A RESTful Notes API built with **Go**, **PostgreSQL**, and **JWT Authentication** following a layered backend architecture.

This project was built to learn and implement real-world backend development concepts including authentication, authorization, database design, middleware, pagination, migrations, and clean architecture patterns.

---

## Features

### Authentication

* User Signup
* User Login
* Password Hashing with bcrypt
* JWT Token Generation
* Protected Routes using Middleware

### Notes Management

* Create Note
* Get All Notes
* Get Note By ID
* Update Note
* Delete Note

### Authorization

* Users can only access their own notes
* Ownership validation on all note operations

### Database

* PostgreSQL
* Foreign Key Relationships
* Cascading Deletes
* Database Migrations

### Additional Features

* Pagination
* Structured JSON Responses
* Layered Architecture
* Environment Variable Configuration

---

## Tech Stack

| Technology     | Purpose                |
| -------------- | ---------------------- |
| Go             | Backend Language       |
| PostgreSQL     | Database               |
| Chi Router     | HTTP Routing           |
| JWT            | Authentication         |
| bcrypt         | Password Hashing       |
| golang-migrate | Database Migrations    |
| godotenv       | Environment Management |

---

## Architecture

```text
Client
   │
   ▼
Handlers
   │
   ▼
Services
   │
   ▼
Repositories
   │
   ▼
PostgreSQL
```

### Layer Responsibilities

#### Handlers

Responsible for:

* Parsing requests
* Returning responses
* HTTP status codes

#### Services

Responsible for:

* Business logic
* Validation
* Authentication logic

#### Repositories

Responsible for:

* Database queries
* PostgreSQL interaction

---

## Project Structure

```text
Notes-Api/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── database/
│   ├── handlers/
│   ├── middleware/
│   ├── models/
│   ├── repositories/
│   ├── services/
│   └── utils/
│
├── migrations/
│
├── .env
├── go.mod
└── README.md
```

---

## Database Schema

### Users Table

```sql
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
```

### Notes Table

```sql
CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT fk_user
        FOREIGN KEY(user_id)
        REFERENCES users(id)
        ON DELETE CASCADE
);
```

---

## API Endpoints

### Authentication

| Method | Endpoint | Description   |
| ------ | -------- | ------------- |
| POST   | /signup  | Register User |
| POST   | /login   | Login User    |

---

### Notes

| Method | Endpoint    | Description    |
| ------ | ----------- | -------------- |
| POST   | /notes      | Create Note    |
| GET    | /notes      | Get User Notes |
| GET    | /notes/{id} | Get Note By ID |
| PUT    | /notes/{id} | Update Note    |
| DELETE | /notes/{id} | Delete Note    |

---

## Pagination

Retrieve notes with pagination.

```http
GET /notes?page=1&limit=10
```

### Query Parameters

| Parameter | Description    |
| --------- | -------------- |
| page      | Page Number    |
| limit     | Notes Per Page |

---

## Environment Variables

Create a `.env` file:

```env
DB_URL=<postgres-connection-string>
JWT_SECRET=<jwt-secret>
```

---

## Running Locally

### Clone Repository

```bash
git clone https://github.com/shantam-sharma/Notes-Api.git
cd Notes-Api
```

### Install Dependencies

```bash
go mod tidy
```

### Configure Environment

Create a `.env` file and add:

```env
DB_URL=<postgres-connection-string>
JWT_SECRET=<jwt-secret>
```

### Run Application

```bash
go run cmd/main.go
```

Server starts on:

```text
http://localhost:8080
```

---

## Sample Login Response

```json
{
  "token": "<jwt-token>"
}
```

---

## Sample Error Response

```json
{
  "error": "unauthorized"
}
```

---

## Concepts Learned

This project demonstrates:

* REST API Design
* Layered Architecture
* Repository Pattern
* JWT Authentication
* Authorization
* Middleware
* PostgreSQL Integration
* Database Migrations
* Pagination
* Error Handling
* Environment Configuration

---

## Future Improvements

* Docker
* Docker Compose
* Unit Tests
* Integration Tests
* Refresh Tokens
* Structured Logging
* CI/CD Pipeline

---

## Version

Current Release:

```text
v1.0.0
```
