# TODO App

A simple TODO application built with Go that allows users to create and manage their todo lists and items. The application follows clean architecture principles and uses JWT for authentication.

## Features

- User authentication with JWT
- Two types of users: regular users and admin users
- CRUD operations for todo lists and items
- Soft delete functionality
- Admin users can view all todo lists
- Regular users can only view and manage their own todo lists

## API Endpoints

### Authentication
- `POST /login` - Login and get JWT token

### Todo Lists
- `GET /api/todos` - Get all todo lists (admin sees all, users see their own)
- `POST /api/todos` - Create a new todo list
- `GET /api/todos/{id}` - Get a specific todo list
- `PUT /api/todos/{id}` - Update a todo list
- `DELETE /api/todos/{id}` - Delete a todo list (soft delete)

### Todo Items
- `GET /api/todos/{listId}/items` - Get all items in a todo list
- `POST /api/todos/{listId}/items` - Create a new todo item
- `PUT /api/todos/{listId}/items/{itemId}` - Update a todo item
- `DELETE /api/todos/{listId}/items/{itemId}` - Delete a todo item (soft delete)

## Predefined Users

1. Regular User:
   - Username: `user1`
   - Password: `password1`

2. Admin User:
   - Username: `admin`
   - Password: `admin123`

## Getting Started

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```
3. Run the application:
   ```bash
   go run cmd/main.go
   ```

The server will start on port 8080.

## Authentication

To authenticate, send a POST request to `/login` with the following JSON body:
```json
{
    "username": "user1",
    "password": "password1"
}
```

The response will include a JWT token that should be included in subsequent requests in the Authorization header:
```
Authorization: Bearer <token>
```

## Project Structure

```
todo-app/
├── cmd/
│   └── main.go
├── internal/
│   ├── handlers/
│   │   └── handlers.go
│   ├── middleware/
│   │   └── auth.go
│   ├── models/
│   │   └── models.go
│   └── services/
│       └── mock_service.go
├── pkg/
│   └── auth/
│       └── jwt.go
└── README.md
```

## Dependencies

- github.com/gorilla/mux - HTTP router
- github.com/dgrijalva/jwt-go - JWT implementation

## Notes

- The application uses in-memory storage (mock service) instead of a database
- All delete operations are soft deletes (items are marked as deleted but not removed)
- The completion percentage of a todo list is calculated based on completed items
- All timestamps are in UTC 