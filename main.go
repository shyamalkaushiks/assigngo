package main

import (
	"context"
	"errors"
	"log"
	"net"

	pb "grpc2/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	users []*pb.User
}

func (s *UserService) GetUserByID(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	for _, user := range s.users {
		if user.Id == req.Id {
			return &pb.GetUserResponse{User: user}, nil
		}
	}
	return nil, errors.New("user not found")
}

func (s *UserService) GetUsersByIDs(ctx context.Context, req *pb.GetUsersRequest) (*pb.GetUsersResponse, error) {
	var foundUsers []*pb.User
	for _, id := range req.Ids {
		for _, user := range s.users {
			if user.Id == id {
				foundUsers = append(foundUsers, user)
				break
			}
		}
	}
	return &pb.GetUsersResponse{Users: foundUsers}, nil
}

func (s *UserService) SearchUsers(ctx context.Context, req *pb.SearchRequest) (*pb.SearchResponse, error) {
	var foundUsers []*pb.User
	for _, user := range s.users {
		matches := true
		if req.City != "" && user.City != req.City {
			matches = false
		}
		if req.Phone != 0 && user.Phone != req.Phone {
			matches = false
		}
		if req.Married && req.Married != user.Married {
			matches = false
		}
		if matches {
			foundUsers = append(foundUsers, user)
		}
	}
	log.Printf("SearchUsers request: %+v, found users: %+v", req, foundUsers)
	return &pb.SearchResponse{Users: foundUsers}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterUserServiceServer(s, &UserService{
		users: []*pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NYC", Phone: 9876543210, Height: 6.0, Married: false},
		},
	})
	log.Println("gRPC server started on port 50051...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
