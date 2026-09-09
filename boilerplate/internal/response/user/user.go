package response

import "boilerplate/internal/model"

type GetUserResponse struct {
	User  model.User `json:"user,omitempty"`
	Error string     `json:"error,omitempty"`
}

type CreateUserResponse struct {
	User  model.User `json:"user,omitempty"`
	Error string     `json:"error,omitempty"`
}
