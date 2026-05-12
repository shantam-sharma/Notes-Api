# Notes API — Go + PostgreSQL Backend Project

## Overview

This project is a backend Notes API built using Go and PostgreSQL.
The goal of the project is not just to build CRUD endpoints, but to deeply understand backend engineering concepts such as:

* API architecture
* authentication and authorization
* database design
* relational data modeling
* password hashing
* middleware
* clean project structure
* Dockerized development
* backend best practices

The project is being built incrementally to focus on learning backend development properly instead of copying large tutorials.

---

# Tech Stack

## Backend

* Go (Golang)

## Database

* PostgreSQL

## Database Driver

* pgx
* database/sql

## Authentication

* bcrypt password hashing
* JWT authentication (in progress)

## Environment Management

* .env configuration

## Planned Infrastructure

* Docker
* Docker Compose

---

# Features Implemented

## Database Setup

* Connected Go application to PostgreSQL
* Implemented DB connection handling
* Added DB ping verification
* Added error handling for DB failures

---

## Database Schema Design

### Users Table

Fields:

* id
* name
* email
* password_hash
* created_at

### Notes Table

Fields:

* id
* user_id
* title
* content
* created_at
* updated_at

---

## Relational Database Design

Implemented:

* one-to-many relationship between users and notes
* ownership structure using user_id
* unique email handling

---

## Authentication — Signup

Implemented:

* POST /signup endpoint
* request body parsing
* JSON decoding
* input validation
* password hashing using bcrypt
* storing hashed passwords securely
* user insertion into PostgreSQL
* proper HTTP status handling

---

# Current Project Structure

```text
notes-api/
│
├── cmd/
│   └── main.go
│
├── internal/
│   ├── database/
│   ├── handlers/
│   ├── repositories/
│   ├── middleware/
│   ├── services/
│   └── models/
│
├── migrations/
│
├── .env
├── go.mod
├── go.sum
└── README.md
```

---

# Concepts Learned So Far

## Backend Architecture

* separation of concerns
* layered backend structure
* handler vs repository responsibilities

---

## Database Concepts

* relational databases
* foreign keys
* schema design
* unique constraints
* ownership relationships

---

## Authentication Concepts

* password hashing
* bcrypt
* why raw passwords should never be stored
* request validation

---

## Go Backend Development

* package organization
* database/sql usage
* pgx driver integration
* environment variables
* structured error handling

---

# API Endpoints

## Implemented

### Signup

```http
POST /signup
```

Request Example:

```json
{
  "name": "John",
  "email": "john@example.com",
  "password": "secret123"
}
```

---

# Features Currently In Progress

## Login System

Planned:

* POST /login
* password verification
* JWT generation
* secure authentication flow

---

## JWT Authentication

Planned:

* JWT token generation
* token validation
* expiration handling
* authorization middleware

---

## Protected Routes

Planned:

* middleware-based route protection
* authenticated note access
* user-specific request handling

---

## Notes CRUD

Planned:

* create notes
* get notes
* update notes
* delete notes

---

## Ownership Authorization

Planned:

* ensure users can only access their own notes
* secure note querying using user_id

---

## Pagination

Planned:

* page and limit query parameters
* SQL LIMIT/OFFSET usage
* paginated API responses

---

## Docker Support

Planned:

* Dockerfile
* Docker Compose
* containerized PostgreSQL
* reproducible development environment

---

# Future Improvements

## Possible Additions

* refresh tokens
* role-based authorization
* request logging
* structured logging
* configuration management
* API versioning
* unit testing
* integration testing
* SQL migrations
* rate limiting
* CI/CD pipeline

---

# Learning Goals Behind This Project

This project is being used to learn and practice:

* backend engineering fundamentals
* production-style API development
* authentication systems
* clean architecture
* relational database design
* Go backend development
* debugging and error handling
* scalable backend structure

---

# Project Status

Current Status:

* Database connected
* Schema designed
* Signup authentication completed
* Login + JWT authentication in progress

---

# Notes

This project is intentionally being built step-by-step to focus on understanding backend systems deeply instead of relying heavily on tutorials or generated boilerplate.
