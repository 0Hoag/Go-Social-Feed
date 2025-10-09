# 🧩 Go-Social-Feed

**Go-Social-Feed** is a personal backend project developed in **Go** after my internship at **Tanca**.  
It’s inspired by the internal social feed system I worked on — a micro social network for employees to share posts, interact, and stay updated.

This project demonstrates practical backend skills: **Clean Architecture**, **Swagger**, **unit tests**, **RabbitMQ**, and **Docker** integration.

---

## 🚀 Features

- 🧱 **Clean Architecture** for scalability & maintainability  
- 📄 **Swagger** API documentation  
- 🧩 RESTful APIs built with **Gin**  
- 💬 Create, list, and interact with posts  
- 🐇 **RabbitMQ** integration for async tasks (e.g. background cleanup, notifications)  
- 🧪 **Unit tests** for core business logic  
- 🐳 Containerized with **Docker Compose**

---

## 🧰 Tech Stack

| Category | Technologies |
|-----------|--------------|
| **Language** | Go (Golang) |
| **Framework** | Gin |
| **Database** | MongoDB |
| **Message Queue** | RabbitMQ |
| **Testing** | Go test, Testify |
| **Documentation** | Swagger (swaggo/gin-swagger) |
| **Containerization** | Docker, Docker Compose |

---

## 🧩 Project Structure

go-social-feed/
├── cmd/
│   └── server/
│       └── main.go
├── internal/
│   ├── domain/
│   │   └── post.go
│   ├── usecase/
│   │   └── post_usecase.go
│   ├── repository/
│   │   └── post_mongo.go
│   ├── delivery/
│   │   └── http/
│   │       └── post_handler.go
│   └── queue/
│       └── rabbitmq.go
├── configs/
│   └── config.go
├── docker/
│   └── docker-compose.yml
├── go.mod
├── go.sum
├── Dockerfile
└── README.md


## ⚙️ Installation & Setup

### 1️⃣ Clone the repository
```bash
git clone https://github.com/0Hoag/go-social-feed.git
cd go-social-feed

### 2️⃣ Create an .env file
```bash
PORT=8080
MONGO_URI=mongodb://localhost:27017
MONGO_DB=social_feed
RABBITMQ_URI=amqp://guest:guest@localhost:5672/

### 3️⃣ Run with Docker
```bash
docker-compose up --build
Or run locally:
go mod tidy
go run ./cmd/app

### 4️⃣ Access
```bash
API Base URL: http://localhost:8080/api/v1
Swagger UI: http://localhost:8080/swagger/index.html