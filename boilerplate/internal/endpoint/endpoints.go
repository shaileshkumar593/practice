package endpoint

import (
	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	GetUser    endpoint.Endpoint
	CreateUser endpoint.Endpoint
}
