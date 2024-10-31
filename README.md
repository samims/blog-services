# Blog Backend Microservices

A robust, scalable backend system for a blogging platform, built with Go and following microservices architecture.

## 🚀 Features

- User authentication and authorization
- Blog post CRUD operations
- Scalable microservices architecture
- PostgreSQL database integration
- Docker containerization

## 🛠 Tech Stack

- Go 1.23+
- PostgreSQL
- Docker & Docker Compose
- JWT for authentication
- Golang Migrate for database migrations

## 🏗 Architecture

The system consists of two main microservices:

1. **Auth Service** (Port: 8080)
    - Handles user registration, login, and token management
    - Connected to Auth Database (Port: 5433)

2. **Blog Service** (Port: 8082)
    - Manages blog post operations (create, read, update, delete)
    - Connected to Blog Database (Port: 5434)

## 🚦 Getting Started

### Prerequisites

- Docker
- Docker Compose
- Go 1.23+

### Running the Application

1. Clone the repository:

    ```shell
    git clone https://github.com/samims/blog-services.git cd blog-services
    ```

2. Start all services 
```shell
docker-compose up -d
```

3. Stop all services
```shell
docker-compose down
```



## 📡 API Endpoints

### Auth Service

- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - User login
- `GET /api/auth/verify` - Verify JWT token

### Blog Service

- `GET /api/posts` - Get all posts
- `POST /api/posts` - Create a new post
- `GET /api/posts/{id}` - Get a specific post
- `PUT /api/posts/{id}` - Update a post
- `DELETE /api/posts/{id}` - Delete a post

## 🧪 Running Tests

To run tests for each service:
# Blog Backend Microservices

A robust, scalable backend system for a blogging platform, built with Go and following microservices architecture.

## 🚀 Features

- User authentication and authorization
- Blog post CRUD operations
- Scalable microservices architecture
- PostgreSQL database integration
- Docker containerization

## 🛠 Tech Stack

- Go 1.23+
- PostgreSQL
- Docker & Docker Compose
- JWT for authentication
- Golang Migrate for database migrations

## 🏗 Architecture

The system consists of two main microservices:

1. **Auth Service** (Port: 8080)
    - Handles user registration, login, and token management
    - Connected to Auth Database (Port: 5433)

2. **Blog Service** (Port: 8082)
    - Manages blog post operations (create, read, update, delete)
    - Connected to Blog Database (Port: 5434)

## 🚦 Getting Started

### Prerequisites

- Docker
- Docker Compose
- Go 1.23+

### Running the Application

1. Clone the repository:

    ```shell
    git clone https://github.com/yourusername/blog-backend.git cd blog-backend
    ```

2. Start all services
```shell
docker-compose up -d
```

3. Stop all services
```shell
docker-compose down
```



## 📡 API Endpoints

### Auth Service

- `POST /api/auth/register` - Register a new user
- `POST /api/auth/login` - User login
- `GET /api/auth/verify` - Verify JWT token

### Blog Service

- `GET /api/posts` - Get all posts
- `POST /api/posts` - Create a new post
- `GET /api/posts/{id}` - Get a specific post
- `PUT /api/posts/{id}` - Update a post
- `DELETE /api/posts/{id}` - Delete a post

## 🧪 Running Tests

To run tests for each service:
```shell
cd auth-service
go test ./...

cd ../blog-service
go test ./...
```

###  Useful Docker Commands
```shell
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
```

## 📊 Monitoring and Logging
Logs are available through Docker Compose logs
Consider integrating Prometheus and Grafana for advanced monitoring

## 🔒 Security Considerations
All endpoints are secured with JWT authentication
Passwords are hashed before storage
HTTPS is recommended for production deployments


## 🚀 Deployment

1. Ensure all environment variables are properly set 
2. Build and push Docker images to your container registry 
3. Deploy using Docker Compose or Kubernetes (sample manifests will be provided later in /k8s)


## 🔄 Continuous Integration / Continuous Deployment
GitHub Actions workflows are provided in .github/workflows
Automated testing and building on push to main branch
Automated deployment to staging environment on successful build

## 🔄 Continuous Integration / Continuous Deployment
GitHub Actions workflows are provided in .github/workflows
Automated testing and building on push to main branch
Automated deployment to staging environment on successful build


📬 Contact


Project Link: https://github.com/samims/blog-services
