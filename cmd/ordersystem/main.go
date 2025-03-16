package main

import (
	"database/sql"
	"fmt"
	"net"

	"github.com/marcelobritu/go-expert-desafio-clean-architecture/configs"
	"github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/infra/grpc/pb"
	"github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/infra/grpc/service"
	"github.com/marcelobritu/go-expert-desafio-clean-architecture/internal/infra/web/webserver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	db, err := sql.Open(configs.DBDriver, fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configs.DBUser, configs.DBPassword, configs.DBHost, configs.DBPort, configs.DBName))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	listOrdersUseCase := NewListOrdersUseCase(db)

	webserver := webserver.NewWebServer(configs.WebServerPort)
	webOrderHandler := NewWebOrderHandler(db)
	webserver.AddHandler("/order", webOrderHandler.ListOrders)
	fmt.Println("Starting web server on port", configs.WebServerPort)
	go webserver.Start()

	grpcServer := grpc.NewServer()
	createOrderService := service.NewOrderService(*listOrdersUseCase)
	pb.RegisterOrderServiceServer(grpcServer, createOrderService)
	reflection.Register(grpcServer)

	fmt.Println("Starting gRPC server on port", configs.GRPCServerPort)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", configs.GRPCServerPort))
	if err != nil {
		panic(err)
	}
	grpcServer.Serve(lis)
}
