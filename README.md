# Gobank API

Welcome to Gobank, a backend API for a simple banking application built with Go.

## 🚀 Technologies

This project is developed with the following technologies:

- **Go:** A statically typed, compiled programming language designed at Google.
- **PostgreSQL:** A powerful, open-source object-relational database system.
- **Docker:** A platform for developing, shipping, and running applications in containers.
- **Gorilla Mux:** A powerful URL router and dispatcher for Go.

## ✨ Features

- Create and manage accounts
- JWT-based authentication
- API documentation with Swagger

## 🏁 Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Installation

1.  **Clone the repository:**

    ```bash
    git clone https://github.com/gregoryAlvim/gobank.git
    cd gobank
    ```

2.  **Install Go dependencies:**

    ```bash
    go mod download
    ```

3.  **Set up environment variables:**

    Create a `.env` file in the root of the project and add the following variables:

    ```env
    DB_USER=user
    DB_PASSWORD=password
    DB_NAME=bank
    DB_HOST=localhost
    DB_PORT=5432
    DATABASE_URL=postgres://user:password@localhost:5432/bank?sslmode=disable
    ```

## 🐳 Running the Project with Docker

The easiest way to get the project running is by using Docker Compose.

1.  **Build and start the containers:**

    ```bash
    docker-compose up --build
    ```

    This command will start a PostgreSQL container and the Go application container. The API will be available at `http://localhost:8080`.

2.  **Stopping the containers:**

    ```bash
    docker-compose down
    ```

## 🏃 Running the Project without Docker

You can also run the project locally without Docker.

1.  **Start a PostgreSQL database:**

    You can use Docker to start a PostgreSQL instance:

    ```bash
    docker run --name gobank-db -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=bank -p 5432:5432 -d postgres
    ```

2.  **Run database migrations:**

    You'll need to manually apply the database migrations. Connect to the PostgreSQL database and run the SQL commands in `migrations/001_create_tables.sql`.

3.  **Run the application:**

    ```bash
    go run main.go
    ```

    The API will be available at `http://localhost:8080`.

## Migrations
To create the tables you need to run the migrations on file migrations/001_create_tables.sql

## 📄 License

This project is licensed under the MIT License - see the [LICENSE.md](LICENSE.md) file for details.
