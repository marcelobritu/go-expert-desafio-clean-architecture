# Go Expert - Desafio Clean Architecture

Este projeto é um exemplo de implementação de Clean Architecture em Go.

## Descrição

Este projeto é uma aplicação de sistema de pedidos que utiliza a arquitetura limpa para separar as preocupações e facilitar a manutenção e escalabilidade do código.

## Pré-requisitos

- Docker

## Instalação

1. Clone o repositório:

   ```sh
   git clone https://github.com/marcelobritu/go-expert-desafio-clean-architecture.git
   cd go-expert-desafio-clean-architecture
   ```

2. Inicie os containers Docker:
   ```sh
   docker-compose up -d
   ```

## Testando a Aplicação

### HTTP

Para testar a aplicação via HTTP, utilize ferramentas como `curl` ou Postman. Por exemplo:

```sh
curl -X GET http://localhost:8000/order -H "Content-Type: application/json"
```

### GraphQL

Para testar via GraphQL, acesse o endpoint GraphQL no navegador ou use ferramentas como GraphQL Playground. O endpoint estará disponível em:

```
http://localhost:8080
```

Exemplo de query:

```graphql
query queryOrders {
  orders {
    ID
    Price
    Tax
    FinalPrice
  }
}
```

### gRPC

Para testar via gRPC, utilize ferramentas como `evans` ou implemente um cliente gRPC.

#### Exemplo com `evans`:

1. Inicie o cliente `evans` no modo interativo:

   ```sh
   evans -r repl
   ```

2. Selecione o package `pb`:

   ```sh
   package pb
   ```

3. Selecione o serviço `OrderService`:

   ```sh
   service OrderService
   ```

4. Para listar pedidos (`ListOrders`), basta chamar:
   ```sh
   call ListOrders
   ```
