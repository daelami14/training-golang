package main

import (
	"context"
	"log"

	user "training-golang/assignment-golang/user-server/proto/user_service/v1"
	wallet "training-golang/assignment-golang/wallet-server/proto/wallet_service/v1"

	"github.com/gin-gonic/gin"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}
	err := user.RegisterUserServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50052", opts)
	if err != nil {
		log.Fatalf("did not connect user service grpc: %v", err)
	}

	err = wallet.RegisterWalletServiceHandlerFromEndpoint(context.Background(), mux, "localhost:50051", opts)
	if err != nil {
		log.Fatalf("did not connect user wallet grpc: %v", err)
	}

	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	r.Any("*any", gin.WrapH(mux))

	log.Println("gateway run on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
