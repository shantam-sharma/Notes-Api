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

- Notes table

Phase 5. Project Structure
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

Phase 6. Each layer meaning and tasks
