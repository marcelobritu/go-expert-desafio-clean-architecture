FROM golang:latest as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-w -s" -o ordersystem cmd/ordersystem/main.go cmd/ordersystem/wire_gen.go

FROM scratch
WORKDIR /app

COPY --from=builder /app/ordersystem .
COPY --from=builder /app/.env .env

CMD [ "./ordersystem" ]