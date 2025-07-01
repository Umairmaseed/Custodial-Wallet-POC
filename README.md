## Table of Contents

- [Features](#features)
- [Prerequisites](#prerequisites)
- [Getting Started](#getting-started)
- [Project Structure](#project-structure)
- [Future Work](#future-work)

## Features

- [Gin](https://github.com/gin-gonic/gin)
- DotEnv with [Viper](https://github.com/spf13/viper)
- Request Validation with [Ozzo Validation](https://github.com/go-ozzo/ozzo-validation)
- Docker support
- Swagger with [gin-swagger](https://github.com/swaggo/gin-swagger)

## Prerequisites

- Git
- Docker
- [Go 1.20](https://go.dev/doc/install) and above

## Getting Started

```bash
# Clone Project
git clone <Repository>

# Change Directory
cd PrismFX_Wallet
```

#### Using Docker

```bash
# Build & Create Docker Containers
docker-compose up -d
```

#### Using Local Environment

```bash
# Copy Example Env file
cp ./env.example .env

# Download Modules
go mod download

# Build Project
go build -o go-starter .

# Run the Project
./go-starter
```

The application starts at port 8080:

- `GET /v1/ping` Health check endpoint, returns 'pong' message

---

- `GET /swagger/*` Auto created swagger endpoint

You can also see: http://localhost:8080/swagger/index.html

> If you want to add new routes and swagger docs, you should run `swag init`
> See: [Gin Swagger](https://github.com/swaggo/gin-swagger)

## Project Structure

```
├── controllers         # contains api functions and main business logic
├── docs                # swagger files
├── logs
├── middlewares         # request/response middlewares
│   └── validators      # data/request validators
├── models
├── routes              # router initialization
├──  services            # general service
└── utils               #Utilities
```
