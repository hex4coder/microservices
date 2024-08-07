package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/hex4coder/ams/common"
	"google.golang.org/grpc"
)

var (
	grpcUserServiceAddr = common.EnvString("GRPC_USER_SERVICE_ADDRESS", "localhost:3000")
)

// main
func main() {
	// grpc server for user service
	server := grpc.NewServer()

	l, err := net.Listen("tcp", grpcUserServiceAddr)
	if err != nil {
		log.Fatalf("Failed to start GRPC Server User Service at address %s", grpcUserServiceAddr)
	}
	defer l.Close()

	// user store
	store := NewStore()
	service := NewService(store)
	NewGRPCHandler(server)

	fmt.Printf("User service in GRPC run at address %s\n", grpcUserServiceAddr)

	service.CreateUser(context.Background())

	err = server.Serve(l)

	if err != nil {
		log.Fatalf("Failed to serve user service at %s", grpcUserServiceAddr)
	}
}
