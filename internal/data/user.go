package data

type CreateUserRequest struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
}

type CreateUserResponse struct {
	ID string `json:"id"`
}