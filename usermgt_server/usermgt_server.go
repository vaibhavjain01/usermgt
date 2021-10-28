package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/vaibhavjain01/usermgt/usermgt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"google.golang.org/grpc/codes"
	"google.golang.org/genproto/googleapis/rpc/code"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserMgtServer
}

func (s *UserManagementServer) CreateNewUser (ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	if len(in.GetName()) < 3 {
		log.Printf("Received Invalid Name: %v", in.GetName())
		error_status := &pb.Error_Status {
			Code: 400, Message: "Name should be more than 3 characters", 
			Status: code.Code_INVALID_ARGUMENT , Details: nil,
		}
		return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: -1, ErrorDetails: &pb.Error{Error: error_status}}, status.Error(codes.InvalidArgument, "id was not found")
	}

	log.Printf("Received: %v", in.GetName())
	var user_id int32 = int32(rand.Intn(1000));
	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id, ErrorDetails: nil}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to launch listener %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserMgtServer(s, &UserManagementServer{})
	log.Printf("Server listening at %v", lis.Addr())

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to launch listener %v", err)
	}
}