package main

import (
	"database/sql"
	"fmt"
	"grpc/internal/pb"
	"grpc/internal/services"
	"grpc/internal/store"
	"net"

	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./database/db.sqlite")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := store.NewCategory(db)
	categoryService := services.NewCategoryService(*categoryDB)

	grpcServer := grpc.NewServer()
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	reflection.Register(grpcServer)

	fmt.Print("gRPC Server running on port: 50051\n")
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}
}
