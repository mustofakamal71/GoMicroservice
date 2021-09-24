package main

import (
	"net"
	"os"

	GoMicroserviceGRPCCore "GoMicroservice/core"
	GoMicroserviceGRPC "GoMicroservice/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// configure our core service
	userService := GoMicroserviceGRPCCore.NewService()
	// configure our gRPC service controller
	userServiceController := NewUserServiceController(userService)
	// start a gRPC server
	server := grpc.NewServer()
	GoMicroserviceGRPC.RegisterUserServiceServer(server, userServiceController)
	reflection.Register(server)
	con, err := net.Listen("tcp", os.Getenv("GRPC_ADDR"))
	if err != nil {
		panic(err)
	}
	err = server.Serve(con)
	if err != nil {
		panic(err)
	}
}
