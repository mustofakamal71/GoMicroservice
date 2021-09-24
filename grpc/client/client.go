package client

import (
	"context"
	"time"

	"GoMicroservice"

	GoMicroserviceGRPC "GoMicroservice/grpc"

	"google.golang.org/grpc"
)

var defaultRequestTimeout = time.Second * 10

type grpcService struct {
	grpcClient GoMicroserviceGRPC.UserServiceClient
}

// NewGRPCService creates a new gRPC user service connection using the specified connection string.
func NewGRPCService(connString string) (GoMicroservice.Service, error) {
	conn, err := grpc.Dial(connString, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	return &grpcService{grpcClient: GoMicroserviceGRPC.NewUserServiceClient(conn)}, nil
}
func (s *grpcService) GetUsers(ids []int64) (result map[int64]GoMicroservice.User, err error) {
	result = map[int64]GoMicroservice.User{}
	req := &GoMicroserviceGRPC.GetUsersRequest{
		Ids: ids,
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.GetUsers(ctx, req)
	if err != nil {
		return
	}
	for _, grpcUser := range resp.GetUsers() {
		u := unmarshalUser(grpcUser)
		result[u.ID] = u
	}
	return
}
func (s *grpcService) GetUser(id int64) (result GoMicroservice.User, err error) {
	req := &GoMicroserviceGRPC.GetUsersRequest{
		Ids: []int64{id},
	}
	ctx, cancelFunc := context.WithTimeout(context.Background(), defaultRequestTimeout)
	defer cancelFunc()
	resp, err := s.grpcClient.GetUsers(ctx, req)
	if err != nil {
		return
	}
	for _, grpcUser := range resp.GetUsers() {
		// sanity check: only the requested user should be present in results
		if grpcUser.GetId() == id {
			return unmarshalUser(grpcUser), nil
		}
	}
	return result, GoMicroservice.ErrNotFound
}
func unmarshalUser(grpcUser *GoMicroserviceGRPC.User) (result GoMicroservice.User) {
	result.ID = grpcUser.Id
	result.Name = grpcUser.Name
	return
}
