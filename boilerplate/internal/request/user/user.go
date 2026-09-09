package request

type GetUserRequest struct {
	ID int64
}

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
