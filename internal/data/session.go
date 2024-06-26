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
	ID     string    `json:"id"`
	Period string    `json:"period"`
	Start  time.Time `json:"start"`
	End    time.Time `json:"end"`
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

type GetUserStatRequest struct {
	ID string `json:"id"`
}

type GetUserStatResponse struct {
	CurrentStreak int `json:"current_streak"`
	BiggestStreak int `json:"biggest_streak"`
	SavedMoney    int `json:"saved_money"`
}
