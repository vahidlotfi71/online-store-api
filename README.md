# 🛍️ Online Store API

A RESTful API for an online store built with Go, Fiber, and GORM. This project demonstrates clean architecture, proper authentication, validation, and testing practices.

[![Go Version](https://img.shields.io/badge/Go-1.24-blue.svg)](https://golang.org)
[![Fiber](https://img.shields.io/badge/Fiber-v2.52-00ADD8.svg)](https://gofiber.io)
[![GORM](https://img.shields.io/badge/GORM-v1.31-red.svg)](https://gorm.io)

## 📋 Table of Contents

- [Features](#features)
- [Tech Stack](#tech-stack)
- [Project Structure](#project-structure)
- [Installation](#installation)
- [Configuration](#configuration)
- [Usage](#usage)
- [API Endpoints](#api-endpoints)
- [Testing](#testing)
- [CLI Commands](#cli-commands)
- [Authentication](#authentication)
- [Contributing](#contributing)

## ✨ Features

### Core Features
- 🔐 **JWT Authentication** - Secure token-based auth for users and admins
- 👥 **Role-Based Access Control** - Separate permissions for users and admins
- 🛒 **Product Management** - CRUD operations with soft deletes
- 📦 **Order Management** - Create and track orders with items
- 👤 **User Management** - Complete user profile management
- 🗑️ **Trash System** - Soft delete with restore capability
- 📄 **Pagination** - Efficient data pagination for all list endpoints
- ✅ **Validation** - Comprehensive request validation
- 🧪 **Testing** - Unit tests for critical components

### Technical Features
- Clean Architecture with Repository Pattern
- Transaction support for data consistency
- Custom validation rules (Phone, National ID, etc.)
- Resource transformation for API responses
- CLI tools for database management
- Environment-based configuration

## 🛠 Tech Stack

- **Framework:** [Fiber](https://gofiber.io) - Fast HTTP framework
- **ORM:** [GORM](https://gorm.io) - Database ORM
- **Database:** MySQL / SQLite (for testing)
- **Authentication:** JWT with [golang-jwt](https://github.com/golang-jwt/jwt)
- **Password Hashing:** bcrypt
- **CLI:** [Cobra](https://github.com/spf13/cobra)
- **Testing:** [Testify](https://github.com/stretchr/testify)
- **Environment:** [godotenv](https://github.com/joho/godotenv)

## 📁 Project Structure

```
online-store-api/
├── cmd/                    # CLI commands
│   ├── migrate.go         # Database migration
│   ├── seed.go            # Data seeding
│   ├── serve.go           # Start server
│   └── routes.go          # List routes
├── Config/                # Configuration
│   ├── App.go             # Fiber app
│   ├── DBConnection.go    # Database connection
│   └── EnvVars.go         # Environment variables
├── Controllers/           # HTTP handlers
│   ├── AdminController/
│   ├── AuthController/
│   ├── OrderController/
│   ├── ProductController/
│   └── UserController/
├── Middlewares/           # HTTP middlewares
│   ├── AdminAuthMiddleware.go
│   ├── AuthMiddleware.go
│   └── ValidationMiddleware.go
├── Models/                # Database models
│   ├── User/              # User repository
│   ├── Product/           # Product repository
│   ├── Admin/             # Admin repository
│   └── Order/             # Order repository
├── Resources/             # API response transformers
├── Routes/                # Route definitions
├── Rules/                 # Validation rules
├── Test/                  # Test files
│   ├── model/
│   └── utils/
├── Utils/                 # Utility functions
│   ├── Http/              # HTTP utilities
│   ├── Constants/         # Constants
│   ├── Password.go        # Password hashing
│   ├── Token.go           # JWT utilities
│   └── IsAdmin.go         # Role checking
├── Validations/           # Request validations
├── main.go                # Application entry point
├── go.mod                 # Go modules
└── .env.example           # Environment template
```

## 🚀 Installation

### Prerequisites

- Go 1.24 or higher
- MySQL 8.0+ (or MariaDB 10.5+)
- Git

### Steps

1. **Clone the repository**
```bash
git clone https://github.com/vahidlotfi71/online-store-api.git
cd online-store-api
```

2. **Install dependencies**
```bash
go mod download
go mod tidy
```

3. **Set up environment variables**
```bash
cp .env.example .env
```

Edit `.env` with your configuration:
```env
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_NAME=online_store
DB_CHARSET=utf8mb4
JWT_SECRET=your-secret-key-here
PORT=8080
```

4. **Run migrations**
```bash
go run main.go migrate
```

This will:
- Create all database tables
- Seed a default admin account:
  - Phone: `09123456789`
  - Password: `12345678`
  - National ID: `1111111111`

5. **Start the server**
```bash
go run main.go serve
```

The API will be available at `http://localhost:8080`

## ⚙️ Configuration

### Environment Variables

| Variable | Description | Required | Default |
|----------|-------------|----------|---------|
| `DB_HOST` | Database host | Yes | localhost |
| `DB_PORT` | Database port | Yes | 3306 |
| `DB_USER` | Database user | Yes | root |
| `DB_PASSWORD` | Database password | Yes | - |
| `DB_NAME` | Database name | Yes | online_store |
| `DB_CHARSET` | Database charset | Yes | utf8mb4 |
| `JWT_SECRET` | JWT signing secret | Yes | - |
| `PORT` | Server port | No | 8080 |

## 📖 Usage

### Starting the Server

```bash
# Using CLI command
go run main.go serve

# Or directly
go run main.go

# With custom port
go run main.go serve -p 3000
```

### Making Requests

**Register a new user:**
```bash
curl -X POST http://localhost:8080/register \
  -F "first_name=John" \
  -F "last_name=Doe" \
  -F "phone=09123456788" \
  -F "address=123 Main St" \
  -F "national_id=1234567890" \
  -F "password=mypassword"
```

**Login:**
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{"phone":"09123456788","password":"mypassword"}'
```

**Get products (authenticated):**
```bash
curl http://localhost:8080/products \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## 🔌 API Endpoints

### Public Routes

| Method | Endpoint | Description |
|--------|----------|-------------|
| POST | `/register` | Register new user |
| POST | `/login` | User login |
| POST | `/admin/login` | Admin login |
| POST | `/logout` | Logout |
| GET | `/products` | List products |
| GET | `/products/:id` | Get product details |

### User Routes (Requires Authentication)

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/user/profile` | Get user profile |
| POST | `/user/profile/update` | Update profile |
| GET | `/user/orders` | List user orders |
| POST | `/user/orders` | Create order |
| GET | `/user/orders/:id` | Get order details |

### Admin Routes (Requires Admin Role)

#### Users Management
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/admin/user` | List all users |
| GET | `/admin/user/show/:id` | Get user details |
| POST | `/admin/user/store` | Create user |
| POST | `/admin/user/update/:id` | Update user |
| POST | `/admin/user/delete/:id` | Soft delete user |
| GET | `/admin/user/restore/:id` | Restore deleted user |
| GET | `/admin/user/trash` | List deleted users |
| POST | `/admin/user/clear-trash` | Permanently delete users |

#### Products Management
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/admin/product` | List all products |
| GET | `/admin/product/show/:id` | Get product details |
| POST | `/admin/product/store` | Create product |
| POST | `/admin/product/update/:id` | Update product |
| POST | `/admin/product/delete/:id` | Soft delete product |
| GET | `/admin/product/restore/:id` | Restore deleted product |
| GET | `/admin/product/trash` | List deleted products |
| POST | `/admin/product/clear-trash` | Permanently delete products |

#### Orders Management
| Method | Endpoint | Description |
|--------|----------|-------------|
| GET | `/admin/order` | List all orders |
| GET | `/admin/order/show/:id` | Get order details |
| POST | `/admin/order/update/:id` | Update order status |
| GET | `/admin/order/trash` | List deleted orders |

### Pagination

All list endpoints support pagination:

```bash
# Default (page 1, 15 items per page)
GET /admin/user

# Custom pagination
GET /admin/user?page=2&per_page=20

# Maximum per_page is 100
```

Response includes metadata:
```json
{
  "data": [...],
  "metadata": {
    "totalPages": 5,
    "previousPage": 1,
    "nextPage": 3,
    "offset": 30,
    "limitPerPage": 15,
    "currentPage": 2
  }
}
```

## 🧪 Testing

### Run All Tests

```bash
go test ./...
```

### Run Tests with Coverage

```bash
go test -cover ./...
```

### Generate Coverage Report

```bash
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out -o coverage.html
```

### Run Specific Tests

```bash
# Test Utils
go test -v ./Test/utils

# Test Models
go test -v ./Test/model
```

### Using Makefile

```bash
# Run all tests
make test

# Run with verbose output
make test-verbose

# Generate coverage report
make coverage

# Run specific package tests
make test-utils
make test-models
```

### Test Coverage

Current test coverage includes:

- ✅ **Utils Package** (90%+)
  - Password hashing and verification
  - JWT token creation and validation
  
- ✅ **Models Package** (85%+)
  - User CRUD operations
  - Product CRUD operations
  - Soft delete and restore
  
- ✅ **Validation Rules** (80%+)
  - All validation rules tested
  - Edge cases covered

## 🖥️ CLI Commands

The project includes a powerful CLI for common tasks:

### Available Commands

```bash
# Show all commands
go run main.go help

# Show version
go run main.go version

# Start server
go run main.go serve
go run main.go serve -p 3000

# Database migration
go run main.go migrate
go run main.go migrate -f  # Force drop tables

# Seed database
go run main.go seed
go run main.go seed -u 100  # Seed 100 users

# List routes
go run main.go routes
```

### Migration

```bash
# Create tables
go run main.go migrate

# Drop and recreate tables
go run main.go migrate --force
```

### Seeding

```bash
# Seed 100 fake users (default)
go run main.go seed

# Seed custom number of users
go run main.go seed --users 500
```

Seeded users will have:
- Username: User1, User2, etc.
- Phone: 09120000001, 09120000002, etc.
- Password: `password`

## 🔐 Authentication

### JWT Authentication

The API uses JWT tokens for authentication with two roles:

- **user**: Regular user access
- **admin**: Administrative access

### Token Structure

```json
{
  "id": 1,
  "role": "user",
  "name": "John Doe",
  "phone": "09123456789",
  "exp": 1234567890
}
```

### Token Expiration

- **Standard**: 1 month
- **Remember Me**: 6 months

### Using Tokens

Include the token in the `Authorization` header:

```bash
Authorization: Bearer YOUR_JWT_TOKEN
```

### Getting a Token

**User Login:**
```bash
curl -X POST http://localhost:8080/login \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "09123456789",
    "password": "mypassword",
    "remember_me": true
  }'
```

**Admin Login:**
```bash
curl -X POST http://localhost:8080/admin/login \
  -H "Content-Type: application/json" \
  -d '{
    "phone": "09123456789",
    "password": "12345678"
  }'
```

## 📝 Validation

The API includes comprehensive validation for all inputs:

### Built-in Validation Rules

- **Required**: Field must be present
- **MinLength/MaxLength**: String length validation
- **LengthBetween**: String length range
- **PhoneNumber**: Iranian phone number validation
- **NationalID**: Iranian national ID validation with checksum
- **Numeric**: Numeric value validation
- **BooleanStrict**: Only accepts "true" or "false"
- **Optional**: Field is optional

### Custom Validation

Example of using validation:

```go
Middlewares.ValidationMiddleware([]Rules.FieldRules{
    {
        FieldName: "phone",
        Rules: []Rules.ValidationRule{
            Rules.Required,
            Rules.PhoneNumber(),
        },
    },
    {
        FieldName: "password",
        Rules: []Rules.ValidationRule{
            Rules.Required,
            Rules.LengthBetween(8, 16),
        },
    },
})
```

## 🗄️ Database Schema

### Users Table
```sql
- id (PK)
- first_name
- last_name
- phone (unique)
- address
- national_id (unique)
- password (hashed)
- role (user/admin)
- is_verified
- created_at
- updated_at
- deleted_at (soft delete)
```

### Products Table
```sql
- id (PK)
- name
- brand
- price
- description
- stock
- is_active
- created_at
- updated_at
- deleted_at (soft delete)
```

### Orders Table
```sql
- id (PK)
- user_id (FK)
- status (pending/paid/cancelled)
- total_price
- created_at
- updated_at
- deleted_at (soft delete)
```

### Order Items Table
```sql
- id (PK)
- order_id (FK)
- product_id (FK)
- quantity
- price
- created_at
- updated_at
```

## 🏗️ Architecture

### Repository Pattern

The project uses the Repository Pattern for data access:

```go
// Repository interface
type UserRepository interface {
    Create(dto UserCreateDTO) (User, error)
    FindByID(id uint) (User, error)
    Update(id uint, dto UserUpdateDTO) error
    SoftDelete(id uint) error
}

// Implementation
func Create(tx *gorm.DB, dto UserCreateDTO) (user Models.User, err error) {
    user = Models.User{
        FirstName: dto.FirstName,
        // ...
    }
    err = tx.Create(&user).Error
    return
}
```

### Resource Transformation

API responses are transformed using Resource classes:

```go
type UserResource struct {
    ID         uint      `json:"id"`
    FirstName  string    `json:"first_name"`
    LastName   string    `json:"last_name"`
    // Excludes sensitive data like password
}

func Single(u Models.User) UserResource {
    return UserResource{
        ID:        u.ID,
        FirstName: u.FirstName,
        LastName:  u.LastName,
    }
}
```

## 🤝 Contributing

Contributions are welcome! Here's how you can help:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

### Guidelines

- Write tests for new features
- Follow Go conventions and best practices
- Update documentation as needed
- Keep commits atomic and meaningful

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 👨‍💻 Author

**Vahid Lotfi**
- GitHub: [@vahidlotfi71](https://github.com/vahidlotfi71)

## 🙏 Acknowledgments

- [Fiber](https://gofiber.io) - Web framework
- [GORM](https://gorm.io) - ORM library
- [Testify](https://github.com/stretchr/testify) - Testing toolkit
- Go community for excellent tools and libraries

## 📞 Support

If you have any questions or issues, please:

1. Check the [documentation](#table-of-contents)
2. Search [existing issues](https://github.com/vahidlotfi71/online-store-api/issues)
3. Create a new issue if needed

## 🗺️ Roadmap

- [ ] Add email verification
- [ ] Implement payment gateway integration
- [ ] Add product categories
- [ ] Implement product search
- [ ] Add product reviews
- [ ] Implement caching with Redis
- [ ] Add rate limiting
- [ ] Docker support
- [ ] API documentation with Swagger
- [ ] GraphQL support

---

**⭐ If you find this project helpful, please give it a star!**