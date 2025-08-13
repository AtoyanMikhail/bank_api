# <img id="go" src="https://devicon-website.vercel.app/api/go/plain.svg?color=%2300ACD7" width="40" /> Go Banking System

A comprehensive banking backend system built with Go, implementing modern backend development practices including RESTful APIs, gRPC services, microservices architecture, and cloud deployment.

## 🚀 Features

- **Account Management**: Create and manage bank accounts with different currencies
- **Transaction Processing**: Secure money transfers between accounts with ACID compliance
- **User Authentication**: JWT and PASETO token-based authentication
- **Session Management**: Automatic session handling with refresh tokens
- **gRPC APIs**: High-performance gRPC services with HTTP gateway
- **Async Processing**: Background workers for email verification and notifications
- **Role-based Access Control**: User roles and permissions system
- **API Documentation**: Auto-generated Swagger documentation
- **Database Management**: PostgreSQL with migrations and connection pooling
- **Testing**: Comprehensive unit tests with database mocking
- **DevOps**: Docker containerization and Kubernetes deployment

## 🏗️ Architecture

### Database Schema
The system uses PostgreSQL with the following main entities:
- **Users**: User accounts with authentication
- **Accounts**: Bank accounts with balances and currencies
- **Entries**: Transaction logs for balance changes
- **Transfers**: Money transfer records between accounts
- **Sessions**: User session management
- **Verify Emails**: Email verification tokens

### Tech Stack

**Backend Framework & Language:**
- **Go** 1.24+ - Primary programming language
- **Gin** - HTTP web framework
- **gRPC** - High-performance RPC framework

**Database & Caching:**
- **PostgreSQL** 15.4+ - Primary database
- **Redis** - Caching and message broker
- **SQLC** - Type-safe SQL code generation

**Authentication & Security:**
- **JWT** - JSON Web Tokens
- **PASETO** - Platform-Agnostic Security Tokens
- **bcrypt** - Password hashing

**DevOps & Deployment:**
- **Docker** - Containerization
- **Kubernetes** - Container orchestration
- **AWS EKS** - Managed Kubernetes service
- **GitHub Actions** - CI/CD pipeline

**Development Tools:**
- **GoMock** - Mock generation for testing
- **Testify** - Testing assertions
- **Viper** - Configuration management
- **Zerolog** - Structured logging
- **Asynq** - Distributed task queue

## 📋 Prerequisites

- **Go** 1.24 or higher
- **Docker** and Docker Compose
- **PostgreSQL** 15.4+
- **Redis** (for background workers)
- **Make** (for build automation)

## 🛠️ Installation & Setup

### 1. Clone the Repository
```bash
git clone https://github.com/AtoyanMikhai/bank_api.git
cd go-basic-bank
```

### 2. Environment Configuration
Copy the example environment file and configure your settings:
```bash
cp app.env.example app.env
```

Configure the following environment variables in `app.env`:
```env
DB_DRIVER=postgres
DB_SOURCE=postgresql://root:secret@localhost:5432/bank?sslmode=disable
MIGRATION_URL=file://db/migration
REDIS_ADDRESS=localhost:6379
SERVER_ADDRESS=0.0.0.0:8080
GRPC_SERVER_ADDRESS=0.0.0.0:9090
TOKEN_SYMMETRIC_KEY=12345678901234567890123456789012
ACCESS_TOKEN_DURATION=15m
REFRESH_TOKEN_DURATION=24h
EMAIL_SENDER_NAME=Bank System
EMAIL_SENDER_ADDRESS=noreply@bank.com
EMAIL_SENDER_PASSWORD=secret
```

### 3. Database Setup
Start PostgreSQL and Redis using Docker Compose:
```bash
docker-compose up -d
```

Or manually start the database:
```bash
# Start PostgreSQL
make postgres

# Create database
make db_create

# Run migrations
make migrateup
```

### 4. Build and Run
```bash
# Install dependencies
go mod download

# Generate code (SQLC, gRPC, mocks)
make sqlc
make proto
make mock

# Build the application
go build -o main main.go

# Run the server
./main
```

## 🐳 Docker Deployment

### Build Docker Image
```bash
docker build -t bank-api .
```

### Run with Docker Compose
```bash
docker-compose up -d
```

This will start:
- PostgreSQL database on port 5432
- Redis server on port 6379
- Bank API server on port 8080
- gRPC server on port 9090

## 📖 API Documentation

### REST API Endpoints

The server provides both REST and gRPC APIs:

- **REST API**: `http://localhost:8080`
- **gRPC API**: `localhost:9090`
- **Swagger Documentation**: `http://localhost:8080/swagger/`

