package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/vaibhavjain01/usermgt/usermgt"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserMgtServer
}

func (s *UserManagementServer) CreateNewUser (ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received: %v", in.GetName())
	var user_id int32 = int32(rand.Intn(1000));
	return &pb.User{Name: in.GetName(), Age: in.GetAge(), Id: user_id}, nil
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