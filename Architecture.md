Senior approach:

1. define requirements
2. design entities
3. design API
4. design DB schema
5. define architecture
6. THEN code

______________________________________________________________________

Phase 1. System
- user can
    sign up , login , own notes
- Note
    belongs to user , has content , can be updated and deleted

Phase 2. Relationship
- one user = many notes
    notes table need [UserID]
    ownership must be forced
    queries must filter by authenticated users

Phase 3. Api design
- Auth
    POST /signup
    POST /login
- Notes
    POST   /notes
    GET    /notes
    GET    /notes/:id
    PUT    /notes/:id
    DELETE /notes/:id

Phase 4. Database design
- User table
    id , name , email , created at .
    you should never store raw password
- Notes table
    id, user_id , title , content , created at , created by.
Phase 5. Project Structure
```
    notes-api/
    │
    ├── cmd/
    │   └── main.go
    │
    ├── internal/
    │   ├── handlers/
    │   ├── services/
    │   ├── repositories/
    │   ├── middleware/
    │   ├── models/
    │   └── database/
    │
    ├── migrations/
    │
    ├── docker-compose.yml
    ├── Dockerfile
    ├── go.mod
    └── .env
```

Phase 6. Each layer meaning and tasks
- handlers
    http only , parse request , validate method , send response
- services
    Business Logic (like a brain)
    signup user , hash password , give id , create note etc.
- repositories
    database access only (this layer talks to SQL)
    insert user , fetch note , delete note.

Phase 7. AUTHENTICATION
- Sign up flow
```
    User sends email/password
            ↓
    Server validates input
            ↓
    Password gets hashed
            ↓
    User stored in DB
            ↓
    Success response
```
- Log in flow
```
    User sends credentials
            ↓
    Find user by email
            ↓
    Compare password hash
            ↓
    Generate JWT token
            ↓
    Return token
```
- Protected route flow
```
    Client sends JWT
            ↓
    Middleware validates token
            ↓
    Extract user ID
            ↓
    Attach user to request context
            ↓
    Handler uses authenticated user
```
