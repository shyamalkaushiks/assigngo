# assigngo

# gRPC User Service

This is a simple gRPC service for managing users.

## Building and Running the Application

### Prerequisites

- Go 1.16 or higher

### Steps

1. Clone the repository:

 
   git clone https://github.com/shyamalkaushiks/assigngo
   cd assigngo

2   Install the necessary dependencies:

go mod tidy

3 run the code 
go run main.go

run go test
4 all test cases passing

Accessing the gRPC Service Endpoints
grpcurl -plaintext -d '{"id": 1}' localhost:50051 UserService/GetUserByID
grpcurl -plaintext -d '{"ids": [1, 2]}' localhost:50051 UserService/GetUsersByIDs
grpcurl -plaintext -d '{"city": "LA"}' localhost:50051 UserService/SearchUsers


