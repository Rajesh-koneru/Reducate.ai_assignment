# Golang Authentication Backend (Gin + JWT + GORM + SQLite)

A simple backend authentication system built using Golang.  
It supports user registration, login, JWT-based authentication, and role-based access control (Admin/User).

---

##  Features

- User Signup (Registration)
- User Login
- Password Hashing using bcrypt
- JWT Authentication
- Role-Based Access (Admin / User)
- Protected Routes (Middleware)
- SQLite Database Integration
- GORM ORM for database handling

---

## 🛠 Tech Stack

- Golang
- Gin Web Framework
- GORM ORM
- SQLite Database
- JWT (JSON Web Token)
- bcrypt (Password Hashing)

---

## 📁 Project Structure
    backend-assignment/
    │
    ├── main.go
    ├── go.mod
    ├── database/
    │ └── database.go
    │
    ├── handlers/
    │ └── Auth.go
      └── Signin.go
      └── AdminDashboard.go
      └── UserDashboard.g0
    │
    ├── models/
    │ └── users.go
    │
    ├── middleware/
    │ └── JWT_auth.go
    │
    ├── utils/
    │ └── GenerateJWT.go
    │
    └── test.db


---

## 🔐 Authentication Flow

1. User registers using `/signup`
2. Password is hashed using bcrypt
3. User logs in using `/login`
4. Server verifies credentials
5. JWT token is generated and returned
6. Token is used to access protected routes

---

## 📡 API Endpoints

### 🔓 Public Routes

#### Signup
```
POST /signup
  Request:
  
  {
    "name": "John",
    "email": "john@example.com",
    "password": "123456"
  }
```
```
Response:

{
  "message": "User created successfully"
}
```

Login
```
POST /login

Request:

{
  "email": "john@example.com",
  "password": "123456"
}

```
```
Response:

{
  "token": "JWT_TOKEN_HERE"
}
```

Protected Routes

```
Get Profile
GET /profile
```
```
Headers:

Authorization: Bearer <token>
```
```
Response:

{
  "name": "John",
  "email": "john@example.com",
  "role": "user"
}
```
```
Get All Users (Admin Only)
GET /users
```
```
Headers:
Authorization: Bearer <token>
```
```
Response:

[
  {
    "id": 1,
    "name": "John",
    "email": "john@example.com"
  }
]
```

Setup Instructions

1. Clone repository
```
git clone https://github.com/your-username/repo-name.git
cd repo-name

```

2. Install dependencies
```
go mod tidy

```
3. Run project
```
go run main.go


```
5. Server runs at
```
http://localhost:8000

```
