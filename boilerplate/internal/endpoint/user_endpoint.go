package endpoint

import (
	"context"

	"boilerplate/internal/service"

	userreq "boilerplate/internal/request/user"
	userresp "boilerplate/internal/response/user"

	"github.com/go-kit/kit/endpoint"
)

/*type Endpoints struct {
	GetUser    endpoint.Endpoint
	CreateUser endpoint.Endpoint
}*/

// --------------------------------------------------
// Make User Endpoints
// --------------------------------------------------

func MakeEndpoints(s service.IService) Endpoints {
	return Endpoints{
		GetUser:    MakeGetUserEndpoint(s),
		CreateUser: MakeCreateUserEndpoint(s),
	}
}

// --------------------------------------------------
// Get User Endpoint
// --------------------------------------------------

func MakeGetUserEndpoint(s service.IService) endpoint.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {

		req := request.(userreq.GetUserRequest)

		user, err := s.GetUser(
			ctx,
			req.ID,
		)

		if err != nil {
			return userresp.GetUserResponse{
				Error: mapError(err),
			}, nil
		}

		return userresp.GetUserResponse{
			User: user,
		}, nil
	}
}

// --------------------------------------------------
// Create User Endpoint
// --------------------------------------------------

func MakeCreateUserEndpoint(s service.IService) endpoint.Endpoint {
	return func(
		ctx context.Context,
		request interface{},
	) (interface{}, error) {

		req := request.(userreq.CreateUserRequest)

		user, err := s.CreateUser(
			ctx,
			req.Name,
			req.Email,
		)

		if err != nil {
			return userresp.CreateUserResponse{
				Error: mapError(err),
			}, nil
		}

		return userresp.CreateUserResponse{
			User: user,
		}, nil
	}
}

// --------------------------------------------------
// Error Mapping
// --------------------------------------------------

func mapError(err error) string {
	if err == nil {
		return ""
	}

	return err.Error()
}
