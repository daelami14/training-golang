package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"

	"training-golang/assignment-golang/wallet-server/entity"
	grpcHandler "training-golang/assignment-golang/wallet-server/handler/grpc"
	pb "training-golang/assignment-golang/wallet-server/proto/wallet_service/v1"
	"training-golang/assignment-golang/wallet-server/repository/postgres_gorm"
	"training-golang/assignment-golang/wallet-server/service"
)

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	dsn := "postgresql://postgres:postgres@localhost:5432/postgres"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "assignment2_wallet.", // schema name
			SingularTable: false,
		}})
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	// Migrate the schema
	err = db.AutoMigrate(entity.Transaction{}, entity.Wallet{})
	if err != nil {
		fmt.Println("Failed to migrate database schema:", err)
	} else {
		fmt.Println("Database schema migrated successfully")
	}

	repo := postgres_gorm.NewWalletRepository(db) // Initialize your repository implementation
	walletService := service.NewWalletService(repo)
	walletHandler := grpcHandler.NewWalletHandler(walletService)

	grpcServer := grpc.NewServer()
	pb.RegisterWalletServiceServer(grpcServer, walletHandler)

	// Register reflection service on gRPC server.
	reflection.Register(grpcServer)

	log.Printf("gRPC server started at %s", listen.Addr().String())
	if err := grpcServer.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
