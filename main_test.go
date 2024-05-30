package main

import (
	"context"
	"testing"

	pb "grpc2/pb"

	"github.com/stretchr/testify/assert"
)

func TestGetUserByID(t *testing.T) {
	service := &UserService{
		users: []*pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NYC", Phone: 9876543210, Height: 6.0, Married: false},
		},
	}

	resp, err := service.GetUserByID(context.Background(), &pb.GetUserRequest{Id: 1})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Equal(t, "Steve", resp.User.Fname)

	_, err = service.GetUserByID(context.Background(), &pb.GetUserRequest{Id: 3})
	assert.Error(t, err)
}

func TestGetUsersByIDs(t *testing.T) {
	service := &UserService{
		users: []*pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NYC", Phone: 9876543210, Height: 6.0, Married: false},
		},
	}

	resp, err := service.GetUsersByIDs(context.Background(), &pb.GetUsersRequest{Ids: []int32{1, 2}})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Users, 2)

	resp, err = service.GetUsersByIDs(context.Background(), &pb.GetUsersRequest{Ids: []int32{3}})
	assert.NoError(t, err)
	assert.NotNil(t, resp)
	assert.Len(t, resp.Users, 0)
}

func TestSearchUsers(t *testing.T) {
	service := &UserService{
		users: []*pb.User{
			{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
			{Id: 2, Fname: "John", City: "NYC", Phone: 9876543210, Height: 6.0, Married: false},
		},
	}

	t.Run("Search by city", func(t *testing.T) {
		resp, err := service.SearchUsers(context.Background(), &pb.SearchRequest{City: "LA"})
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Users, 1)
		assert.Equal(t, "Steve", resp.Users[0].Fname)
	})

	t.Run("Search by phone", func(t *testing.T) {
		resp, err := service.SearchUsers(context.Background(), &pb.SearchRequest{Phone: 9876543210})
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Users, 1)
		assert.Equal(t, "John", resp.Users[0].Fname)
	})

	t.Run("Search by marital status", func(t *testing.T) {
		resp, err := service.SearchUsers(context.Background(), &pb.SearchRequest{Married: true})
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Users, 1)
		assert.Equal(t, "Steve", resp.Users[0].Fname)
	})

	t.Run("Search with multiple criteria", func(t *testing.T) {
		resp, err := service.SearchUsers(context.Background(), &pb.SearchRequest{City: "NYC", Married: false})
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Users, 1)
		assert.Equal(t, "John", resp.Users[0].Fname)
	})

	t.Run("Search with no matching criteria", func(t *testing.T) {
		resp, err := service.SearchUsers(context.Background(), &pb.SearchRequest{City: "SF"})
		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Users, 0)
	})
}
