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
git clone https://github.com/Umairmaseed/Custodial-Wallet-POC.git

# Change Directory
cd Custodial-Wallet-POC
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
├── handlers            # used as bridge between controllers and services
├── logs
├── middlewares         # request/response middlewares
│   └── validators      # data/request validators
├── models
├── routes              # router initialization
├──  services            # general service
└── utils               #Utilities
```

## Project Future Structure

```
├── controllers/
│   ├── wallet_controller.go           # Handles API endpoints for wallet operations
│   └── transaction_controller.go      # Initiates/send internal transactions
├── docs                                # swagger files
├── handlers                            # used as bridge between controllers and services
├── middlewares/
│   └── validators/
│       └── wallet_validators.go       # Validate wallet creation / transaction input

├── models/
│   ├── wallet.go                      # Wallet DB models
│   └── transaction.go                 # Transaction history/log model

├── routes/
│   └── wallet_routes.go               # Define wallet-related routes

├── services/
│   ├── wallet/                        # Wallet-specific business logic
│   │   ├── generator.go               # Wallet generator using Shamir
│   │   ├── signer.go                  # Internal signer using private key shares
│   │   └── transfer.go                # Logic to send asset internally
│   ├── chains/                        # Logic per blockchain (ERC20, TRC20, etc.)
│   │   ├── ethereum.go
│   │   └── tron.go
│   ├── shamir/                        # Shamir Secret Sharing helper module
│   │   ├── split.go
│   │   └── reconstruct.go
│   └── keyvault/                      # Storage/retrieval for key shares

├── utils/
│   ├── crypto.go                # General ECDSA, AES helpers
│   └── constants.go                   # Network IDs, gas fees, etc.
```