#### Authentication Endpoints
```
POST /users          - Create new user
POST /users/login    - User login
POST /tokens/renew   - Renew access token
```

#### Account Management
```
GET    /accounts           - List accounts (paginated)
POST   /accounts           - Create new account
GET    /accounts/:id       - Get account by ID
```

#### Money Transfers
```
POST /transfers - Transfer money between accounts
```

### gRPC Services

The system also provides gRPC services for better performance:
- `CreateUser` - User registration
- `LoginUser` - User authentication
- `UpdateUser` - Update user information
- `VerifyEmail` - Email verification

## 🧪 Testing

### Run All Tests
```bash
make test
```

### Test Coverage
```bash
go test -v -cover ./...
```

### Database Testing
The project uses GoMock for database mocking in tests, ensuring:
- Independent test execution
- Fast test performance
- 100% test coverage capability

## 🔧 Development Commands

The project includes a comprehensive Makefile with useful commands:

```bash
# Database operations
make postgres          # Start PostgreSQL container
make db_create         # Create database
make db_drop           # Drop database
make migrateup         # Run all migrations
make migratedown       # Rollback migrations

# Code generation
make sqlc              # Generate SQLC code
make proto             # Generate gRPC code
make mock              # Generate mocks

# Testing
make test              # Run tests
make server            # Start server
```

## 🏛️ Project Structure

```
├── api/              # REST API handlers
├── db/
│   ├── migration/    # Database migrations
│   ├── query/        # SQL queries
│   └── sqlc/         # Generated SQLC code
├── gapi/             # gRPC API handlers
├── pb/               # Generated protobuf code
├── proto/            # Protocol buffer definitions
├── token/            # JWT/PASETO token makers
├── util/             # Utility functions
├── worker/           # Background workers
├── mail/             # Email service
├── docs/             # API documentation
├── aws/              # Kubernetes manifests
├── docker-compose.yml
├── Dockerfile
├── Makefile
└── main.go
```

## 🔐 Security Features

- **Password Encryption**: bcrypt hashing
- **Token-based Authentication**: JWT and PASETO support
- **Session Management**: Secure session handling
- **Role-based Access Control**: User permissions system
- **SQL Injection Prevention**: SQLC type-safe queries
- **CORS Protection**: Configurable CORS policies

## 📊 Database Features

- **ACID Transactions**: Guaranteed data consistency
- **Connection Pooling**: Efficient database connections
- **Migration Management**: Version-controlled schema changes
- **Deadlock Prevention**: Optimized transaction ordering
- **Database Isolation Levels**: Configurable isolation settings

## 🔄 Background Processing

The system implements asynchronous task processing using Redis and Asynq:

- **Email Verification**: Async email sending for user verification
- **Notification System**: Background workers for system notifications
- **Task Queuing**: Reliable task distribution and execution
- **Error Handling**: Automatic retry mechanisms for failed tasks

## 🌐 Production Deployment

### Kubernetes Deployment
The project includes Kubernetes manifests for production deployment:

```bash
# Apply Kubernetes configurations
kubectl apply -f aws/
```

### AWS EKS Deployment
- Automated CI/CD with GitHub Actions
- Container registry with AWS ECR
- Production database with AWS RDS
- Secret management with AWS Secrets Manager

### Local Development
For local development, use Docker Compose:

```bash
# Start all services
docker-compose up -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

## 📊 Performance & Monitoring

- **Connection Pooling**: Optimized database connections
- **Structured Logging**: JSON-formatted logs with Zerolog
- **Health Checks**: Built-in health check endpoints
- **Metrics**: Performance monitoring capabilities
- **Graceful Shutdown**: Proper cleanup on server termination

## 🧪 Testing Strategy

### Unit Tests
- Comprehensive test coverage for all business logic
- Database mocking with GoMock
- Isolated test execution

### Integration Tests
- End-to-end API testing
- Database transaction testing
- gRPC service testing

### Performance Tests
- Load testing capabilities
- Concurrent transaction testing
- Deadlock prevention validation

## 🚀 Getting Started

1. **Clone and Setup**:
   ```bash
   git clone https://github.com/AtoyanMikhai/bank_api.git
   cd go-basic-bank
   make postgres
   make db_create
   make migrateup
   ```

2. **Start Development Server**:
   ```bash
   make server
   ```

3. **Access APIs**:
   - REST API: http://localhost:8080
   - Swagger Docs: http://localhost:8080/swagger/
   - gRPC: localhost:9090

4. **Run Tests**:
   ```bash
   make test
   ```

## 🤝 Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

If you encounter any issues or have questions:

2. Review the API documentation at `/swagger/`
3. Check the database schema documentation
4. Ensure all prerequisites are installed and configured correctly

---

*Happy Banking! 🏦*
