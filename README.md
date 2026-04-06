# Notes API

A RESTful Notes API built with Go, Gin, and PostgreSQL. Supports user authentication with JWT and full-text note search.

## Tech Stack
- **Language**: Go
- **Framework**: Gin
- **Database**: PostgreSQL
- **Auth**: JWT (golang-jwt)
- **Password Hashing**: bcrypt

## Features
- User registration and login
- JWT-based authentication
- Create, read, update, delete notes
- Full-text search across notes

## Project Structure
notes-api/
├── config/       # Database connection
├── handlers/     # HTTP request handlers
├── middleware/   # JWT auth middleware
├── models/       # Data structures
├── repository/   # Database queries
└── routes/       # Route definitions

## Getting Started

### Prerequisites
- Go 1.21+
- PostgreSQL

### Setup
1. Clone the repo
```bash
   git clone https://github.com/yourusername/notes-api.git
   cd notes-api
```

2. Create a `.env` file
DB_URL=postgres://postgres:password@localhost:5432/notes_db?sslmode=disable
JWT_SECRET=your_secret_key

3. Run database migrations
```sql
   CREATE TABLE users (
       id SERIAL PRIMARY KEY,
       name VARCHAR(100) NOT NULL,
       email VARCHAR(100) UNIQUE NOT NULL,
       password VARCHAR(255) NOT NULL,
       created_at TIMESTAMP DEFAULT NOW()
   );

   CREATE TABLE notes (
       id SERIAL PRIMARY KEY,
       user_id INT NOT NULL REFERENCES users(id) ON DELETE CASCADE,
       title VARCHAR(255) NOT NULL,
       content TEXT,
       created_at TIMESTAMP DEFAULT NOW(),
       updated_at TIMESTAMP DEFAULT NOW()
   );
```

4. Run the server
```bash
   go run main.go
```
   Server runs at `http://localhost:8080`

## API Endpoints

| Method | Endpoint | Auth | Description |
|--------|----------|------|-------------|
| POST | /auth/register | No | Register new user |
| POST | /auth/login | No | Login and get JWT token |
| POST | /notes/ | Yes | Create a note |
| GET | /notes/ | Yes | Get all notes |
| GET | /notes/:id | Yes | Get note by ID |
| PUT | /notes/:id | Yes | Update a note |
| DELETE | /notes/:id | Yes | Delete a note |
| GET | /notes/search?q= | Yes | Search notes |