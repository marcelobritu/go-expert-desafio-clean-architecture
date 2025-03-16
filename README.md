# Go Expert - Desafio Clean Architecture

Este projeto é um exemplo de implementação de Clean Architecture em Go.

## Descrição

Este projeto é uma aplicação de sistema de pedidos que utiliza a arquitetura limpa para separar as preocupações e facilitar a manutenção e escalabilidade do código.

## Pré-requisitos

- Go 1.16+
- Docker
- Docker Compose
- [migrate](https://github.com/golang-migrate/migrate)

## Instalação

1. Clone o repositório:

   ```sh
   git clone https://github.com/marcelobritu/go-expert-desafio-clean-architecture.git
   cd go-expert-desafio-clean-architecture
   ```

2. Instale as dependências:

   ```sh
   go mod download
   ```

3. Inicie os containers Docker:
   ```sh
   docker-compose up -d
   ```

## Migrações do Banco de Dados

Para iniciar o banco de dados e popular com dados, execute:

```sh
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose up
```

ou

```sh
make migrate
```

Para desfazer as tabelas e os dados, execute:

```sh
migrate -path=sql/migrations -database "mysql://root:root@tcp(localhost:3306)/orders" -verbose down
```

ou

```sh
make migratedown
```

## Executando a Aplicação

Para rodar a aplicação, execute:

```sh
go run cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go
```
