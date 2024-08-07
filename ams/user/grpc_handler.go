package main

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/hex4coder/ams/common"
	pb "github.com/hex4coder/ams/common/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedUserServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	handler := &grpcHandler{}
	pb.RegisterUserServiceServer(grpcServer, handler)
}

func (h *grpcHandler) CreateUser(ctx context.Context, userReq *pb.CreateUserRequest) (*pb.User, error) {
	fmt.Println("Creating user in GRPC")

	u, err := uuid.NewUUID()

	if err != nil {
		return nil, err
	}

	user := &pb.User{
		ID:       u.String(),
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	return user, nil
}

func (h *grpcHandler) GetUser(ctx context.Context, getReq *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	if getReq.ID != "1" {
		return nil, common.ErrNoUser
	}

	return &pb.GetUserResponse{
		Email: "admin@google.com",
		Name:  "Administrator",
	}, nil
}
