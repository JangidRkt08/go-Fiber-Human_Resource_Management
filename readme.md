# Go Fiber Human Resource Management (HRMS) API

A simple and clean REST API for managing Employee data, built using:

- **Go (Golang)**
- **Fiber v2** (Web Framework)
- **MongoDB Atlas (Cloud MongoDB)**
- **MongoDB Go Driver**

This project demonstrates CRUD operations (Create, Read, Update, Delete) for employees using RESTful API endpoints.

---

## 🚀 Features

- Connects to **MongoDB Atlas** using official MongoDB Go driver
- Employee CRUD operations:
  - `GET /employee` — fetch all employees
  - `POST /employee` — create new employee
  - `PUT /employee/:id` — update existing employee
  - `DELETE /employee/:id` — delete employee
- Clean Fiber routes
- BSON/JSON compatible models
- MongoDB ObjectID handling
- Fully working code inside a single `main.go`

---


## 📁 Project Structure

```
├── main.go
├── go.mod
├── Dockerfile

```


Everything (routes, DB connection, models) lives inside `main.go` for simplicity.

---

## 🧰 Tech Stack / Libraries Used

| Tool | Purpose |
|------|----------|
| **Fiber v2** (`github.com/gofiber/fiber/v2`) | HTTP server & routing |
| **MongoDB Go Driver** (`go.mongodb.org/mongo-driver`) | MongoDB connection and queries |
| **bson / primitive.ObjectID** | BSON encoding + MongoDB ObjectID handling |
| **context** | Timeout for DB connections |
| **log & time** | System utilities |

---

## 🗄️ Database Model

### **Employee**

```go
type Employee struct {
    ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
    Name   string             `json:"name"`
    Salary float64            `json:"salary"`
    Age    float64            `json:"age"`
}

```
## ▶️ Running the Project

### 1. Clone the Repository
```bash
git clone https://github.com/JangidRkt08/go-Fiber-Human_Resource_Management
cd go-Fiber-Human_Resource_Management

```
### 2. Install Dependencies
```bash
go mod tidy
```
### 3. Run the API
```bash
go run main.go
```