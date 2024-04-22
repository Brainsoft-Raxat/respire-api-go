package data

import (
	"time"

	"github.com/Brainsoft-Raxat/respire-api-go/internal/models"
)

type CreateSessionRequest struct {
	UID       string    `json:"uid"`
	Count     int       `json:"count"`
	Timestamp time.Time `json:"timestamp"`
}

type CreateSessionResponse struct {
	ID string `json:"id"`
}

type GetSessionByIDRequest struct {
	ID string `json:"id"`
}

type GetSessionByIDResponse struct {
	Session *models.SmokeSession
}

type GetSessionByUserIDRequest struct {
	ID string `json:"id"`
}

type GetSessionByUserIDResponse struct {
	Sum int `json:"sum"`
}

type GetSessionByUserIDAndDateRequest struct {
	ID string `json:"id"`
	DR [2]time.Time
}

type GetSessionByUserIDAndDateResponse struct {
	Sum int `json:"sum"`
}

type UpdateSessionRequest struct {
	Count int `json:"count"`
}

type UpdateSessionResponse struct {
	ID string `json:"id"`
}
