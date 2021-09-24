package server

import (
	GoMicroservice "GoMicroservice"
	GoMicroserviceGRPC "GoMicroservice/grpc"
	"context"
)

// userServiceController implements the gRPC UserServiceServer interface.
type userServiceController struct {
	userService GoMicroservice.Service
}

// NewUserServiceController instantiates a new UserServiceServer.
func NewUserServiceController(userService GoMicroservice.Service) GoMicroserviceGRPC.UserServiceServer {
	return &userServiceController{
		userService: userService,
	}
}

// GetUsers calls the core service's GetUsers method and maps the result to a grpc service response.
func (ctlr *userServiceController) GetUsers(ctx context.Context, req *GoMicroserviceGRPC.GetUsersRequest) (resp *GoMicroserviceGRPC.GetUsersResponse, err error) {
	resultMap, err := ctlr.userService.GetUsers(req.GetIds())
	if err != nil {
		return
	}
	resp = &GoMicroserviceGRPC.GetUsersResponse{}
	for _, u := range resultMap {
		resp.Users = append(resp.Users, marshalUser(&u))
	}
	return
}

// marshalUser marshals a business object User into a gRPC layer User.
func marshalUser(u *GoMicroservice.User) *GoMicroserviceGRPC.User {
	return &GoMicroserviceGRPC.User{Id: u.ID, Name: u.Name}
}
