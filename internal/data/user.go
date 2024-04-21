package data

import (
	"time"

	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
)

type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CreateUserResponse struct {
	ID string `json:"id"`
}

type GetUserByIDRequest struct {
	ID string `json:"id"`
}

type GetUserByIDResponse struct {
	*models.User
	Friends int `json:"friends"`
}

type UpdateUserRequest struct {
	Name string `json:"name"`
	// Email	 string `json:"email"`
	Username string    `json:"username"`
	Avatar   string    `json:"avatar"`
	QuitDate time.Time `json:"quit_date"`
}

type UpdateUserResponse struct {
	ID string `json:"id"`
}

type SearchUsersByUsernameRequest struct {
	Username string `json:"username"`
	Limit    int    `json:"limit,omitempty"`
}

type SearchUsersByUsernameResponse struct {
	Users []*models.ShortUser `json:"users"`
}
