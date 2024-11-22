# Go Application with GORF and MongoDB

## Overview

This Go-based application leverages **GORF** (Go ORM Framework) for database interaction and **MongoDB** as the primary database. The architecture is designed for scalability, simplicity, and high performance, suitable for modern, production-ready web applications.

## Features

* **GORF Integration** : Simplifies database queries with an intuitive ORM layer.
* **MongoDB** : A powerful NoSQL database for both structured and unstructured data.
* **RESTful API** : Fully functional endpoints for seamless interaction with the application.
* **Environment Configurations** : Supports multiple environments such as development, staging, and production.
* **Docker Support** : Provides containerized deployment for portability and scalability.

## Prerequisites

Ensure you have the following installed on your system:

* [Go](https://golang.org/doc/install) (version 1.18+)
* [MongoDB](https://www.mongodb.com/try/download/community)
* [Docker](https://www.docker.com/products/docker-desktop) (for optional containerization)

## Setup

### 1. Clone the Repository

```bash
git clone https://github.com/p-shubh/goFr_with_mongo.git
cd goFr_with_mongo
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Configure Environment Variables

Update the `.env` file in the `configs` directory with your MongoDB connection details:

```dotenv
MONGO_URI=mongodb://localhost:27017
DATABASE_NAME=your_database_name
```

### 4. Run the Application

Start the application locally:

```bash
go run main.go
```

### 5. Access the API

The API will be available at:

```
http://localhost:8080
```

## Project Structure

```plaintext
.
├── main.go                           # Entry point of the application
├── .vscode/                          # VSCode launch configurations
│   └── launch.json
├── config/                           # Configuration setup
│   ├── config.go
│   └── type.go
├── configs/                          # Environment-specific configuration files
│   ├── .dev.env
│   ├── .staging.env
│   ├── .prod.env
│   └── .env
├── connection/database/mongoDb/      # MongoDB connection setup
│   ├── database.go
│   └── type.go
├── router/                           # API routes
│   ├── router.go
│   └── type.go
├── .gitignore                        # Git ignore file
├── go.mod                            # Dependency management
├── go.sum                            # Dependency versions
├── README.md                         # Project documentation
├── makefile                          # Build automation file
└── Dockerfile                        # Docker configuration
```

## API Endpoints

| Method | Endpoint        | Description          |
| ------ | --------------- | -------------------- |
| GET    | `/items`      | Retrieve all items   |
| POST   | `/items`      | Create a new item    |
| PUT    | `/items/{id}` | Update an item by ID |
| DELETE | `/items/{id}` | Delete an item by ID |

## Docker Setup

### 1. Build the Docker Image

```bash
docker build -t go-application .
```

### 2. Run the Container

```bash
docker run -p 8080:8080 go-application
```

### 3. Verify the Container

Access the API running inside the container at:

```
http://localhost:8080
```

## Contributing

We welcome contributions! Please fork this repository, create a feature branch, and submit a pull request. For major changes, open an issue to discuss your ideas first.

## License

This project is licensed under the MIT License. See the [LICENSE](https://chatgpt.com/c/LICENSE) file for more details.

---

Let me know if you need further adjustments!
