package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/hex4coder/ams/common"
	pb "github.com/hex4coder/ams/common/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	gatewayServerAddr = common.EnvString("USER_GATEWAY_SERVER_ADDR", ":8080")
	userServiceAddr   = "localhost:3000"
)

func main() {

	// create grpc client connection
	conn, err := grpc.NewClient(userServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect to GRPC Server at user service on address %s", userServiceAddr)
	}
	defer conn.Close()

	c := pb.NewUserServiceClient(conn)
	fmt.Printf("Connected to GRPC Server at user service address %s\n", userServiceAddr)

	mux := http.NewServeMux()
	handler := NewHandler(c)
	handler.registerRoutes(mux)

	log.Printf("Starting user gateway on address %s", gatewayServerAddr)

	if err := http.ListenAndServe(gatewayServerAddr, mux); err != nil {
		log.Fatalf("Failed to start user gateway on address %s", gatewayServerAddr)
	}
}
