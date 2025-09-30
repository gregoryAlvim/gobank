# API Gobank

Bem-vindo ao Gobank, uma API backend para uma aplicação bancária simples construída com Go.

## 🚀 Tecnologias

Este projeto foi desenvolvido com as seguintes tecnologias:

- **Go:** Uma linguagem de programação compilada e estaticamente tipada, projetada no Google.
- **PostgreSQL:** Um poderoso sistema de banco de dados objeto-relacional de código aberto.
- **Docker:** Uma plataforma para desenvolver, enviar e executar aplicações em contêineres.
- **Gorilla Mux:** Um poderoso roteador de URL e despachante para Go.
- **Tern:** Uma ferramenta de migração de banco de dados para Go.

## ✨ Funcionalidades

- Criar e gerenciar contas
- Autenticação baseada em JWT
- Documentação da API com Swagger
- Migrações de banco de dados com Tern

## 🏁 Começando

Estas instruções fornecerão uma cópia do projeto em execução em sua máquina local para fins de desenvolvimento e teste.

### Pré-requisitos

- [Go](https://golang.org/doc/install)
- [Docker](https://docs.docker.com/get-docker/)
- [Docker Compose](https://docs.docker.com/compose/install/)
- [Make](https://www.gnu.org/software/make/)

### Instalação

1.  **Clone o repositório:**

    ```bash
    git clone https://github.com/gregoryAlvim/gobank.git
    cd gobank
    ```

2.  **Instale as dependências do Go:**

    ```bash
    go mod download
    ```

3.  **Configure as variáveis de ambiente:**

    Crie um arquivo `.env` na raiz do projeto e adicione as seguintes variáveis:

    ```env
    DB_USER=user
    DB_PASSWORD=password
    DB_NAME=bank
    DB_HOST=localhost
    DB_PORT=5432
    DATABASE_URL=postgres://user:password@localhost:5432/bank?sslmode=disable
    ```

## 🐳 Executando o Projeto com Docker

A maneira mais fácil de executar o projeto é usando o Docker Compose.

1.  **Construa e inicie os contêineres:**

    ```bash
    docker compose up --build
    ```

    Este comando iniciará um contêiner PostgreSQL e o contêiner da aplicação Go. A API estará disponível em `http://localhost:8080`.

2.  **Parando os contêineres:**

    ```bash
    docker compose down
    ```

## 🏃 Executando o Projeto sem Docker

Você também pode executar o projeto localmente sem o Docker.

1.  **Inicie um banco de dados PostgreSQL:**

    Você pode usar o Docker para iniciar uma instância do PostgreSQL:

    ```bash
    docker run --name gobank-db -e POSTGRES_USER=user -e POSTGRES_PASSWORD=password -e POSTGRES_DB=bank -p 5432:5432 -d postgres
    ```

2.  **Execute as migrações do banco de dados:**

    ```bash
    make migrate-up
    ```

3.  **Execute a aplicação:**

    ```bash
    go run cmd/api/main.go
    ```

    A API estará disponível em `http://localhost:8080`.

## Migrações

As migrações de banco de dados são gerenciadas com `tern`. Você pode usar os seguintes comandos `make` para executá-las:

- `make migrate-up`: Aplica todas as migrações disponíveis.
- `make migrate-down`: Reverte todas as migrações.

## 📄 Licença

Este projeto está licenciado sob a Licença MIT - veja o arquivo [LICENSE.md](LICENSE.md) para mais detalhes.
