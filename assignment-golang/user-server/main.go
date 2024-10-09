package main

import (
	"fmt"
	"log"
	"net"
	"training-golang/assignment-golang/user-server/entity"

	grpcHandler "solution1/assignment-2-simple-ewallet-system/user-service/handler/grpc"
	pb "solution1/assignment-2-simple-ewallet-system/user-service/proto/user_service/v1"
	"solution1/assignment-2-simple-ewallet-system/user-service/repository/postgres_gorm"
	"solution1/assignment-2-simple-ewallet-system/user-service/service"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	listen, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dsn := "postgresql://postgres:admin@localhost:5432/user_db"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{SkipDefaultTransaction: true})

	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}

	// Migrate the schema
	err = db.AutoMigrate(entity.User{})
	if err != nil {
		fmt.Println("Failed to migrate database schema:", err)
	} else {
		fmt.Println("Database schema migrated successfully")
	}

	repo := postgres_gorm.NewUserRepository(db) // Initialize your repository implementation
	userService := service.NewUserService(repo)
	userHandler := grpcHandler.NewUserHandler(userService)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, userHandler)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	log.Printf("gRPC server started at %s", listen.Addr().String())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
