Blog Backend Microservices
A microservices-based backend system written in Go, consisting of authentication and blog services.

Quick Start
Prerequisites
Docker
Docker Compose
Go 1.23+
Running the Application
Clone the repository
bash
Insert Code
Edit
Copy code
git clone https://github.com/yourusername/blog-backend.git
cd blog-backend
Start all services
bash
Insert Code
Edit
Copy code
docker-compose up -d
Stop services
bash
Insert Code
Edit
Copy code
docker-compose down
Architecture
Services
Auth Service (Port: 8080)

Handles user authentication
JWT token management
User management
Auth Database (Port: 5433)

PostgreSQL database
Stores user data
Credentials:
DB: authdb
User: user
Password: password
Blog Service (Port: 8082)

Manages blog posts
Handles blog operations
Requires authentication
Blog Database (Port: 5434)

PostgreSQL database
Stores blog data
Credentials:
DB: blogdb
User: user
Password: password
Project Structure
Insert Code
Edit
Copy code
.
├── auth-service/
│   ├── cmd/
│   ├── internal/
│   ├── Dockerfile
│   └── go.mod
├── blog-service/
│   ├── cmd/
│   ├── internal/
│   ├── Dockerfile
│   └── go.mod
└── docker-compose.yml
API Documentation
Auth Service Endpoints
POST /api/auth/register - Register new user
POST /api/auth/login - User login
GET /api/auth/verify - Verify JWT token
Blog Service Endpoints
GET /api/posts - Get all posts
POST /api/posts - Create new post
GET /api/posts/{id} - Get specific post
PUT /api/posts/{id} - Update post
DELETE /api/posts/{id} - Delete post
Environment Variables
Auth Service
env
Insert Code
Edit
Copy code
DB_HOST=auth-db
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=authdb
Blog Service
env
Insert Code
Edit
Copy code
DB_HOST=blog-db
DB_PORT=5432
DB_USER=user
DB_PASSWORD=password
DB_NAME=blogdb
Development
Tech Stack
Go 1.23
PostgreSQL
Docker & Docker Compose
JWT for authentication
RESTful API
Building Individual Services
bash
Insert Code
Edit
Copy code
# Build auth service
cd auth-service
docker build -t auth-service .

# Build blog service
cd blog-service
docker build -t blog-service .
Useful Docker Commands
bash
Insert Code
Edit
Copy code
# View logs
docker-compose logs

# View specific service logs
docker-compose logs auth-service
docker-compose logs blog-service

# Rebuild services
docker-compose build

# Rebuild and start specific service
docker-compose up --build auth-service

# Check status
docker-compose ps
Contributing
Fork the repository
Create your feature branch (git checkout -b feature/amazing-feature)
Commit your changes (git commit -m 'Add some amazing feature')
Push to the branch (git push origin feature/amazing-feature)
Open a Pull Request
License
This project is licensed under the MIT License - see the LICENSE file for details